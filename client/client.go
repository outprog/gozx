package client

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/outprog/gozx/models"
)

// Client for 0x protocol
type Client struct {
	key    *ecdsa.PrivateKey
	config *models.Config

	ec     *ethclient.Client
	signer types.EIP155Signer
}

func New(keyHex string, config *models.Config) *Client {
	k, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		panic(err)
	}

	c, err := rpc.Dial(config.EthRPC)
	if err != nil {
		panic(err)
	}
	ec := ethclient.NewClient(c)

	networkID, err := ec.NetworkID(context.TODO())
	if err != nil {
		panic(err)
	}

	return &Client{
		key:    k,
		config: config,

		ec:     ec,
		signer: types.NewEIP155Signer(networkID),
	}
}

func (c *Client) Address() common.Address {
	return crypto.PubkeyToAddress(c.key.PublicKey)
}
