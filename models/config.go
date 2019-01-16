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

		Erc20ProxyContractAddress:  common.HexToAddress("0xf1ec01d6236d3cd881a0bf0130ea25fe4234003e"),
		Erc721ProxyContractAddress: common.HexToAddress("0x2a9127c745688a165106c11cd4d647d2220af821"),
		ExchangeContractAddress:    common.HexToAddress("0x35dd2932454449b14cee11a94d3674a936d5d7b2"),

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
