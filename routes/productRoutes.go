package routes

import (
	"GoInventory/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Get("/fetch/products", controllers.GetAllProducts)

	app.Get("/fetch/product/:id", controllers.GetProduct)

	app.Post("/create/product", controllers.CreateProduct)

	app.Delete("/delete/product/:id", controllers.Deleteproduct)

	app.Put("/update/product/:id", controllers.UpdateProduct)

	app.Get("fetch/orders", controllers.GetAllOrders)

	app.Post("create/orders", controllers.CreateOrder)

}
