package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/outprog/gozx/models"
)

func GetOrderHash(order *models.Order) []byte {
	domainHash := crypto.Keccak256(EIP712_DOMAIN_SEPARATOR_SCHEMA_HASH,
		crypto.Keccak256([]byte("0x Protocol")),
		crypto.Keccak256([]byte("2")),
		addressToByte32(order.ExchangeAddress),
	)
	return crypto.Keccak256([]byte(EIP191_HEADER),
		domainHash,
		hashOrder(order))
}

func IsValidOrderHash(order *models.Order) bool {
	return true
}

func hashOrder(order *models.Order) []byte {
	return crypto.Keccak256(EIP712_ORDER_SCHEMA_HASH,
		addressToByte32(order.MakerAddress),
		addressToByte32(order.TakerAddress),
		addressToByte32(order.FeeRecipientAddress),
		addressToByte32(order.SenderAddress),
		math.PaddedBigBytes(order.MakerAssetAmount, 32),
		math.PaddedBigBytes(order.TakerAssetAmount, 32),
		math.PaddedBigBytes(order.MakerFee, 32),
		math.PaddedBigBytes(order.TakerFee, 32),
		math.PaddedBigBytes(order.ExpirationTimeSeconds, 32),
		math.PaddedBigBytes(order.Salt, 32),
		crypto.Keccak256(order.MakerAssetData),
		crypto.Keccak256(order.TakerAssetData),
	)
}

func addressToByte32(addr common.Address) []byte {
	return append(make([]byte, 12), addr.Bytes()...)
}
