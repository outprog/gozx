package utils

import (
	"bytes"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func EncodeERC20AssetData(tokenAddr common.Address) (assetData string, err error) {
	erc20, err := genABI(ERC20ABI)
	if err != nil {
		return
	}

	res, err := erc20.Pack("ERC20Token", tokenAddr)
	if err != nil {
		return
	}

	assetData = common.ToHex(res)
	return
}

func DecodeERC20AssetData(assetData string) (tokenAddr common.Address, err error) {
	by := common.FromHex(assetData)

	if len(by) != 36 {
		err = errors.New("Wrong length of assetData")
		return
	}

	erc20, err := genABI(ERC20ABI)
	if err != nil {
		return
	}

	if bytes.Compare(by[:4], erc20.Methods["ERC20Token"].Id()) != 0 {
		err = errors.New("Wrong transfer proxy ID")
		return
	}

	tokenAddr = common.BytesToAddress(by[4:])
	return
}

func EncodeERC721AssetData() {}

func DecodeERC721AssetData() string {
	return ""
}

func genABI(sABI string) (abi.ABI, error) {
	tokenABI, err := abi.JSON(strings.NewReader(sABI))
	if err != nil {
		return tokenABI, err
	}
	return tokenABI, nil
}
