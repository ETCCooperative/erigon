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

package params

// ClassicChainConfig is the chain parameters to run a node on the main network.
var ClassicChainConfig = readChainSpec("chainspecs/classic.json")

// IsClassic returns true if the config's chain id is 61 (ETC).
func (c *ChainConfig) IsClassic() bool {
	if c.ChainID == nil {
		return false
	}
	return c.ChainID.Cmp(ClassicChainConfig.ChainID) == 0
}

// IsProtectedSigner returns true if the signer is protected from replay attacks with Chain ID.
// This feature was activated for the ETH network (et al) with the Spurious Dragon hard fork,
// while it was activated for the ETC network with a different hard fork, which differentially
// did not include EIP-161 State Trie Clearing.
func (c *ChainConfig) IsProtectedSigner(num uint64) bool {
	if c.IsClassic() {
		return isForked(c.ClassicEIP155Block, num)
	}
	return isForked(c.SpuriousDragonBlock, num)
}

// ClassicDNS is the Ethereum Classic DNS network identifier.
const ClassicDNS = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@all.classic.blockd.info"

// IsECIP1010 returns true if the block number is greater than or equal to the ECIP1010 block number.
// ECIP1010 disables the difficulty bomb for 2000000 blocks.
func (c *ChainConfig) IsECIP1010(num uint64) bool {
	return isForked(c.ECIP1010Block, num)
}

// IsECIP1010Disable returns true if the block number is greater than or equal to the ECIP1010Disable block number.
// This is the block number where the difficulty bomb is re-enabled after ECIP1010's pause runs out.
func (c *ChainConfig) IsECIP1010Disable(num uint64) bool {
	return isForked(c.ECIP1010DisableBlock, num)
}

// IsECIP1017 returns true if the block number is greater than or equal to the ECIP1017 block number.
// ECIP1017 defines the ETC monetary policy known as 5M20, which reduces the block reward by 20% every 5 million blocks.
func (c *ChainConfig) IsECIP1017(num uint64) bool {
	return isForked(c.ECIP1017Block, num)
}

// IsECIP1041 returns true if the block number is greater than or equal to the ECIP1041 block number.
// ECIP1041 removes the difficulty bomb.
func (c *ChainConfig) IsECIP1041(num uint64) bool {
	return isForked(c.ECIP1041Block, num)
}

// ECIP1099ForkBlockUint64 returns the ECIP1099ForkBlock as a uint64 pointer.
// If the ECIP1099ForkBlock is not defined, nil is returned.
// ECIP1099 defines an 'etchash' vs. ethash algorithm change, doubling the length of the DAG epoch.
func (c *ChainConfig) ECIP1099ForkBlockUint64() *uint64 {
	if c.ECIP1099Block == nil {
		return nil
	}
	n := c.ECIP1099Block.Uint64()
	return &n
}
