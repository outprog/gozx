package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/outprog/gozx/contracts/weth9"
	"github.com/outprog/gozx/models"
)

func (c *Client) WethDeposit(amount, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	input, err := weth9.Deposit()
	if err != nil {
		return
	}

	tx, err := types.SignTx(types.NewTransaction(
		nonce,
		c.config.WethAddress,
		amount,
		models.GasLimit,
		gasPrice,
		input),
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

func (c *Client) WethWithdraw(amount, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	input, err := weth9.Withdraw(amount)
	if err != nil {
		return
	}

	tx, err := types.SignTx(types.NewTransaction(
		nonce,
		c.config.WethAddress,
		big.NewInt(0),
		models.GasLimit,
		gasPrice,
		input),
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
