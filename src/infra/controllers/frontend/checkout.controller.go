package controllers_frontend

import (
	"encoding/json"
	"fmt"

	"github.com/IsaacDSC/pdv-golang/src/infra/settings"
	"github.com/gofiber/fiber/v2"
)

func CheckoutController(c *fiber.Ctx) error {
	template := settings.GetTemplateSingleton()
	return c.Render("checkout", fiber.Map{
		"title":    template.AppName,
		"navbar":   template.Navbar,
		"footer":   template.Footer,
		"delivery": template.Delivery,
	})
}

type Checkout struct {
	CheckoutId  string `json:"checkoutId"`
	DeliveryId  string `json:"deliveryId"`
	Address     string `json:"address"`
	HomeNumber  string `json:"homeNumber"`
	Observation string `json:"observation"`
	TypePayment string `json:"typePayment"`
}

func SendCheckoutController(c *fiber.Ctx) error {
	var body Checkout
	buff := c.Body()
	err := json.Unmarshal(buff, &body)
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	fmt.Printf("%+v", body)
	checkoutId := "123"
	send := fmt.Sprintf("{\"checkoutId\":\"%s\"}", checkoutId)
	return c.Status(fiber.StatusOK).Send([]byte(send))
}
