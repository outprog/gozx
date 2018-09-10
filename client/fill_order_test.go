package client

import (
	"fmt"
	"testing"

	"github.com/outprog/gozx/models"
)

func TestFillOrder(t *testing.T) {
	txHash, err := TestClient.FillOrder(&models.Order{}, "", "1")
	fmt.Println(txHash, err)
}
