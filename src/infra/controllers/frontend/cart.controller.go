package controllers_frontend

import (
	"github.com/IsaacDSC/pdv-golang/src/infra/settings"
	"github.com/gofiber/fiber/v2"
)

func CartController(c *fiber.Ctx) error {
	template := settings.GetTemplateSingleton()
	return c.Render("cart", fiber.Map{
		"title":  template.AppName,
		"navbar": template.Navbar,
	})
}
