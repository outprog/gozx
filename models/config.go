package models

type Config struct {
	EthRPC string

	Erc20ProxyContractAddress  string
	Erc721ProxyContractAddress string
	ExchangeContractAddress    string
	ForwarderContractAddress   string
	GasPrice                   int64
	ZrxContractAddress         string
}
