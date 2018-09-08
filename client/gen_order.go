package client

import (
	"strconv"
	"time"

	"github.com/outprog/gozx/models"
	"github.com/outprog/gozx/utils"
)

func (c *Client) GenOrder(makerToken, takerToken models.Token, makerAssetAmount, takerAssetAmount string) (*models.Order, error) {
	makerAssetData, err := utils.EncodeERC20AssetData(makerToken.Address)
	if err != nil {
		return nil, err
	}
	takerAssetData, err := utils.EncodeERC20AssetData(takerToken.Address)
	if err != nil {
		return nil, err
	}
	salt := ""

	return &models.Order{
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
	}, nil
}
