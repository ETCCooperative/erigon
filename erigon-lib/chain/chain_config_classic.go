package chain

import (
	"fmt"
	"math/big"
	"strings"
)

var classicChainID = big.NewInt(61)

// IsClassic returns true if the config's chain id is 61 (ETC).
func (c *Config) IsClassic() bool {
	return c.ChainID != nil && c.ChainID.Cmp(classicChainID) == 0
}

// IsProtectedSigner returns true if the signer is protected from replay attacks with Chain ID.
// This feature was activated for the ETH network (et al) with the Spurious Dragon hard fork,
// while it was activated for the ETC network with a different hard fork, which differentially
// did not include EIP-161 State Trie Clearing.
func (c *Config) IsProtectedSigner(num uint64) bool {
	if c.IsClassic() {
		return isForked(c.ClassicEIP155Block, num)
	}
	return isForked(c.SpuriousDragonBlock, num)
}

// IsECIP1010 returns true if the block number is greater than or equal to the ECIP1010 block number.
// ECIP1010 disables the difficulty bomb for 2000000 blocks.
func (c *Config) IsECIP1010(num uint64) bool {
	return isForked(c.ECIP1010Block, num)
}

// IsECIP1010Disable returns true if the block number is greater than or equal to the ECIP1010Disable block number.
// This is the block number where the difficulty bomb is re-enabled after ECIP1010's pause runs out.
func (c *Config) IsECIP1010Disable(num uint64) bool {
	return isForked(c.ECIP1010DisableBlock, num)
}

// IsECIP1017 returns true if the block number is greater than or equal to the ECIP1017 block number.
// ECIP1017 defines the ETC monetary policy known as 5M20, which reduces the block reward by 20% every 5 million blocks.
func (c *Config) IsECIP1017(num uint64) bool {
	return isForked(c.ECIP1017Block, num)
}

// IsECIP1041 returns true if the block number is greater than or equal to the ECIP1041 block number.
// ECIP1041 removes the difficulty bomb.
func (c *Config) IsECIP1041(num uint64) bool {
	return isForked(c.ECIP1041Block, num)
}

// ECIP1099ForkBlockUint64 returns the ECIP1099ForkBlock as a uint64 pointer.
// If the ECIP1099ForkBlock is not defined, nil is returned.
// ECIP1099 defines an 'etchash' vs. ethash algorithm change, doubling the length of the DAG epoch.
func (c *Config) ECIP1099ForkBlockUint64() *uint64 {
	if c.ECIP1099Block == nil {
		return nil
	}
	n := c.ECIP1099Block.Uint64()
	return &n
}

func ConfigString_Classic(c *Config) string {
	if c == nil {
		return "nil"
	}

	entries := []struct {
		name  string
		value *big.Int
	}{
		{"ECIP1010Block", c.ECIP1010Block},
		{"ECIP1010DisableBlock", c.ECIP1010DisableBlock},
		{"ECIP1017Block", c.ECIP1017Block},
		{"ECIP1041Block", c.ECIP1041Block},
		{"ECIP1099Block", c.ECIP1099Block},
		{"ClassicEIP155Block", c.ClassicEIP155Block},
		{"ClassicEIP160Block", c.ClassicEIP160Block},
		{"ClassicMystiqueBlock", c.ClassicMystiqueBlock},
	}

	// entryStrings become "name: value" for each entry.
	entryStrings := make([]string, 0, len(entries))
	for _, entry := range entries {
		entryStrings = append(entryStrings, fmt.Sprintf("%s: %v", entry.name, entry.value))
	}
	// entryStrings are joined by a comma and space, then inserted in the config string's curly braces.
	custom := strings.Join(entryStrings, ", ")

	engine := c.getEngine()
	return fmt.Sprintf("{ChainID: %v, Homestead: %v, DAO: %v, DAO Support: %v, Tangerine Whistle: %v, Spurious Dragon: %v, Byzantium: %v, Constantinople: %v, Petersburg: %v, Istanbul: %v, Muir Glacier: %v, Berlin: %v, London: %v, Arrow Glacier: %v, Gray Glacier: %v, Terminal Total Difficulty: %v, Merge Netsplit: %v, Shanghai: %v, Cancun: %v, Engine: %v%s}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.TangerineWhistleBlock,
		c.SpuriousDragonBlock,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		c.IstanbulBlock,
		c.MuirGlacierBlock,
		c.BerlinBlock,
		c.LondonBlock,
		c.ArrowGlacierBlock,
		c.GrayGlacierBlock,
		c.TerminalTotalDifficulty,
		c.MergeNetsplitBlock,
		c.ShanghaiTime,
		c.CancunTime,
		engine,
		custom,
	)
}
