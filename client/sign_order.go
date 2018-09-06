package client

import (
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/outprog/gozx/models"
	"github.com/outprog/gozx/utils"
)

func (c *Client) SignOrder(makerToken, takerToken models.Token) (models.Order, error) {
	maker := c.Address()
	makerAssetData, err := utils.EncodeERC20AssetData(makerToken.Address)
	if err != nil {
		return models.Order{}, err
	}
	takerAssetData, err := utils.EncodeERC20AssetData(takerToken.Address)
	if err != nil {
		return models.Order{}, err
	}

	order := c.genOrder(maker, "1", "1", makerAssetData, takerAssetData)
	return order, nil
}

func (c *Client) genOrder(maker common.Address, makerAssetAmount, takerAssetAmount, makerAssetData, takerAssetData string) models.Order {
	salt := ""

	return models.Order{
		ExchangeAddress:       c.config.ExchangeContractAddress,
		ExpirationTimeSeconds: strconv.FormatInt((time.Now().Unix() + 60*60), 10),
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
