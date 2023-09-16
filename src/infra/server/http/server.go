package server

import (
	"fmt"

	"github.com/IsaacDSC/pdv-golang/src/infra/server/http/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func Start_framework_server() {
	engine := handlebars.New("./src/views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./src/public")
	routes.StartRouter(app)
	fmt.Println("[ * ] starting server in port 3000")
	app.Listen(":3000")
}
