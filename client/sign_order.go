package client

import (
	"time"

	"github.com/outprog/gozx/models"
)

func (c *Client) SignOrder(makerToken, takerToken models.Token) models.Order {
	maker := c.Address()
	makerAssetData := ""
	takerAssetData := ""

	order := c.genOrder(maker, "", "", makerAssetData, takerAssetData)
	return order
}

func (c *Client) genOrder(maker, makerAssetAmount, takerAssetAmount, makerAssetData, takerAssetData string) models.Order {
	salt := ""

	return models.Order{
		ExchangeAddress:       c.config.ExchangeContractAddress,
		ExpirationTimeSeconds: time.Now().Unix() + 60*60,
		FeeRecipientAddress:   models.NullAddress,
		MakerAddress:          maker,
		MakerAssetAmount:      makerAssetAmount,
		MakerAssetData:        makerAssetData,
		MakerFee:              "0",
		Salt:                  salt,
		SenderAddress:         models.NullAddress,
		TakerAddress:          models.NullAddress,
		TakerAssetAmount:      takerAssetAmount,
		TakerAssetData:        takerAssetData,
		TakerFee:              "0",
	}
}
