// Copyright 2022 The erigon Authors
// This file is part of the erigon library.
//
// The erigon library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The erigon library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"math/big"

	"github.com/ledgerwatch/erigon-lib/chain"
	"github.com/ledgerwatch/erigon/params"
)

func init() {
	Forks["ETC_Atlantis"] = &chain.Config{
		ChainID:                 big.NewInt(61),
		HomesteadBlock:          big.NewInt(0),
		TangerineWhistleBlock:   big.NewInt(0),
		SpuriousDragonBlock:     big.NewInt(0),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     nil,
		PetersburgBlock:         nil,
		IstanbulBlock:           nil,
		MuirGlacierBlock:        nil,
		BerlinBlock:             nil,
		LondonBlock:             nil,
		ArrowGlacierBlock:       nil,
		TerminalTotalDifficulty: nil,
		ECIP1010Block:           big.NewInt(0),
		ECIP1010DisableBlock:    big.NewInt(0),
		ECIP1017Block:           big.NewInt(0),
		ECIP1041Block:           big.NewInt(0),
		ECIP1099Block:           nil,
		ClassicEIP155Block:      big.NewInt(0),
		ClassicEIP160Block:      big.NewInt(0),
		ClassicMystiqueBlock:    nil,
	}
	Forks["ETC_Agharta"] = &chain.Config{
		ChainID:                 big.NewInt(61),
		HomesteadBlock:          big.NewInt(0),
		TangerineWhistleBlock:   big.NewInt(0),
		SpuriousDragonBlock:     big.NewInt(0),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     big.NewInt(0),
		PetersburgBlock:         big.NewInt(0),
		IstanbulBlock:           nil,
		MuirGlacierBlock:        nil,
		BerlinBlock:             nil,
		LondonBlock:             nil,
		ArrowGlacierBlock:       nil,
		TerminalTotalDifficulty: nil,
		ECIP1010Block:           big.NewInt(0),
		ECIP1010DisableBlock:    big.NewInt(0),
		ECIP1017Block:           big.NewInt(0),
		ECIP1041Block:           big.NewInt(0),
		ECIP1099Block:           nil,
		ClassicEIP155Block:      big.NewInt(0),
		ClassicEIP160Block:      big.NewInt(0),
		ClassicMystiqueBlock:    nil,
	}
	Forks["ETC_Phoenix"] = &chain.Config{
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
		ECIP1010Block:           big.NewInt(0),
		ECIP1010DisableBlock:    big.NewInt(0),
		ECIP1017Block:           big.NewInt(0),
		ECIP1041Block:           big.NewInt(0),
		ECIP1099Block:           nil,
		ClassicEIP155Block:      big.NewInt(0),
		ClassicEIP160Block:      big.NewInt(0),
		ClassicMystiqueBlock:    nil,
	}
	Forks["ETC_Magneto"] = &chain.Config{
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
		ECIP1010Block:           big.NewInt(0),
		ECIP1010DisableBlock:    big.NewInt(0),
		ECIP1017Block:           big.NewInt(0),
		ECIP1041Block:           big.NewInt(0),
		ECIP1099Block:           big.NewInt(0),
		ClassicEIP155Block:      big.NewInt(0),
		ClassicEIP160Block:      big.NewInt(0),
		ClassicMystiqueBlock:    nil,
	}
	Forks["ETC_Mystique"] = &chain.Config{
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
		ECIP1010Block:           big.NewInt(0),
		ECIP1010DisableBlock:    big.NewInt(0),
		ECIP1017Block:           big.NewInt(0),
		ECIP1041Block:           big.NewInt(0),
		ECIP1099Block:           big.NewInt(0),
		ClassicEIP155Block:      big.NewInt(0),
		ClassicEIP160Block:      big.NewInt(0),
		ClassicMystiqueBlock:    big.NewInt(0),
	}
	// Forks:ETC is a configuration used exclusively by the Difficulty tests.
	Forks["ETC"] = params.ClassicChainConfig
}
