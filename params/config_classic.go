package params

import (
	"math/big"

	"github.com/ledgerwatch/erigon/common/math"
)

// ClassicChainConfig is the chain parameters to run a node on the main network.
var ClassicChainConfig = readChainSpec("chainspecs/classic.json")

// IsClassic returns true if the config's chain id is 61 (ETC).
func (c *ChainConfig) IsClassic() bool {
	if c.ChainID == nil {
		return false
	}
	return c.ChainID.Cmp(ClassicChainConfig.ChainID) == 0
}

// ECIP1010Block_Classic defines the block number where the ECIP-1010 difficulty bomb delay is activated,
// delaying the bomb for 2M blocks.
var ECIP1010Block_Classic = big.NewInt(3_000_000)

func (c *ChainConfig) ECIP1010Block() *big.Int {
	if c.IsClassic() {
		return ECIP1010Block_Classic
	}
	return nil
}

// ECIP1041Block_Classic is the ultimate difficulty bomb diffuser block number for the Ethereum Classic network.
var ECIP1041Block_Classic = big.NewInt(5_900_000)

func (c *ChainConfig) ECIP1041Block() *big.Int {
	if c.IsClassic() {
		return ECIP1041Block_Classic
	}
	return nil
}

// ECIP1017Block_Classic defines the block number where the ECIP-1017 monetary policy is activated,
var ECIP1017Block_Classic = big.NewInt(5_000_000)
var infinity = big.NewInt(math.MaxInt64)

func (c *ChainConfig) ECIP1017Block() *big.Int {
	if c.IsClassic() {
		return ECIP1017Block_Classic
	}
	return infinity
}

// ECIP1099Block_Classic defines the block number where the ECIP-1099 Etchash PoW algorithm is activated.
var ECIP1099Block_Classic = big.NewInt(11_700_000)

var (
	classicEIP155Block   = big.NewInt(3_000_000)
	classicEIP160        = big.NewInt(3_000_000)
	classicMystiqueBlock = big.NewInt(14_525_000)
)

func (c *ChainConfig) IsProtectedSigner(num uint64) bool {
	if c.IsClassic() {
		return isForked(classicEIP155Block, num)
	}
	return isForked(c.SpuriousDragonBlock, num)
}

const ClassicDNS = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@all.classic.blockd.info"
