package controllers

import (
	"pdv-example/settings"

	"github.com/gofiber/fiber/v2"
)

func HomeController(c *fiber.Ctx) error {
	template := settings.GetTemplateSingleton()
	return c.Render("index", fiber.Map{
		"title":  template.AppName,
		"navbar": template.Navbar,
		"home":   template.Pages[0].Home,
		"footer": template.Footer,
	})
}
