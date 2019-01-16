package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/outprog/gozx/contracts/exchange"
	"github.com/outprog/gozx/models"
)

func (c *Client) MatchOrders(leftOrder, rightOrder models.Order, leftSignature, rightSignature []byte, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	input, err := exchange.MatchOrders(leftOrder, rightOrder, leftSignature, rightSignature)
	if err != nil {
		return common.Hash{}, err
	}

	return c.exchangeSender(input, gasPrice, nonce)
}
