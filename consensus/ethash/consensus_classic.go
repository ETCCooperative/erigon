package ethash

import (
	"math/big"

	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/params"
)

var (
	calcDifficultyNoBombByzantium = makeDifficultyCalculatorClassic(true, true, true)
	calcDifficultyNoBombHomestead = makeDifficultyCalculatorClassic(false, true, true)
	calcDifficulty1010Pause       = makeDifficultyCalculatorClassic(false, false, true)
	calcDifficulty1010Explode     = makeDifficultyCalculatorClassic(false, false, false)
)

func CalcDifficulty_Classic(config *params.ChainConfig, time, parentTime uint64, parentDifficulty *big.Int, parentNumber uint64, parentUncleHash common.Hash) *big.Int {
	next := parentNumber + 1
	switch {
	case config.IsByzantium(next):
		return calcDifficultyNoBombByzantium(time, parentTime, parentDifficulty, parentNumber, parentUncleHash)
	case config.IsECIP1041(next):
		return calcDifficultyNoBombHomestead(time, parentTime, parentDifficulty, parentNumber, parentUncleHash)
	case config.IsECIP1010Disable(next):
		return calcDifficulty1010Explode(time, parentTime, parentDifficulty, parentNumber, parentUncleHash)
	case config.IsECIP1010(next):
		return calcDifficulty1010Pause(time, parentTime, parentDifficulty, parentNumber, parentUncleHash)
	case config.IsHomestead(next):
		return calcDifficultyHomestead(time, parentTime, parentDifficulty, parentNumber, parentUncleHash)
	default:
		return calcDifficultyFrontier(time, parentTime, parentDifficulty, parentNumber, parentUncleHash)
	}
}

func makeDifficultyCalculatorClassic(eip100b, defuse, pause bool) func(time, parentTime uint64, parentDifficulty *big.Int, parentNumber uint64, parentUncleHash common.Hash) *big.Int {
	return func(time, parentTime uint64, parentDifficulty *big.Int, parentNumber uint64, parentUncleHash common.Hash) *big.Int {
		// https://github.com/ethereum/EIPs/issues/100.
		// algorithm:
		// diff = (parent_diff +
		//         (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
		//        ) + 2^(periodCount - 2)

		bigTime := new(big.Int).SetUint64(time)
		bigParentTime := new(big.Int).SetUint64(parentTime)

		// holds intermediate values to make the algo easier to read & audit
		x := new(big.Int)
		y := new(big.Int)

		if eip100b {
			// (2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // 9
			x.Sub(bigTime, bigParentTime)
			x.Div(x, big9)
			if parentUncleHash == types.EmptyUncleHash {
				x.Sub(big1, x)
			} else {
				x.Sub(big2, x)
			}
			// max((2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // 9, -99)
			if x.Cmp(bigMinus99) < 0 {
				x.Set(bigMinus99)
			}
			// parent_diff + (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
			y.Div(parentDifficulty, params.DifficultyBoundDivisor)
			x.Mul(y, x)
			x.Add(parentDifficulty, x)

			// minimum difficulty can ever be (before exponential factor)
			if x.Cmp(params.MinimumDifficulty) < 0 {
				x.Set(params.MinimumDifficulty)
			}
		} else {
			// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.md
			// algorithm:
			// diff = (parent_diff +
			//         (parent_diff / 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
			//        ) + 2^(periodCount - 2)

			// 1 - (block_timestamp - parent_timestamp) // 10
			x.Sub(bigTime, bigParentTime)
			x.Div(x, big10)
			x.Sub(big1, x)

			// max(1 - (block_timestamp - parent_timestamp) // 10, -99)
			if x.Cmp(bigMinus99) < 0 {
				x.Set(bigMinus99)
			}
			// (parent_diff + parent_diff // 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
			y.Div(parentDifficulty, params.DifficultyBoundDivisor)
			x.Mul(y, x)
			x.Add(parentDifficulty, x)

			// minimum difficulty can ever be (before exponential factor)
			if x.Cmp(params.MinimumDifficulty) < 0 {
				x.Set(params.MinimumDifficulty)
			}
		}

		if defuse {
			return x
		}

		// exPeriodRef the explosion clause's reference point
		exPeriodRef := big.NewInt(int64(parentNumber) + 1)

		if pause {
			exPeriodRef.Set(params.ClassicChainConfig.ECIP1010Block)
		} else {
			// unpaused (exploded) difficulty bomb
			length := int64(2_000_000)
			exPeriodRef.Sub(exPeriodRef, big.NewInt(length))
		}

		// EXPLOSION

		// the 'periodRef' (from above) represents the many ways of hackishly modifying the reference number
		// (ie the 'currentBlock') in order to lie to the function about what time it really is
		//
		//   2^(( periodRef // EDP) - 2)
		//
		z := new(big.Int)
		z.Div(exPeriodRef, big.NewInt(100_000)) // (periodRef // EDP)
		if z.Cmp(big1) > 0 {                    // if result large enough (not in algo explicitly)
			z.Sub(z, big2)      // - 2
			z.Exp(big2, z, nil) // 2^
		} else {
			z.SetUint64(0)
		}
		x.Add(x, z)

		return x
	}
}

