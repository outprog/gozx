package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetTransactionHash(salt *big.Int, takerAddress common.Address, data []byte, exchangeAddress common.Address) []byte {
	domainHash := crypto.Keccak256(EIP712_DOMAIN_SEPARATOR_SCHEMA_HASH,
		crypto.Keccak256([]byte("0x Protocol")),
		crypto.Keccak256([]byte("2")),
		addressToByte32(exchangeAddress),
	)
	return crypto.Keccak256([]byte(EIP191_HEADER),
		domainHash,
		hashTransaction(salt, takerAddress, data, exchangeAddress))
}

func hashTransaction(salt *big.Int, takerAddress common.Address, data []byte, exchangeAddress common.Address) []byte {
	return crypto.Keccak256(EIP712_ZEROEX_TRANSACTION_SCHEMA_HASH,
		math.PaddedBigBytes(salt, 32),
		addressToByte32(takerAddress),
		crypto.Keccak256(data),
	)
}
