package main

import (
	"GoInventory/database"
	"GoInventory/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	routes.Routers(app)
	app.Get("/", hello)
	app.Listen(":3000")
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to Go API")
}
