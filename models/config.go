package models

import "github.com/ethereum/go-ethereum/common"

type Config struct {
	EthRPC string

	Erc20ProxyContractAddress  common.Address
	Erc721ProxyContractAddress common.Address
	ExchangeContractAddress    common.Address
	ForwarderContractAddress   common.Address
	GasPrice                   int64
	ZrxContractAddress         common.Address
}
