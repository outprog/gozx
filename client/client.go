package client

import (
	"crypto/ecdsa"

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

func (c *Client) GenOrder() models.Order {
	return models.Order{}
}

func (c *Client) FillOrder() {}
