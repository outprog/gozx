package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/outprog/gozx/contracts/erc20"
	"github.com/outprog/gozx/models"
)

func (c *Client) Approve(tokenAddr, spender common.Address, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	return c.approve(tokenAddr, spender, math.MaxBig256, gasPrice, nonce)
}

func (c *Client) UnApprove(tokenAddr, spender common.Address, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	return c.approve(tokenAddr, spender, big.NewInt(0), gasPrice, nonce)
}

func (c *Client) approve(tokenAddr, spender common.Address, value *big.Int, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	input, err := erc20.Approve(spender, value)
	if err != nil {
		return
	}

	tx, err := types.SignTx(types.NewTransaction(
		nonce,
		tokenAddr,
		big.NewInt(0),
		models.GasLimit,
		gasPrice,
		input),
		c.signer, c.key)
	if err != nil {
		return
	}

	return tx.Hash(), c.ec.SendTransaction(context.TODO(), tx)
}
