package utils

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/outprog/gozx/models"
)

func GetOrderHash(order *models.Order) (string, error) {
	hash, err := hashOrder(order)
	if err != nil {
		return "", err
	}

	domainHash := crypto.Keccak256(EIP712_DOMAIN_SEPARATOR_SCHEMA_HASH,
		crypto.Keccak256([]byte("0x Protocol")),
		crypto.Keccak256([]byte("2")),
		addressToByte32(order.ExchangeAddress),
	)
	orderHash := crypto.Keccak256([]byte(EIP191_HEADER),
		domainHash,
		hash)
	return common.ToHex(orderHash), nil
}

func IsValidOrderHash(order *models.Order) bool {
	return true
}

func hashOrder(order *models.Order) ([]byte, error) {

	byMakerAssetAmount, err := bigNumberToByte32(order.MakerAssetAmount)
	if err != nil {
		return []byte{}, err
	}
	byTakerAssetAmount, err := bigNumberToByte32(order.TakerAssetAmount)
	if err != nil {
		return []byte{}, err
	}
	byMakerFee, err := bigNumberToByte32(order.MakerFee)
	if err != nil {
		return []byte{}, err
	}
	byTakerFee, err := bigNumberToByte32(order.TakerFee)
	if err != nil {
		return []byte{}, err
	}
	byExpirationTimeSeconds, err := bigNumberToByte32(order.ExpirationTimeSeconds)
	if err != nil {
		return []byte{}, err
	}
	bySalt, err := bigNumberToByte32(order.Salt)
	if err != nil {
		return []byte{}, err
	}

	return crypto.Keccak256(EIP712_ORDER_SCHEMA_HASH,
		addressToByte32(order.MakerAddress),
		addressToByte32(order.TakerAddress),
		addressToByte32(order.FeeRecipientAddress),
		addressToByte32(order.SenderAddress),
		byMakerAssetAmount,
		byTakerAssetAmount,
		byMakerFee,
		byTakerFee,
		byExpirationTimeSeconds,
		bySalt,
		crypto.Keccak256(common.FromHex(order.MakerAssetData)),
		crypto.Keccak256(common.FromHex(order.TakerAssetData)),
	), nil
}

func addressToByte32(addr common.Address) []byte {
	return append(make([]byte, 12), addr.Bytes()...)
}

func bigNumberToByte32(big string) ([]byte, error) {
	i, ok := math.ParseBig256(big)
	if ok == false {
		err := errors.New("Order's Number type attributes Parse Failed")
		return []byte{}, err
	}
	return math.PaddedBigBytes(i, 32), nil
}
