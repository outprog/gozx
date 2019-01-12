package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/outprog/gozx/contracts/exchange"
	"github.com/outprog/gozx/models"
)

func (c *Client) FillOrder(order *models.Order, takerAssetFillAmount *big.Int, signature []byte, nonce uint64) (txHash common.Hash, err error) {
	input, err := exchange.FillOrder(order, takerAssetFillAmount, signature)
	if err != nil {
		return
	}

	tx, err := types.SignTx(types.NewTransaction(
		nonce,
		c.config.ExchangeContractAddress,
		big.NewInt(0),
		models.GasLimit,
		big.NewInt(10000),
		input),
		c.signer, c.key)
	if err != nil {
		return
	}

	err = c.ec.SendTransaction(context.TODO(), tx)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}
