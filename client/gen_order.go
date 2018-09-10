package client

import (
	"math/big"
	"time"

	"github.com/outprog/gozx/models"
	"github.com/outprog/gozx/utils"
)

func (c *Client) GenOrder(makerToken, takerToken models.Token, makerAssetAmount, takerAssetAmount *big.Int) (order *models.Order, err error) {
	makerAssetData, err := utils.EncodeERC20AssetData(makerToken.Address)
	if err != nil {
		return
	}
	takerAssetData, err := utils.EncodeERC20AssetData(takerToken.Address)
	if err != nil {
		return
	}
	salt := new(big.Int)

	order = &models.Order{
		ExchangeAddress:       c.config.ExchangeContractAddress,
		ExpirationTimeSeconds: big.NewInt((time.Now().Unix() + 60*60)),
		FeeRecipientAddress:   models.NullAddress,
		MakerAddress:          c.Address(),
		MakerAssetAmount:      makerAssetAmount,
		MakerAssetData:        makerAssetData,
		MakerFee:              big.NewInt(0),
		Salt:                  salt,
		SenderAddress:         models.NullAddress,
		TakerAddress:          models.NullAddress,
		TakerAssetAmount:      takerAssetAmount,
		TakerAssetData:        takerAssetData,
		TakerFee:              big.NewInt(0),
	}
	return
}

func (c *Client) SignOrder(order *models.Order, signType int) (sign []byte, err error) {
	return utils.Signature(c.key, utils.GetOrderHash(order), signType)
}
