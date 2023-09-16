package routes

import (
	controllers_backend "github.com/IsaacDSC/pdv-golang/src/infra/controllers/backend"
	controllers_frontend "github.com/IsaacDSC/pdv-golang/src/infra/controllers/frontend"
	"github.com/gofiber/fiber/v2"
)

const (
	frontend_path = "/"
	backend_path  = "/api/v1"
)

func StartRouter(app *fiber.App) {
	routerFrontend(app)
	routerBackend(app)
}

func routerFrontend(app *fiber.App) {
	frontend := app.Group(frontend_path, func(c *fiber.Ctx) error {
		return c.Next()
	})
	frontend.Get("/", controllers_frontend.HomeController)
	frontend.Get("/checkout", controllers_frontend.CheckoutController)
	frontend.Get("/cart", controllers_frontend.CartController)
}

func routerBackend(app *fiber.App) {
	v1 := app.Group(backend_path, func(c *fiber.Ctx) error {
		return c.Next()
	})
	order := controllers_backend.NewOrderController()
	v1.Post(order.Path, order.Create)
	delivery := controllers_backend.NewDeliveryController()
	v1.Post(delivery.Path, delivery.Create)
	v1.Post("/client", func(c *fiber.Ctx) error { return nil })
}
