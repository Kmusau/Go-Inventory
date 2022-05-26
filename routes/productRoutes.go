package routes

import (
	"GoInventory/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	//product end points
	app.Get("/fetch/products", controllers.GetAllProducts)
	app.Get("/fetch/product/:id", controllers.GetProduct)
	app.Post("/create/product", controllers.CreateProduct)
	app.Delete("/delete/product/:id", controllers.Deleteproduct)
	app.Put("/update/product/:id", controllers.UpdateProduct)

	//orders endpoints
	app.Get("fetch/orders", controllers.GetAllOrders)
	app.Get("/fetch/order/:id", controllers.GetOrder)
	app.Post("create/order", controllers.CreateOrder)
	app.Delete("/delete/order/:id", controllers.DeleteOrder)
	app.Put("/update/order/:id", controllers.UpdateOrder)

	//stock endpoints
	app.Get("/fetch/stock", controllers.GetAllStocks)
	app.Get("/fetch/stock/:id", controllers.GetStock)
	app.Post("create/stock", controllers.AddStock)
	app.Delete("/delete/stock/:id", controllers.DeleteStock)
	app.Put("/update/stock/:id", controllers.UpdateStock)

}
