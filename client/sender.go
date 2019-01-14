package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/outprog/gozx/models"
)

func (c *Client) exchangeSender(input []byte, gasPrice *big.Int, nonce uint64) (txHash common.Hash, err error) {
	tx, err := types.SignTx(types.NewTransaction(
		nonce,
		c.config.ExchangeContractAddress,
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

	txHash = tx.Hash()
	return
}
