# gozx

Implement 0x client & 0x utils

- [0x-v2-specification](https://github.com/0xProject/0x-protocol-specification/blob/master/v2/v2-specification.md)
- [0x.js doc](https://0xproject.com/docs/0xjs)

## Test

```
make test
```

## Config

Set contracts info into `./models/config.go`

```
// Config for kovan
var (
	KovanConfig = Config{
		EthRPC: "https://kovan.infura.io",

		Erc20ProxyContractAddress:  common.HexToAddress("0x9ad1b8209cea603892c9dfaa676bc737088b499a"),
		Erc721ProxyContractAddress: common.HexToAddress("0xbb428f7108971ecf1144700c4d37792f8b74f6b0"),
		ExchangeContractAddress:    common.HexToAddress("0xa458ec0709468996ef2ef668f5e52f37ceb66627"),

		WethAddress: KovanTokens["WETH"].Address,
	}

	// token list
	KovanTokens = map[string]Token{
		"WETH": Token{
			Symbol:   "WETH",
			Address:  common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c"),
			Decimals: 18,
			Type:     "ERC20",
		},
		"ZRX": Token{
			Symbol:   "ZRX",
			Address:  common.HexToAddress("0x6ff6c0ff1d68b964901f986d4c9fa3ac68346570"),
			Decimals: 18,
			Type:     "ERC20",
		}
	}
)
```

## Example

#### Approve

Before fillOrder, you should approve token's `transferFrom` to `Erc20ProxyContractAddress`.

```
	// init client from your private key
	prv := "set your prv key here"
	client := New(prv, &models.KovanConfig)
	
	// get nonce
	nonce, _ := makerClient.Nonce()
	
	// approve
	txHash, _ := client.Approve(models.KovanTokens["WETH"].Address, models.KovanConfig.Erc20ProxyContractAddress, big.NewInt(100000), nonce)
	fmt.Println("txHash:", txHash.String())
```

#### FillOrder

```
	// new maker & taker Client
	makerPrv := "set your prv key here"
	takerPrv := "set your prv key here"
	makerClient := New(makerPrv, &models.KovanConfig)
	takerClient := New(takerPrv, &models.KovanConfig)
	
	// maker: new order
	order, err := makerClient.NewOrder(models.KovanTokens["WETH"], models.KovanTokens["CZT"], new(big.Int).SetUint64(1234567*1e9), new(big.Int).SetUint64(1e18))
	require.NoError(t, err)
	fmt.Println(order)

	// maker: sign order
	signature, err := makerClient.SignOrder(order, utils.SIGNTYPE_EthSign)
	require.NoError(t, err)
	fmt.Println(common.Bytes2Hex(signature))

	// taker: fillorder
	nonce, err := takerClient.Nonce()
	gasPrice := new(big.Int).SetUint64(13*1e9)
	require.NoError(t, err)
	txHash, err := takerClient.FillOrder(order, new(big.Int).SetUint64(1e18), signature, gasPrice, nonce)
	require.NoError(t, err)
	fmt.Println("txHash:", txHash.String())
```