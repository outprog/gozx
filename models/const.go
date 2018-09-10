package models

import "github.com/ethereum/go-ethereum/common"

const (
	GasLimit = uint64(400000)
)

var NullAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

// Config for kovan
var (
	KovanConfig = Config{
		EthRPC: "https://kovan.infura.io",

		Erc20ProxyContractAddress:  common.HexToAddress("0x9ad1b8209cea603892c9dfaa676bc737088b499a"),
		Erc721ProxyContractAddress: common.HexToAddress("0xbb428f7108971ecf1144700c4d37792f8b74f6b0"),
		ExchangeContractAddress:    common.HexToAddress("0xa458ec0709468996ef2ef668f5e52f37ceb66627"),

		WethAddress: KovanTokens["WETH"].Address,
	}

	// token list
	KovanTokens = map[string]Token{
		"WETH": Token{
			Symbol:   "WETH",
			Address:  common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c"),
			Decimals: 18,
			Type:     "ERC20",
		},
		"ZRX": Token{
			Symbol:   "ZRX",
			Address:  common.HexToAddress("0x6ff6c0ff1d68b964901f986d4c9fa3ac68346570"),
			Decimals: 18,
			Type:     "ERC20",
		},
	}
)
