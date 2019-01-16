package client

import (
	"testing"
)

func TestMatchOrders(t *testing.T) {
	// // new maker1 & maker2 & sender Client
	// maker1Prv := ""
	// maker2Prv := ""
	// senderPrv := ""
	// maker1Client := New(maker1Prv, &models.KovanConfig)
	// maker2Client := New(maker2Prv, &models.KovanConfig)
	// senderClient := New(senderPrv, &models.KovanConfig)

	// // new two orders
	// order1, _ := maker1Client.NewOrder(
	// 	models.KovanTokens["WETH"],
	// 	models.KovanTokens["CZT"],
	// 	new(big.Int).SetUint64(1e16),
	// 	new(big.Int).SetUint64(1e18))
	// signature1, _ := maker1Client.SignOrder(order1, utils.SIGNTYPE_EthSign)
	// order2, _ := maker2Client.NewOrder(
	// 	models.KovanTokens["CZT"],
	// 	models.KovanTokens["WETH"],
	// 	new(big.Int).SetUint64(1e18),
	// 	new(big.Int).SetUint64(9*1e15))
	// signature2, _ := maker2Client.SignOrder(order2, utils.SIGNTYPE_EthSign)

	// // match orders
	// nonce, _ := senderClient.Nonce()
	// txHash, _ := senderClient.MatchOrders(order1, order2, signature1, signature2, new(big.Int).SetUint64(1*1e7), nonce)
	// fmt.Println(txHash.String())
}
