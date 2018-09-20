package exchange

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/outprog/gozx/models"
	"github.com/outprog/gozx/utils"
)

func TestFillOrder(t *testing.T) {
	makerAsset, _ := utils.EncodeERC20AssetData(models.KovanTokens["WETH"].Address)
	takerAsset, _ := utils.EncodeERC20AssetData(models.KovanTokens["ZRX"].Address)
	salt, _ := new(big.Int).SetString("99392977337186991769370078621467871065903823087405146848913118809902183766712", 10)
	by, err := FillOrder(
		&models.Order{
			ExchangeAddress:       models.KovanConfig.ExchangeContractAddress,
			ExpirationTimeSeconds: big.NewInt(time.Now().Unix()),
			FeeRecipientAddress:   models.NullAddress,
			MakerAddress:          common.HexToAddress("0x7139c965955f706a1133b90a2f4bb8ed372dbfa9"),
			MakerAssetAmount:      big.NewInt(1),
			MakerAssetData:        makerAsset,
			MakerFee:              big.NewInt(1),
			// Salt:                  big.NewInt(1),
			Salt:             salt,
			SenderAddress:    models.NullAddress,
			TakerAddress:     models.NullAddress,
			TakerAssetData:   takerAsset,
			TakerAssetAmount: big.NewInt(1),
			TakerFee:         big.NewInt(1),
		},
		big.NewInt(1), common.FromHex("0x1bf904bee2638cdc35d351167860c9b2081d516314cfcb6de6f61098f0a2935b8a437e5afc7efa9eea287919dcc68a84ef868e8ba26af514bade179ffa3e9a279803"))
	fmt.Println(common.ToHex(by), err)
}
