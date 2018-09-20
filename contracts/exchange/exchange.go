package exchange

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/outprog/gozx/models"
)

var exchangeABI abi.ABI

func init() {
	var err error
	exchangeABI, err = abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		panic(err)
	}
}

func FillOrder(order *models.Order, takerAssetFillAmount *big.Int, signatrue []byte) ([]byte, error) {
	packNoTuple, err := exchangeABI.Pack("fillOrder",
		order.MakerAddress,
		order.TakerAddress,
		order.FeeRecipientAddress,
		order.SenderAddress,
		order.MakerAssetAmount,
		order.TakerAssetAmount,
		order.MakerFee,
		order.TakerFee,
		order.ExpirationTimeSeconds,
		order.Salt,
		order.MakerAssetData,
		order.TakerAssetData,
		takerAssetFillAmount, signatrue)
	if err != nil {
		return nil, err
	}

	methodHash := crypto.Keccak256([]byte("fillOrder({address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes},uint256,bytes)"))
	fmt.Println(common.ToHex(methodHash))

	return packNoTuple, nil
}
