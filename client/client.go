package client

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/outprog/gozx/models"
)

// Client for 0x protocol
type Client struct {
	key    *ecdsa.PrivateKey
	config *models.Config
}

func New(keyHex string, config *models.Config) *Client {
	k, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		panic(err)
	}
	return &Client{
		key:    k,
		config: config,
	}
}

func (c *Client) Address() common.Address {
	return crypto.PubkeyToAddress(c.key.PublicKey)
}
