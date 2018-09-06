package utils

import "github.com/ethereum/go-ethereum/crypto"

const ERC20_ABI = `[{"name":"ERC20Token","inputs":[{"type":"address"}],"type":"function"}]`

const EIP191_HEADER = "\x19\x01"

var EIP712_DOMAIN_SEPARATOR_SCHEMA_HASH = crypto.Keccak256([]byte("EIP712Domain(string name,string version,address verifyingContract)"))

var EIP712_ORDER_SCHEMA_HASH = crypto.Keccak256([]byte("Order(address makerAddress,address takerAddress,address feeRecipientAddress,address senderAddress,uint256 makerAssetAmount,uint256 takerAssetAmount,uint256 makerFee,uint256 takerFee,uint256 expirationTimeSeconds,uint256 salt,bytes makerAssetData,bytes takerAssetData)"))
