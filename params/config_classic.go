package params

var (
	// ClassicChainConfig is the chain parameters to run a node on the Ethereum Classic main network.
	ClassicChainConfig = readChainSpec("chainspecs/classic.json")
)
