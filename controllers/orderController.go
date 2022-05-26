package controllers

import (
	"GoInventory/database"
	"GoInventory/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Read all orders
func GetAllOrders(c *fiber.Ctx) error {
	var orders []models.Order
	database.DB.Find(&orders)
	return c.JSON(&orders)
}

//function to find first. find only a single item with id
func FindOrder(id int, order *models.Order) error {
	database.DB.Find(&order, "id = ?", id)

	if order.ID == 0 {

		return errors.New("Order does not exist")

	}
	return nil
}

//Read a single order
func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err = FindOrder(id, &order); err != nil {
		c.Status(400).JSON(err.Error())
	}

	return c.JSON(&order)
}

//Create /(make) an order
func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)

	if err := c.BodyParser(order); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Create(&order)
	return c.JSON(&order)
}

//Update / order and order
func UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	order := new(models.Order)

	database.DB.First(&order, id)

	// if order.SellingPrice == 0 {
	// 	return c.Status(500).SendString("Order not available")
	// }

	if err := c.BodyParser(order); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&order)
	return c.JSON(&order)
}

//Delete an order
func DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order
	database.DB.First(&order, id)

	// if order.ProductID == 0 {
	// 	return c.Status(500).SendString("Order not available")
	// }

	database.DB.Delete(&order)
	return c.SendString("Order deleted successfully")
}
