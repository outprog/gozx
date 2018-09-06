package client

import (
	"fmt"
	"testing"

	"github.com/outprog/gozx/models"
	"github.com/stretchr/testify/require"
)

func TestSignOrder(t *testing.T) {
	order, err := TestClient.SignOrder(models.KovanTokens["WETH"], models.KovanTokens["ZRX"])
	require.NoError(t, err)
	fmt.Printf("%+v\n", order)
}
