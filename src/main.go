package main

import (
	"fmt"

	"github.com/IsaacDSC/pdv-golang/src/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func start_framework_server() {
	engine := handlebars.New("./src/views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./src/public")
	app.Get("/", controllers.HomeController)
	app.Get("/checkout", controllers.CheckoutController)
	app.Post("/checkout", controllers.SendCheckoutController)
	app.Get("/cart", controllers.CartController)
	fmt.Println("[ * ] starting server in port 3000")
	app.Listen(":3000")
}

func main() {
	start_framework_server()
}
