package server

import (
	"fmt"
	"strconv"

	"github.com/IsaacDSC/pdv-golang/src/infra/server/http/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	// . "github.com/sqrthree/toFixed"
)

func Start_framework_server() {
	engine := handlebars.New("./src/views", ".hbs")
	engine.AddFunc("toInt", func(input string) string {
		value, err := strconv.Atoi(input)
		if err != nil {
			return "0"
		}
		price := float64(value) / float64(100)
		return fmt.Sprintf("%.2f", price)
	})
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./src/public")
	routes.StartRouter(app)
	fmt.Println("[ * ] starting server in port 3000")
	app.Listen(":3000")
}
