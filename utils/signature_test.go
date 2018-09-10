package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestSignature(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	address := crypto.PubkeyToAddress(key.PublicKey)
	orderHash := common.FromHex("0xff077b1e304c3c30a2f3cf64285bff8434d97a6008203ba309f324c462242bca")

	_, err = Signature(key, orderHash, 7)
	assert.Equal(t, "Not support signatrue", err.Error())

	testSign, err := Signature(key, orderHash, SIGNTYPE_EthSign)
	require.NoError(t, err)

	isValid := IsValidSignture(orderHash, address, testSign)
	assert.Equal(t, true, isValid)
}
