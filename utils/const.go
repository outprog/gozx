package utils

import "github.com/ethereum/go-ethereum/crypto"

// assetData
const ERC20_ABI = `[{"name":"ERC20Token","inputs":[{"type":"address"}],"type":"function"},{"name":"res","inputs":[{"type":"address"}],"type":"event"}]`

// orderHash
const EIP191_HEADER = "\x19\x01"

var EIP712_DOMAIN_SEPARATOR_SCHEMA_HASH = crypto.Keccak256([]byte("EIP712Domain(string name,string version,address verifyingContract)"))

var EIP712_ORDER_SCHEMA_HASH = crypto.Keccak256([]byte("Order(address makerAddress,address takerAddress,address feeRecipientAddress,address senderAddress,uint256 makerAssetAmount,uint256 takerAssetAmount,uint256 makerFee,uint256 takerFee,uint256 expirationTimeSeconds,uint256 salt,bytes makerAssetData,bytes takerAssetData)"))

var EIP712_ZEROEX_TRANSACTION_SCHEMA_HASH = crypto.Keccak256([]byte("ZeroExTransaction(uint256 salt,address signerAddress,bytes data)"))

// signature
const (
	SIGNTYPE_Illegal                = iota // 0x00, default vae
	SIGNTYPE_Invalid                       // 0x01
	SIGNTYPE_EIP712                        // 0x02
	SIGNTYPE_EthSign                       // 0x03
	SIGNTYPE_Wallet                        // 0x04
	SIGNTYPE_Validator                     // 0x05
	SIGNTYPE_PreSigned                     // 0x06
	SIGNTYPE_NSignatureTypesIllegal        // 0x07, number of signature types. Always leave at end.
)
