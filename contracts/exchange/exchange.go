package exchange

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/outprog/gozx/models"
)

var exchangeABI abi.ABI

func init() {
	var err error
	exchangeABI, err = abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		panic(err)
	}
}

func FillOrder(order *models.Order, takerAssetFillAmount *big.Int, signatrue []byte) ([]byte, error) {
	data, err := exchangeABI.Pack("fillOrder", order, takerAssetFillAmount, signatrue)
	if err != nil {
		return nil, err
	}

	return data, nil
}
