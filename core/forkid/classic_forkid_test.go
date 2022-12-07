package forkid

import (
	"testing"

	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/params"
)

// TestCreation_Classic tests that different genesis and fork rule combinations result in
// the correct fork ID.
func TestCreation_Classic(t *testing.T) {
	type testcase struct {
		head uint64
		want ID
	}
	tests := []struct {
		config  *params.ChainConfig
		genesis common.Hash
		cases   []testcase
	}{
		// Classic test cases
		{
			params.ClassicChainConfig,
			params.MainnetGenesisHash,
			[]testcase{
				{0, ID{Hash: checksumToBytes(0xfc64ec04), Next: 1150000}},
				{1, ID{Hash: checksumToBytes(0xfc64ec04), Next: 1150000}},
				{2, ID{Hash: checksumToBytes(0xfc64ec04), Next: 1150000}},
				{3, ID{Hash: checksumToBytes(0xfc64ec04), Next: 1150000}},
				{9, ID{Hash: checksumToBytes(0xfc64ec04), Next: 1150000}},
				{10, ID{Hash: checksumToBytes(0xfc64ec04), Next: 1150000}},
				{1149999, ID{Hash: checksumToBytes(0xfc64ec04), Next: 1150000}},
				{1150000, ID{Hash: checksumToBytes(0x97c2c34c), Next: 2500000}},
				{1150001, ID{Hash: checksumToBytes(0x97c2c34c), Next: 2500000}},
				{2499999, ID{Hash: checksumToBytes(0x97c2c34c), Next: 2500000}},
				{2500000, ID{Hash: checksumToBytes(0xdb06803f), Next: 3000000}},
				{2500001, ID{Hash: checksumToBytes(0xdb06803f), Next: 3000000}},
				{2999999, ID{Hash: checksumToBytes(0xdb06803f), Next: 3000000}},
				{3000000, ID{Hash: checksumToBytes(0xaff4bed4), Next: 5000000}},
				{3000001, ID{Hash: checksumToBytes(0xaff4bed4), Next: 5000000}},
				{4999999, ID{Hash: checksumToBytes(0xaff4bed4), Next: 5000000}},
				{5000000, ID{Hash: checksumToBytes(0xf79a63c0), Next: 5900000}},
				{5000001, ID{Hash: checksumToBytes(0xf79a63c0), Next: 5900000}},
				{5899999, ID{Hash: checksumToBytes(0xf79a63c0), Next: 5900000}},
				{5900000, ID{Hash: checksumToBytes(0x744899d6), Next: 8772000}},
				{5900001, ID{Hash: checksumToBytes(0x744899d6), Next: 8772000}},
				{8771999, ID{Hash: checksumToBytes(0x744899d6), Next: 8772000}},
				{8772000, ID{Hash: checksumToBytes(0x518b59c6), Next: 9573000}},
				{8772001, ID{Hash: checksumToBytes(0x518b59c6), Next: 9573000}},
				{9572999, ID{Hash: checksumToBytes(0x518b59c6), Next: 9573000}},
				{9573000, ID{Hash: checksumToBytes(0x7ba22882), Next: 10500839}},
				{9573001, ID{Hash: checksumToBytes(0x7ba22882), Next: 10500839}},
				{10500838, ID{Hash: checksumToBytes(0x7ba22882), Next: 10500839}},
				{10500839, ID{Hash: checksumToBytes(0x9007bfcc), Next: 11_700_000}},
				{10500840, ID{Hash: checksumToBytes(0x9007bfcc), Next: 11_700_000}},
				{11_699_999, ID{Hash: checksumToBytes(0x9007bfcc), Next: 11_700_000}},
				{11_700_000, ID{Hash: checksumToBytes(0xdb63a1ca), Next: 13_189_133}},
				{11_700_001, ID{Hash: checksumToBytes(0xdb63a1ca), Next: 13_189_133}},
				{13_189_132, ID{Hash: checksumToBytes(0xdb63a1ca), Next: 13_189_133}},
				{13_189_133, ID{Hash: checksumToBytes(0x0f6bf187), Next: 14_525_000}},
				{13_189_134, ID{Hash: checksumToBytes(0x0f6bf187), Next: 14_525_000}},
				{14_524_999, ID{Hash: checksumToBytes(0x0f6bf187), Next: 14_525_000}},
				{14_525_000, ID{Hash: checksumToBytes(0x7fd1bb25), Next: 0}},
				{14_525_001, ID{Hash: checksumToBytes(0x7fd1bb25), Next: 0}},
			},
		},
	}
	for i, tt := range tests {
		for j, ttt := range tt.cases {
			if have := NewID(tt.config, tt.genesis, ttt.head); have != ttt.want {
				t.Errorf("test %d, case %d: fork ID mismatch: have %x, want %x", i, j, have, ttt.want)
			}
		}
	}
}

func TestGatherForks(t *testing.T) {
	cases := []struct {
		name   string
		config *params.ChainConfig
		wantNs []uint64
	}{
		{
			"classic",
			params.ClassicChainConfig,
			[]uint64{1150000, 2500000, 3000000, 5000000, 5900000, 8772000, 9573000, 10500839, 11_700_000, 13_189_133, 14_525_000},
		},
		{
			"mainnet",
			params.MainnetChainConfig,
			[]uint64{1150000, 1920000, 2463000, 2675000, 4370000, 7280000, 9069000, 9200000, 12_244_000, 12_965_000, 13_773_000, 15050000},
		},
	}
	sliceContains := func(sl []uint64, u uint64) bool {
		for _, s := range sl {
			if s == u {
				return true
			}
		}
		return false
	}
	for _, c := range cases {
		gotForkNs := GatherForks(c.config)
		if len(gotForkNs) != len(c.wantNs) {
			for _, n := range c.wantNs {
				if !sliceContains(gotForkNs, n) {
					t.Errorf("config=%s missing wanted fork at block number: %d", c.name, n)
				}
			}
			for _, n := range gotForkNs {
				if !sliceContains(c.wantNs, n) {
					t.Errorf("config=%s gathered unwanted fork at block number: %d", c.name, n)
				}
			}
		}
	}
}