var (
	disinflationRateQuotient = uint256.NewInt(4)
	disinflationRateDivisor  = uint256.NewInt(5)
	ecip1017EraLen           = uint256.NewInt(5_000_000)
	big32                    = uint256.NewInt(32)
	big8                     = uint256.NewInt(8)
)

// As of "Era 2" (zero-index era 1), uncle miners and winners are rewarded equally for each included block.
// So they share this function.
func getEraUncleBlockReward(era *uint256.Int, blockReward *uint256.Int) *uint256.Int {
	return new(uint256.Int).Div(GetBlockWinnerRewardByEra(era, blockReward), big32)
}

// GetBlockUncleRewardByEra gets called _for each uncle miner_ associated with a winner block's uncles.
func GetBlockUncleRewardByEra(era *uint256.Int, header, uncle *types.Header, blockReward *uint256.Int) *uint256.Int {
	// Era 1 (index 0):
	//   An extra reward to the winning miner for including uncles as part of the block, in the form of an extra 1/32 (0.15625ETC) per uncle included, up to a maximum of two (2) uncles.
	if era.Cmp(uint256.NewInt(0)) == 0 {
		r := new(uint256.Int)
		r.Add(new(uint256.Int).SetUint64(uncle.Number.Uint64()), big8) // 2,534,998 + 8              = 2,535,006
		r.Sub(r, new(uint256.Int).SetUint64(header.Number.Uint64()))   // 2,535,006 - 2,534,999        = 7
		r.Mul(r, blockReward)                                          // 7 * 5e+18               = 35e+18
		r.Div(r, big8)                                                 // 35e+18 / 8                            = 7/8 * 5e+18

		return r
	}
	return getEraUncleBlockReward(era, blockReward)
}

// GetBlockWinnerRewardForUnclesByEra gets called _per winner_, and accumulates rewards for each included uncle.
// Assumes uncles have been validated and limited (@ func (v *BlockValidator) VerifyUncles).
func GetBlockWinnerRewardForUnclesByEra(era *uint256.Int, uncles []*types.Header, blockReward *uint256.Int) *uint256.Int {
	r := uint256.NewInt(0)

	for range uncles {
		r.Add(r, getEraUncleBlockReward(era, blockReward)) // can reuse this, since 1/32 for winner's uncles remain unchanged from "Era 1"
	}
	return r
}

// GetBlockWinnerRewardByEra gets a block reward at disinflation rate.
// Constants MaxBlockReward, disinflationRateQuotient, and disinflationRateDivisor assumed.
func GetBlockWinnerRewardByEra(era *uint256.Int, blockReward *uint256.Int) *uint256.Int {
	if era.Cmp(uint256.NewInt(0)) == 0 {
		return new(uint256.Int).Set(blockReward)
	}

	// MaxBlockReward _r_ * (4/5)**era == MaxBlockReward * (4**era) / (5**era)
	// since (q/d)**n == q**n / d**n
	// qed
	var q, d, r = new(uint256.Int), new(uint256.Int), new(uint256.Int)

	q.Exp(disinflationRateQuotient, era)
	d.Exp(disinflationRateDivisor, era)

	r.Mul(blockReward, q)
	r.Div(r, d)

	return r
}

func ecip1017BlockReward(header *types.Header, uncles []*types.Header) (uint256.Int, []uint256.Int) {
	blockReward := FrontierBlockReward

	// Ensure value 'era' is configured.
	eraLen := ecip1017EraLen
	era := GetBlockEra(new(uint256.Int).SetUint64(header.Number.Uint64()), new(uint256.Int).Set(eraLen))
	wr := GetBlockWinnerRewardByEra(era, blockReward)                    // wr "winner reward". 5, 4, 3.2, 2.56, ...
	wurs := GetBlockWinnerRewardForUnclesByEra(era, uncles, blockReward) // wurs "winner uncle rewards"
	wr.Add(wr, wurs)

	// Reward uncle miners.
	uncleRewards := make([]uint256.Int, len(uncles))
	for i, uncle := range uncles {
		ur := GetBlockUncleRewardByEra(era, header, uncle, blockReward)
		uncleRewards[i] = *ur
	}

	return *wr, uncleRewards
}

// GetBlockEra gets which "Era" a given block is within, given an era length (ecip-1017 has era=5,000,000 blocks)
// Returns a zero-index era number, so "Era 1": 0, "Era 2": 1, "Era 3": 2 ...
func GetBlockEra(blockNum, eraLength *uint256.Int) *uint256.Int {
	// If genesis block or impossible negative-numbered block, return zero-val.
	if blockNum.Sign() < 1 {
		return new(uint256.Int)
	}

	remainder := uint256.NewInt(0).Mod(uint256.NewInt(0).Sub(blockNum, uint256.NewInt(1)), eraLength)
	base := uint256.NewInt(0).Sub(blockNum, remainder)

	d := uint256.NewInt(0).Div(base, eraLength)
	dremainder := uint256.NewInt(0).Mod(d, uint256.NewInt(1))

	return new(uint256.Int).Sub(d, dremainder)
}
