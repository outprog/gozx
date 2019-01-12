package client

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/outprog/gozx/models"
	"github.com/outprog/gozx/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewOrder(t *testing.T) {
	_, err := TestClient.NewOrder(models.KovanTokens["WETH"], models.KovanTokens["ZRX"], big.NewInt(1), big.NewInt(1))
	require.NoError(t, err)
}

func TestSignOrder(t *testing.T) {
	order, err := TestClient.NewOrder(models.KovanTokens["WETH"], models.KovanTokens["ZRX"], big.NewInt(1), big.NewInt(1))
	require.NoError(t, err)

	sign, err := TestClient.SignOrder(order, utils.SIGNTYPE_EthSign)
	require.NoError(t, err)

	assert.Equal(t, 66, len(sign))
	assert.Equal(t, common.Hex2Bytes("0"+strconv.Itoa(utils.SIGNTYPE_EthSign)), sign[65:])
}
