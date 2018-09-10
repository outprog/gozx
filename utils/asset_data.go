package utils

import (
	"bytes"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var erc20ABI abi.ABI

func init() {
	var err error
	erc20ABI, err = abi.JSON(strings.NewReader(ERC20_ABI))
	if err != nil {
		panic(err)
	}
}

func EncodeERC20AssetData(tokenAddr common.Address) ([]byte, error) {
	return erc20ABI.Pack("ERC20Token", tokenAddr)
}

func DecodeERC20AssetData(assetData []byte) (tokenAddr common.Address, err error) {
	if len(assetData) != 36 {
		err = errors.New("Wrong length of assetData")
		return
	}

	if bytes.Compare(assetData[:4], erc20ABI.Methods["ERC20Token"].Id()) != 0 {
		err = errors.New("Wrong transfer proxy ID")
		return
	}

	tokenAddr = common.BytesToAddress(assetData[4:])
	return
}

func EncodeERC721AssetData() {}

func DecodeERC721AssetData() string {
	return ""
}
