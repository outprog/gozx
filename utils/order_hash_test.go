package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/outprog/gozx/models"
)

func TestGetOrderHashHex(t *testing.T) {
	orderHash, err := GetOrderHashHex(&models.Order{
		ExchangeAddress:       common.HexToAddress("0x35dd2932454449b14cee11a94d3674a936d5d7b2"),
		MakerAddress:          common.HexToAddress("0x5409ed021d9299bf6814279a6a1411a7e866a631"),
		FeeRecipientAddress:   models.NullAddress,
		TakerAddress:          models.NullAddress,
		SenderAddress:         models.NullAddress,
		ExpirationTimeSeconds: "1536225890",
		MakerAssetAmount:      "5000000000000000000",
		TakerAssetAmount:      "100000000000000000",
		MakerFee:              "0",
		TakerFee:              "0",
		Salt:                  "995194064667745912528683602854194084472071061072357932520081785590044667663",
		MakerAssetData:        "0xf47261b00000000000000000000000002002d3812f58e35f0ea1ffbf80a75a38c32175fa",
		TakerAssetData:        "0xf47261b0000000000000000000000000d0a1e359811322d97991e03f863a0c30c2cf029c",
	})
	require.NoError(t, err)
	assert.Equal(t, "0x24d79c3d385e43f90ca9fdf630d64627b18dfd5fc394d4b501b3fd8fea3e941b", orderHash)
}

func TestHashOrder(t *testing.T) {
	hash, err := hashOrder(&models.Order{
		ExchangeAddress:       common.HexToAddress("0x35dd2932454449b14cee11a94d3674a936d5d7b2"),
		MakerAddress:          common.HexToAddress("0x5409ed021d9299bf6814279a6a1411a7e866a631"),
		FeeRecipientAddress:   models.NullAddress,
		TakerAddress:          models.NullAddress,
		SenderAddress:         models.NullAddress,
		ExpirationTimeSeconds: "1536225890",
		MakerAssetAmount:      "5000000000000000000",
		TakerAssetAmount:      "100000000000000000",
		MakerFee:              "0",
		TakerFee:              "0",
		Salt:                  "995194064667745912528683602854194084472071061072357932520081785590044667663",
		MakerAssetData:        "0xf47261b00000000000000000000000002002d3812f58e35f0ea1ffbf80a75a38c32175fa",
		TakerAssetData:        "0xf47261b0000000000000000000000000d0a1e359811322d97991e03f863a0c30c2cf029c",
	})
	require.NoError(t, err)
	assert.Equal(t, common.FromHex("0d1727a1ed3d4e33fc020ece757ad6bdcfae962d135f1596b9d5a98baad5b2a5"), hash)
}
