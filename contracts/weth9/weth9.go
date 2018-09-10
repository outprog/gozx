package weth9

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var weth9ABI abi.ABI

func init() {
	var err error
	weth9ABI, err = abi.JSON(strings.NewReader(Weth9ABI))
	if err != nil {
		panic(err)
	}
}

func Deposit() ([]byte, error) {
	return weth9ABI.Pack("deposit")
}

func Withdraw(wad *big.Int) ([]byte, error) {
	return weth9ABI.Pack("withdraw", wad)
}
