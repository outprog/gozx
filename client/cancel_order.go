package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/outprog/gozx/contracts/exchange"
	"github.com/outprog/gozx/models"
)

func (c *Client) CancelOrder(order models.Order, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	input, err := exchange.CancelOrder(order)
	if err != nil {
		return common.Hash{}, err
	}

	return c.exchangeSender(input, gasPrice, nonce)
}

func (c *Client) BatchCancelOrders(orders []models.Order, gasPrice *big.Int, nonce uint64) (common.Hash, error) {
	input, err := exchange.BatchCancelOrders(orders)
	if err != nil {
		return common.Hash{}, err
	}

	return c.exchangeSender(input, gasPrice, nonce)
}

func (c *Client) CancelOrdersUpTo(targetOrderEpoch *big.Int, gasPrice *big.Int, nonce uint64) (common.Hash, error) {
	input, err := exchange.CancelOrdersUpTo(targetOrderEpoch)
	if err != nil {
		return common.Hash{}, err
	}

	return c.exchangeSender(input, gasPrice, nonce)
}
