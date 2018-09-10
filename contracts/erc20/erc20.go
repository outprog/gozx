package erc20

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var erc20ABI abi.ABI

func init() {
	var err error
	erc20ABI, err = abi.JSON(strings.NewReader(Erc20ABI))
	if err != nil {
		panic(err)
	}
}

func Approve(spender common.Address, value *big.Int) ([]byte, error) {
	return erc20ABI.Pack("approve", spender, value)
}
