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
