package exchange

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/outprog/gozx/models"
	"github.com/outprog/gozx/utils"
	"github.com/stretchr/testify/assert"
)

var testOrder models.Order
var testSignature []byte

func init() {
	makerAsset, _ := utils.EncodeERC20AssetData(models.KovanTokens["WETH"].Address)
	takerAsset, _ := utils.EncodeERC20AssetData(models.KovanTokens["ZRX"].Address)
	salt, _ := new(big.Int).SetString("99392977337186991769370078621467871065903823087405146848913118809902183766712", 10)
	testOrder = models.Order{
		ExchangeAddress:       models.KovanConfig.ExchangeContractAddress,
		ExpirationTimeSeconds: big.NewInt(time.Now().Unix()),
		FeeRecipientAddress:   models.NullAddress,
		MakerAddress:          common.HexToAddress("0x7139c965955f706a1133b90a2f4bb8ed372dbfa9"),
		MakerAssetAmount:      big.NewInt(1),
		MakerAssetData:        makerAsset,
		MakerFee:              big.NewInt(1),
		Salt:                  salt,
		SenderAddress:         models.NullAddress,
		TakerAddress:          models.NullAddress,
		TakerAssetData:        takerAsset,
		TakerAssetAmount:      big.NewInt(1),
		TakerFee:              big.NewInt(1),
	}
	testSignature = common.FromHex("0x1bf904bee2638cdc35d351167860c9b2081d516314cfcb6de6f61098f0a2935b8a437e5afc7efa9eea287919dcc68a84ef868e8ba26af514bade179ffa3e9a279803")
}

func TestFillOrder(t *testing.T) {
	_, err := FillOrder(testOrder, big.NewInt(1), testSignature)
	assert.NoError(t, err)
}

func TestFillOrderNoThrow(t *testing.T) {
	_, err := FillOrderNoThrow(testOrder, big.NewInt(1), testSignature)
	assert.NoError(t, err)
}

func TestBatchFillOrders(t *testing.T) {
	_, err := BatchFillOrders([]models.Order{testOrder, testOrder}, []*big.Int{big.NewInt(1), big.NewInt(1)}, [][]byte{testSignature, testSignature})
	assert.NoError(t, err)
}

func TestBatchFillOrdersNoThrow(t *testing.T) {
	_, err := BatchFillOrdersNoThrow([]models.Order{testOrder, testOrder}, []*big.Int{big.NewInt(1), big.NewInt(1)}, [][]byte{testSignature, testSignature})
	assert.NoError(t, err)
}

func TestCancelOrder(t *testing.T) {
	_, err := CancelOrder(testOrder)
	assert.NoError(t, err)
}

func TestCancelOrdersUpTo(t *testing.T) {
	_, err := CancelOrdersUpTo(testOrder.Salt)
	assert.NoError(t, err)
}

func TestBatchCancelOrders(t *testing.T) {
	_, err := BatchCancelOrders([]models.Order{testOrder, testOrder})
	assert.NoError(t, err)
}

func TestMatchOrders(t *testing.T) {
	_, err := MatchOrders(testOrder, testOrder, testSignature, testSignature)
	assert.NoError(t, err)
}
