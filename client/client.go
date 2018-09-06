package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/outprog/gozx/models"
)

// Client for 0x protocol
type Client struct {
	key    *ecies.PrivateKey
	config *models.Config
}

func New(privateKey *ecies.PrivateKey, config *models.Config) *Client {
	return &Client{
		key:    privateKey,
		config: config,
	}
}

func (c *Client) Address() common.Address {
	return crypto.PubkeyToAddress(c.key.ExportECDSA().PublicKey)
}
