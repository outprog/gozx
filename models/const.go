package models

// Config for mainnet or kovan
var (
	KovanConfig = Config{
		EthRPC: "https://kovan.infura.io",
		Erc20ProxyContractAddress:  "0x9ad1b8209cea603892c9dfaa676bc737088b499a",
		Erc721ProxyContractAddress: "0xbb428f7108971ecf1144700c4d37792f8b74f6b0",
		ExchangeContractAddress:    "0xa458ec0709468996ef2ef668f5e52f37ceb66627",
	}
)
