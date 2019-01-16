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

func FillOrder(order models.Order, takerAssetFillAmount *big.Int, signature []byte) ([]byte, error) {
	data, err := exchangeABI.Pack("fillOrder", order, takerAssetFillAmount, signature)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func FillOrderNoThrow(order models.Order, takerAssetFillAmount *big.Int, signature []byte) ([]byte, error) {
	data, err := exchangeABI.Pack("fillOrderNoThrow", order, takerAssetFillAmount, signature)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func BatchFillOrders(orders []models.Order, takerAssetFillAmounts []*big.Int, signatures [][]byte) ([]byte, error) {
	data, err := exchangeABI.Pack("batchFillOrders", orders, takerAssetFillAmounts, signatures)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func BatchFillOrdersNoThrow(orders []models.Order, takerAssetFillAmounts []*big.Int, signatures [][]byte) ([]byte, error) {
	data, err := exchangeABI.Pack("batchFillOrdersNoThrow", orders, takerAssetFillAmounts, signatures)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CancelOrder(order models.Order) ([]byte, error) {
	data, err := exchangeABI.Pack("cancelOrder", order)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CancelOrdersUpTo(targetOrderEpoch *big.Int) ([]byte, error) {
	data, err := exchangeABI.Pack("cancelOrdersUpTo", targetOrderEpoch)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func BatchCancelOrders(orders []models.Order) ([]byte, error) {
	data, err := exchangeABI.Pack("batchCancelOrders", orders)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func MatchOrders(leftOrder, rightOrder models.Order, leftSignature, rightSignature []byte) ([]byte, error) {
	data, err := exchangeABI.Pack("matchOrders", leftOrder, rightOrder, leftSignature, rightSignature)
	if err != nil {
		return nil, err
	}

	return data, nil
}
