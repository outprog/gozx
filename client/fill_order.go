package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/outprog/gozx/contracts/exchange"
	"github.com/outprog/gozx/models"
)

func (c *Client) FillOrder(order models.Order, takerAssetFillAmount *big.Int, signature []byte, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	input, err := exchange.FillOrder(order, takerAssetFillAmount, signature)
	if err != nil {
		return common.Hash{}, err
	}

	return c.exchangeSender(input, gasPrice, nonce)
}

func (c *Client) FillOrderNoThrow(order models.Order, takerAssetFillAmount *big.Int, signature []byte, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	input, err := exchange.FillOrderNoThrow(order, takerAssetFillAmount, signature)
	if err != nil {
		return common.Hash{}, err
	}

	return c.exchangeSender(input, gasPrice, nonce)
}

func (c *Client) BatchFillOrders(orders []models.Order, takerAssetFillAmounts []*big.Int, signatures [][]byte, gasPrice *big.Int, nonce uint64) (common.Hash, error) {
	input, err := exchange.BatchFillOrders(orders, takerAssetFillAmounts, signatures)
	if err != nil {
		return common.Hash{}, err
	}

	return c.exchangeSender(input, gasPrice, nonce)
}

func (c *Client) BatchFillOrdersNoThrow(orders []models.Order, takerAssetFillAmounts []*big.Int, signatures [][]byte, gasPrice *big.Int, nonce uint64) (common.Hash, error) {
	input, err := exchange.BatchFillOrdersNoThrow(orders, takerAssetFillAmounts, signatures)
	if err != nil {
		return common.Hash{}, err
	}

	return c.exchangeSender(input, gasPrice, nonce)
}
