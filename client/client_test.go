package client

import (
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/outprog/gozx/models"
)

var TestClient *Client

func init() {
	prv, err := ecies.GenerateKey(rand.Reader, crypto.S256(), nil)
	if err != nil {
		panic(err)
	}
	TestClient = New(prv, &models.KovanConfig)
}
