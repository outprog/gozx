package models

import (
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	EthRPC string

	Erc20ProxyContractAddress  common.Address
	Erc721ProxyContractAddress common.Address
	ExchangeContractAddress    common.Address
	ForwarderContractAddress   common.Address
	GasPrice                   int64
	ZrxContractAddress         common.Address

	WethAddress common.Address
}

const (
	GasLimit = uint64(400000)
)

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
		"CZT": Token{
			Symbol:   "CZT",
			Address:  common.HexToAddress("0xa7c8cd567a2321ed6c31395de25bcbcc4ebfce17"),
			Decimals: 22,
			Type:     "ERC20",
		},
	}
)
