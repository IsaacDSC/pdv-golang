package controllers_backend

import (
	"fmt"

	"github.com/IsaacDSC/pdv-golang/shared/dto"
	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	Path string
}

func NewOrderController() OrderController {
	return OrderController{
		Path: "/order",
	}
}

func (oc *OrderController) Create(c *fiber.Ctx) error {
	order := new(dto.OrderInputDto)
	if err := c.BodyParser(order); err != nil {
		return err
	}
	fmt.Printf("\n%+v\n\n", order)

	return c.Status(fiber.StatusCreated).Send(c.Body())
}
