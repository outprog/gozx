package client

import (
	"testing"
)

func TestFillOrder(t *testing.T) {
	// // new maker & taker Client
	// makerPrv := "set your prv key here"
	// takerPrv := "set your prv key here"
	// makerClient := New(makerPrv, &models.KovanConfig)
	// takerClient := New(takerPrv, &models.KovanConfig)

	// // approve first time
	// nonce, err := makerClient.Nonce()
	// require.NoError(t, err)
	// txHash, err := makerClient.Approve(models.KovanTokens["WETH"].Address, models.KovanConfig.Erc20ProxyContractAddress, big.NewInt(100000), nonce)
	// require.NoError(t, err)
	// fmt.Println(txHash.String())
	// nonce, err = takerClient.Nonce()
	// require.NoError(t, err)
	// txHash, err = takerClient.Approve(models.KovanTokens["CZT"].Address, models.KovanConfig.Erc20ProxyContractAddress, big.NewInt(100000), nonce)
	// require.NoError(t, err)
	// fmt.Println(txHash.String())

	// // maker new order
	// order, err := makerClient.NewOrder(models.KovanTokens["WETH"], models.KovanTokens["CZT"], new(big.Int).SetUint64(1234567*1e9), new(big.Int).SetUint64(1e18))
	// require.NoError(t, err)
	// fmt.Println(order)

	// // maker sign order
	// signature, err := makerClient.SignOrder(order, utils.SIGNTYPE_EthSign)
	// require.NoError(t, err)
	// fmt.Println(common.Bytes2Hex(signature))

	// // taker fillorder
	// nonce, err := takerClient.Nonce()
	// require.NoError(t, err)
	// txHash, err := takerClient.FillOrder(order, new(big.Int).SetUint64(1e18), signature, new(big.Int).SetUint64(1*1e9), nonce)
	// require.NoError(t, err)
	// fmt.Println(txHash.String())
}

func TestFillOrderNoThrow(t *testing.T) {
	// // new maker & taker Client
	// makerPrv := "set your prv key here"
	// takerPrv := "set your prv key here"
	// makerClient := New(makerPrv, &models.KovanConfig)
	// takerClient := New(takerPrv, &models.KovanConfig)

	// // maker new order
	// order, err := makerClient.NewOrder(models.KovanTokens["WETH"], models.KovanTokens["CZT"], new(big.Int).SetUint64(1234567*1e9), new(big.Int).SetUint64(1e18))
	// require.NoError(t, err)
	// fmt.Println(order)

	// // maker sign order
	// signature, err := makerClient.SignOrder(order, utils.SIGNTYPE_EthSign)
	// require.NoError(t, err)
	// fmt.Println(common.Bytes2Hex(signature))

	// // taker fillorder
	// nonce, err := takerClient.Nonce()
	// require.NoError(t, err)
	// txHash, err := takerClient.FillOrderNoThrow(order, new(big.Int).SetUint64(1e18), signature, new(big.Int).SetUint64(1*1e9), nonce)
	// require.NoError(t, err)
	// fmt.Println(txHash.String())
}

func TestBatchFillOrders(t *testing.T) {
	// // new maker & taker Client
	// makerPrv := "set your prv key here"
	// takerPrv := "set your prv key here"
	// makerClient := New(makerPrv, &models.KovanConfig)
	// takerClient := New(takerPrv, &models.KovanConfig)

	// // maker order1
	// order1, err := makerClient.NewOrder(models.KovanTokens["WETH"], models.KovanTokens["CZT"], new(big.Int).SetUint64(1234567*1e9), new(big.Int).SetUint64(1e18))
	// require.NoError(t, err)
	// fmt.Println(order1)
	// // maker sign order1
	// signature1, err := makerClient.SignOrder(order1, utils.SIGNTYPE_EthSign)
	// require.NoError(t, err)
	// fmt.Println(common.Bytes2Hex(signature1))
	// // maker order2
	// order2, err := makerClient.NewOrder(models.KovanTokens["WETH"], models.KovanTokens["CZT"], new(big.Int).SetUint64(1234567*1e9), new(big.Int).SetUint64(1e18))
	// require.NoError(t, err)
	// fmt.Println(order2)
	// // maker sign order2
	// signature2, err := makerClient.SignOrder(order2, utils.SIGNTYPE_EthSign)
	// require.NoError(t, err)
	// fmt.Println(common.Bytes2Hex(signature2))

	// // taker fillorder
	// nonce, err := takerClient.Nonce()
	// require.NoError(t, err)
	// txHash, err := takerClient.BatchFillOrders(
	// 	[]models.Order{order1, order2},
	// 	[]*big.Int{new(big.Int).SetUint64(1e17), new(big.Int).SetUint64(1e18)},
	// 	[][]byte{signature1, signature2},
	// 	new(big.Int).SetUint64(1*1e9), nonce)
	// require.NoError(t, err)
	// fmt.Println(txHash.String())
}

func TestBatchFillOrdersNoThrow(t *testing.T) {
	// // new maker & taker Client
	// makerPrv := "set your prv key here"
	// takerPrv := "set your prv key here"
	// makerClient := New(makerPrv, &models.KovanConfig)
	// takerClient := New(takerPrv, &models.KovanConfig)

	// // maker order1
	// order1, err := makerClient.NewOrder(models.KovanTokens["WETH"], models.KovanTokens["CZT"], new(big.Int).SetUint64(1234567*1e9), new(big.Int).SetUint64(1e18))
	// require.NoError(t, err)
	// fmt.Println(order1)
	// // maker sign order1
	// signature1, err := makerClient.SignOrder(order1, utils.SIGNTYPE_EthSign)
	// require.NoError(t, err)
	// fmt.Println(common.Bytes2Hex(signature1))
	// // maker order2
	// order2, err := makerClient.NewOrder(models.KovanTokens["WETH"], models.KovanTokens["CZT"], new(big.Int).SetUint64(1234567*1e9), new(big.Int).SetUint64(1e18))
	// require.NoError(t, err)
	// fmt.Println(order2)
	// // maker sign order2
	// signature2, err := makerClient.SignOrder(order2, utils.SIGNTYPE_EthSign)
	// require.NoError(t, err)
	// fmt.Println(common.Bytes2Hex(signature2))

	// // taker fillorder
	// nonce, err := takerClient.Nonce()
	// require.NoError(t, err)
	// txHash, err := takerClient.BatchFillOrdersNoThrow(
	// 	[]models.Order{order1, order2},
	// 	[]*big.Int{new(big.Int).SetUint64(1e17), new(big.Int).SetUint64(1e18)},
	// 	[][]byte{signature1, signature2},
	// 	new(big.Int).SetUint64(1*1e9), nonce)
	// require.NoError(t, err)
	// fmt.Println(txHash.String())
}
