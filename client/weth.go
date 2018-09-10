package client

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/outprog/gozx/contracts/weth9"
	"github.com/outprog/gozx/models"
)

func (c *Client) WethDeposit(wad string) (txHash common.Hash, err error) {
	nonce, err := c.ec.NonceAt(context.TODO(), c.Address(), nil)
	if err != nil {
		return
	}

	input, err := weth9.Deposit()
	if err != nil {
		return
	}

	amount, ok := new(big.Int).SetString(wad, 10)
	if !ok {
		err = errors.New("Can not convert wad to big.Int")
		return
	}

	tx, err := types.SignTx(types.NewTransaction(
		nonce,
		c.config.WethAddress,
		amount,
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

	return tx.Hash(), nil
}

func (c *Client) WethWithdraw(wad string) (txHash common.Hash, err error) {
	nonce, err := c.ec.NonceAt(context.TODO(), c.Address(), nil)
	if err != nil {
		return
	}

	amount, ok := new(big.Int).SetString(wad, 10)
	if !ok {
		err = errors.New("Can not convert wad to big.Int")
		return
	}

	input, err := weth9.Withdraw(amount)
	if err != nil {
		return
	}

	tx, err := types.SignTx(types.NewTransaction(
		nonce,
		c.config.WethAddress,
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

	return tx.Hash(), nil
}
