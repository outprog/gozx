package client

import (
	"fmt"
	"testing"

	"github.com/outprog/gozx/models"
	"github.com/stretchr/testify/require"
)

func TestGenOrder(t *testing.T) {
	order, err := TestClient.GenOrder(models.KovanTokens["WETH"], models.KovanTokens["ZRX"], "1", "1")
	require.NoError(t, err)
	fmt.Printf("%+v\n", order)
}
