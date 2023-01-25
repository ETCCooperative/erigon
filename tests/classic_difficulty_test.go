//go:build integration

package tests

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ledgerwatch/erigon/common/math"
)

type difficultyTestMarshalingProper struct {
	ParentTimestamp    math.HexOrDecimal64
	ParentDifficulty   *math.HexOrDecimal256
	CurrentTimestamp   math.HexOrDecimal64
	CurrentDifficulty  *math.HexOrDecimal256
	ParentUncles       math.HexOrDecimal64 `json:"parentUncles"`
	CurrentBlockNumber math.HexOrDecimal64
}

func TestDifficulty_Classic(t *testing.T) {
	t.Parallel()

	dt := new(testMatcher)

	dt.walk(t, difficultyTestDirClassic, func(t *testing.T, name string, superTest map[string]json.RawMessage) {
		for fork, rawTests := range superTest {
			if fork == "_info" {
				continue
			}
			var tests = make(map[string]DifficultyTest)
			var testsm map[string]difficultyTestMarshalingProper
			if err := json.Unmarshal(rawTests, &testsm); err != nil {
				t.Error(err, string(rawTests))
				continue
			}

			for k, v := range testsm {
				test := DifficultyTest{
					ParentTimestamp:    uint64(v.ParentTimestamp),
					ParentDifficulty:   (*big.Int)(v.ParentDifficulty),
					ParentUncles:       uint64(v.ParentUncles),
					CurrentTimestamp:   uint64(v.CurrentTimestamp),
					CurrentBlockNumber: uint64(v.CurrentBlockNumber),
					CurrentDifficulty:  (*big.Int)(v.CurrentDifficulty),
				}
				tests[k] = test
			}

			cfg, ok := Forks[fork]
			if !ok {
				t.Error(UnsupportedForkError{fork})
				continue
			}

			for subname, subtest := range tests {
				key := fmt.Sprintf("%s/%s", fork, subname)
				t.Run(key, func(t *testing.T) {
					if err := dt.checkFailure(t, subtest.Run(cfg)); err != nil {
						t.Error(err)
					}
				})
			}
		}
	})
}
