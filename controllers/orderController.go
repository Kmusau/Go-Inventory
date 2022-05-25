package controllers

import (
	"GoInventory/database"
	"GoInventory/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllOrders(c *fiber.Ctx) error {
	var orders []models.Order
	database.DB.Find(&orders)
	return c.JSON(&orders)
}

func GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order
	database.DB.Find(&order, id)
	return c.JSON(&order)
}

func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)

	if err := c.BodyParser(order); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Create(&order)
	return c.JSON(&order)
}
