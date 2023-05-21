package main

import (
	"fmt"
	"pdv-example/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func start_framework_server() {
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	app.Get("/", controllers.HomeController)
	app.Get("/checkout", controllers.CheckoutController)
	app.Get("/cart", controllers.CartController)
	fmt.Println("[ * ] starting server in port 3000")
	app.Listen(":3000")
}

func main() {
	start_framework_server()
}
