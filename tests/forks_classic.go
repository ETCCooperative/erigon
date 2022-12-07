package tests

import (
	"math/big"

	"github.com/ledgerwatch/erigon/params"
)

func init() {
	Forks["ETC_Magneto"] = &params.ChainConfig{
		ChainID:                 big.NewInt(61),
		HomesteadBlock:          big.NewInt(0),
		TangerineWhistleBlock:   big.NewInt(0),
		SpuriousDragonBlock:     big.NewInt(0),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     big.NewInt(0),
		PetersburgBlock:         big.NewInt(0),
		IstanbulBlock:           big.NewInt(0),
		MuirGlacierBlock:        nil,
		BerlinBlock:             nil,
		LondonBlock:             nil,
		ArrowGlacierBlock:       nil,
		TerminalTotalDifficulty: nil,
	}
	Forks["ETC_Mystique"] = &params.ChainConfig{
		ChainID:                 big.NewInt(61),
		HomesteadBlock:          big.NewInt(0),
		TangerineWhistleBlock:   big.NewInt(0),
		SpuriousDragonBlock:     big.NewInt(0),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     big.NewInt(0),
		PetersburgBlock:         big.NewInt(0),
		IstanbulBlock:           big.NewInt(0),
		MuirGlacierBlock:        nil,
		BerlinBlock:             big.NewInt(0),
		LondonBlock:             nil,
		ArrowGlacierBlock:       nil,
		TerminalTotalDifficulty: nil,
	}
}
