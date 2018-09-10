package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/outprog/gozx/models"
)

func (c *Client) FillOrder(order *models.Order, signature string, takerAssetFillAmount string) (txHash common.Hash, err error) {
	nonce, err := c.ec.NonceAt(context.TODO(), c.Address(), nil)
	if err != nil {
		return
	}

	// TODO: replace data to fillOrder data
	tx, err := types.SignTx(types.NewTransaction(
		nonce,
		c.config.ExchangeContractAddress,
		big.NewInt(0),
		models.GasLimit,
		big.NewInt(10000),
		nil),
		c.signer, c.key)
	if err != nil {
		return
	}

	err = c.ec.SendTransaction(context.TODO(), tx)
	if err != nil {
		return
	}

	return tx.Hash(), nil
}
