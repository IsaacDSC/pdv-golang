package controllers

import (
	"pdv-example/settings"

	"github.com/gofiber/fiber/v2"
)

func CheckoutController(c *fiber.Ctx) error {
	template := settings.GetTemplateSingleton()
	return c.Render("checkout", fiber.Map{
		"title":  template.AppName,
		"navbar": template.Navbar,
		"footer": template.Footer,
	})
}
