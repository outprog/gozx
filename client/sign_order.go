package client

import (
	"fmt"

	"github.com/outprog/gozx/models"
	"github.com/outprog/gozx/utils"
)

func (c *Client) SignOrder(order *models.Order) (*models.Order, error) {
	orderHash, err := utils.GetOrderHash(order)
	if err != nil {
		return nil, err
	}
	fmt.Println(orderHash)
	return nil, nil
}
