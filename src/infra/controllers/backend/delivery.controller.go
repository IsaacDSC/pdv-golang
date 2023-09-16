package controllers_backend

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type DeliveryController struct {
	Path string
}

func NewDeliveryController() DeliveryController {
	return DeliveryController{
		Path: "/delivery",
	}
}

type DeliveryDTO struct {
	CheckoutId  string `json:"checkoutId"`
	DeliveryId  string `json:"deliveryId"`
	Address     string `json:"address"`
	HomeNumber  string `json:"homeNumber"`
	Observation string `json:"observation"`
	TypePayment string `json:"typePayment"`
}

func (*DeliveryController) Create(c *fiber.Ctx) error {
	var body DeliveryDTO
	buff := c.Body()
	err := json.Unmarshal(buff, &body)
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	fmt.Printf("\n%+v\n\n", body)
	return c.Status(fiber.StatusOK).Send([]byte("true"))
}
