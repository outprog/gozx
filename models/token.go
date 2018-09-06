package models

import "github.com/ethereum/go-ethereum/common"

type Token struct {
	Symbol   string
	Name     string
	Address  common.Address
	Decimals int64
	Type     string
}
