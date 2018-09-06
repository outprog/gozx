package models

import (
	"github.com/ethereum/go-ethereum/common"
)

type Order struct {
	ExchangeAddress       common.Address
	ExpirationTimeSeconds string
	FeeRecipientAddress   common.Address
	MakerAddress          common.Address
	MakerAssetAmount      string
	MakerAssetData        string
	MakerFee              string
	Salt                  string
	SenderAddress         common.Address
	TakerAddress          common.Address
	TakerAssetAmount      string
	TakerAssetData        string
	TakerFee              string
}
