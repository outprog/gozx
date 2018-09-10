package exchange

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var exchangeABI abi.ABI

// func init() {
// 	var err error
// 	exchangeABI, err = abi.JSON(strings.NewReader(ExchangeABI))
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func FillOrder(order *models.Order, takerAssetFillAmount, signatrue string) ([]byte, error) {
// 	return exchangeABI.Pack("fillOrder", []byte{}, takerAssetFillAmount, signatrue)
// }
