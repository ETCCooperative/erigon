package core

import (
	"math/big"

	"github.com/ledgerwatch/erigon-lib/common/hexutil"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/params"
)

// ClassicGenesisBlock returns the Ethereum main net genesis block.
func ClassicGenesisBlock() *types.Genesis {
	return &types.Genesis{
		Config:     params.ClassicChainConfig,
		Nonce:      66,
		ExtraData:  hexutil.MustDecode("0x11bbe8db4e347b4e8c937c1c8370e4b5ed33adb3db69cbdb7a38e1e50b1b82fa"),
		GasLimit:   5000,
		Difficulty: big.NewInt(17179869184),
		Alloc:      readPrealloc("allocs/mainnet.json"),
	}
}
