package client

import (
	"strconv"
	"time"

	"github.com/outprog/gozx/models"
	"github.com/outprog/gozx/utils"
)

func (c *Client) GenOrder(makerToken, takerToken models.Token, makerAssetAmount, takerAssetAmount string) (order *models.Order, err error) {
	makerAssetData, err := utils.EncodeERC20AssetData(makerToken.Address)
	if err != nil {
		return
	}
	takerAssetData, err := utils.EncodeERC20AssetData(takerToken.Address)
	if err != nil {
		return
	}
	salt := ""

	order = &models.Order{
		ExchangeAddress:       c.config.ExchangeContractAddress,
		ExpirationTimeSeconds: strconv.FormatInt((time.Now().Unix() + 60*60), 10),
		FeeRecipientAddress:   models.NullAddress,
		MakerAddress:          c.Address(),
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
	return
}

func (c *Client) SignOrder(order *models.Order, signType int) (signHex string, err error) {
	orderHash, err := utils.GetOrderHash(order)
	if err != nil {
		return "", err
	}

	return utils.Signature(c.key, orderHash, signType)
}
