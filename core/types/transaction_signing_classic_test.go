package types

import (
	"math/big"
	"testing"

	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/params"
)

func TestMakeSigner_Classic(t *testing.T) {
	// Test number: 8772000 - 3 = 8771997; a little below Byzantium.
	testN := new(big.Int).Sub(params.ClassicChainConfig.SpuriousDragonBlock, common.Big3)
	signer := MakeSigner(params.ClassicChainConfig, testN.Uint64(), 0)
	if !signer.protected {
		t.Fatal("expected protected signer")
	}

	// Test number: a little below EIP155 activation for Classic.
	testN = big.NewInt(2_900_000)
	signer = MakeSigner(params.ClassicChainConfig, testN.Uint64(), 0)
	if signer.protected {
		t.Fatal("expected unprotected signer")
	}
}
