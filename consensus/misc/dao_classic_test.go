package misc

import (
	"errors"
	"testing"

	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/params"
)

func TestVerifyDAOHeaderExtraData(t *testing.T) {
	etcConfig := params.ClassicChainConfig

	// Verify that the canonical config is correct
	if etcConfig.DAOForkBlock == nil {
		t.Fatal("DAOForkBlock is nil")
	}
	if etcConfig.DAOForkSupport {
		t.Fatal("DAOForkSupport is true")
	}

	// First, the header will have Pro-Fork extra data, we'll test that this is rejected with the correct error
	header := &types.Header{
		Number: etcConfig.DAOForkBlock,
		Extra:  params.DAOForkBlockExtra,
	}
	if err := VerifyDAOHeaderExtraData(etcConfig, header); !errors.Is(err, ErrBadNoDAOExtra) {
		t.Fatal(err)
	}

	// Then unset the extra-data and verify that it's accepted (no error)
	header.Extra = nil
	if err := VerifyDAOHeaderExtraData(etcConfig, header); err != nil {
		t.Fatal(err)
	}
}
