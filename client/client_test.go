package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/outprog/gozx/models"
)

var TestClient *Client

func init() {
	prv, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	prvHex := fmt.Sprintf("%x", crypto.FromECDSA(prv))
	TestClient = New(prvHex, &models.KovanConfig)
}
