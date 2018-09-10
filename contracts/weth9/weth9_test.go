package weth9

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeposit(t *testing.T) {
	by, err := Deposit()
	require.NoError(t, err)
	assert.Equal(t, "0xd0e30db0", common.ToHex(by))
}

func TestWithdraw(t *testing.T) {
	by, err := Withdraw(big.NewInt(10000000000000000))
	require.NoError(t, err)
	assert.Equal(t, "0x2e1a7d4d000000000000000000000000000000000000000000000000002386f26fc10000", common.ToHex(by))
}
