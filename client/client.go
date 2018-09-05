package client

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/outprog/gozx/models"
)

// Client for 0x protocol
type Client struct {
	key    ecdsa.PrivateKey
	config models.Config
}

func New(privateKey ecdsa.PrivateKey, config models.Config) *Client {
	return &Client{
		key:    privateKey,
		config: config,
	}
}

func (c *Client) Address() string {
	addr := crypto.PubkeyToAddress(c.key.PublicKey)
	return hexutil.Encode(addr[:])
}
