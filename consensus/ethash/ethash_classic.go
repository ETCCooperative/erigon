package ethash

import (
	"encoding/binary"

	"github.com/ledgerwatch/erigon/common/hexutil"
	"github.com/ledgerwatch/erigon/crypto"
)

// uint32Array2ByteArray returns the bytes represented by uint32 array c
func uint32Array2ByteArray(c []uint32) []byte {
	buf := make([]byte, len(c)*4)
	if isLittleEndian() {
		for i, v := range c {
			binary.LittleEndian.PutUint32(buf[i*4:], v)
		}
	} else {
		for i, v := range c {
			binary.BigEndian.PutUint32(buf[i*4:], v)
		}
	}
	return buf
}

// bytes2Keccak256 returns the keccak256 hash as a hex string (0x prefixed)
// for a given uint32 array (cache/dataset)
func uint32Array2Keccak256(data []uint32) string {
	// convert to bytes
	bytes := uint32Array2ByteArray(data)
	// hash with keccak256
	digest := crypto.Keccak256(bytes)
	// return hex string
	return hexutil.Encode(digest)
}
