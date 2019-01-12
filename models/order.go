package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Order struct {
	ExchangeAddress       common.Address
	ExpirationTimeSeconds *big.Int
	FeeRecipientAddress   common.Address
	MakerAddress          common.Address
	MakerAssetAmount      *big.Int
	MakerAssetData        []byte
	MakerFee              *big.Int
	Salt                  *big.Int
	SenderAddress         common.Address
	TakerAddress          common.Address
	TakerAssetAmount      *big.Int
	TakerAssetData        []byte
	TakerFee              *big.Int
}

func (o *Order) String() string {
	return "exchangeAddress: " + o.ExchangeAddress.String() +
		"\nexpirationTimeSeconds: " + o.ExpirationTimeSeconds.String() +
		"\nFeeRecipientAddress: " + o.FeeRecipientAddress.String() +
		"\nMakerAddress: " + o.MakerAddress.String() +
		"\nMakerAssetAmount: " + o.MakerAssetAmount.String() +
		"\nMakerAssetData: " + common.Bytes2Hex(o.MakerAssetData) +
		"\nMakerFee: " + o.MakerFee.String() +
		"\nSalt: " + o.Salt.String() +
		"\nSenderAddress: " + o.SenderAddress.String() +
		"\nTakerAddress: " + o.TakerAddress.String() +
		"\nTakerAssetAmount: " + o.TakerAssetAmount.String() +
		"\nTakerAssetData: " + common.Bytes2Hex(o.TakerAssetData) +
		"\nTakerFee: " + o.TakerFee.String()
}
