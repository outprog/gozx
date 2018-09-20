package exchange

const ExchangeABI = `[
	{
		"constant": false,
		"inputs": [
			{
				"name": "makerAddress",
				"type": "address"
			},
			{
				"name": "takerAddress",
				"type": "address"
			},
			{
				"name": "feeRecipientAddress",
				"type": "address"
			},
			{
				"name": "senderAddress",
				"type": "address"
			},
			{
				"name": "makerAssetAmount",
				"type": "uint256"
			},
			{
				"name": "takerAssetAmount",
				"type": "uint256"
			},
			{
				"name": "makerFee",
				"type": "uint256"
			},
			{
				"name": "takerFee",
				"type": "uint256"
			},
			{
				"name": "expirationTimeSeconds",
				"type": "uint256"
			},
			{
				"name": "salt",
				"type": "uint256"
			},
			{
				"name": "makerAssetData",
				"type": "bytes"
			},
			{
				"name": "takerAssetData",
				"type": "bytes"
			},
			{
				"name": "takerAssetFillAmount",
				"type": "uint256"
			},
			{
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "fillOrder",
		"outputs": [
			{
				"name": "makerAssetFilledAmount",
				"type": "uint256"
			},
			{
				"name": "takerAssetFilledAmount",
				"type": "uint256"
			},
			{
				"name": "makerFeePaid",
				"type": "uint256"
			},
			{
				"name": "takerFeePaid",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
