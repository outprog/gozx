package utils

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeERC20AssetData(t *testing.T) {
	address := common.HexToAddress("0x1dc4c1cefef38a777b15aa20260a54e584b16c48")
	assetData, err := EncodeERC20AssetData(address)
	require.NoError(t, err)

	assert.Equal(t, "0xf47261b00000000000000000000000001dc4c1cefef38a777b15aa20260a54e584b16c48", assetData)
}

func TestDecodeERC20AssetData(t *testing.T) {
	_, err := DecodeERC20AssetData("0x123")
	assert.Equal(t, "Wrong length of assetData", err.Error())

	_, err = DecodeERC20AssetData("0x08e937fa0000000000000000000000001dc4c1cefef38a777b15aa20260a54e584b16c48")
	assert.Equal(t, "Wrong transfer proxy ID", err.Error())

	tokenAddr, err := DecodeERC20AssetData("0xf47261b00000000000000000000000001dc4c1cefef38a777b15aa20260a54e584b16c48")
	require.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x1dc4c1cefef38a777b15aa20260a54e584b16c48"), tokenAddr)
}

func TestEncodeERC721AssetData(t *testing.T) {
}

func TestDecodeERC721AssetData(t *testing.T) {
}

func TestGenABI(t *testing.T) {
	_, err := genABI(ERC20_ABI)
	assert.NoError(t, err)
}
