//go:build integration

package tests

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ledgerwatch/erigon-lib/kv/memdb"
	"github.com/ledgerwatch/erigon/core/vm"
	"github.com/ledgerwatch/log/v3"
)

var (
	baseDirClassic           = filepath.Join(".", "testdata-etc")
	stateTestDirClassic      = filepath.Join(baseDirClassic, "GeneralStateTests")
	difficultyTestDirClassic = filepath.Join(baseDirClassic, "DifficultyTests")
)

func TestState_Classic(t *testing.T) {
	defer log.Root().SetHandler(log.Root().GetHandler())
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlError, log.StderrHandler))
	if runtime.GOOS == "windows" {
		t.Skip("fix me on win please") // it's too slow on win, need generally improve speed of this tests
	}
	//t.Parallel()

	st := new(testMatcher)

	// Very time consuming
	st.skipLoad(`^stTimeConsuming/`)
	st.skipLoad(`.*vmPerformance/loop.*`)

	st.walk(t, stateTestDirClassic, func(t *testing.T, name string, test *StateTest) {
		db := memdb.NewTestDB(t)
		for _, subtest := range test.Subtests() {
			subtest := subtest
			key := fmt.Sprintf("%s/%d", subtest.Fork, subtest.Index)
			t.Run(key, func(t *testing.T) {
				withTrace(t, func(vmconfig vm.Config) error {
					tx, err := db.BeginRw(context.Background())
					if err != nil {
						t.Fatal(err)
					}
					defer tx.Rollback()
					_, err = test.Run(tx, subtest, vmconfig)
					tx.Rollback()
					if err != nil && len(test.json.Post[subtest.Fork][subtest.Index].ExpectException) > 0 {
						// Ignore expected errors
						return nil
					}
					return st.checkFailure(t, err)
				})
			})
		}
	})
}
