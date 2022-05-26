package controllers

import (
	"GoInventory/database"
	"GoInventory/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

//function to find everything in the stock
func GetAllStocks(c *fiber.Ctx) error {
	var stock []models.Stock
	database.DB.Find(&stock)
	return c.JSON(&stock)
}

//function to find first. find only a single item with id
func FindStock(id int, stock *models.Stock) error {
	database.DB.Find(&stock, "id = ?", id)

	if stock.ID == 0 {

		return errors.New("Stock does not exist")

	}
	return nil
}

//Read a single Stock
func GetStock(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var stock models.Stock

	if err = FindStock(id, &stock); err != nil {
		c.Status(400).JSON(err.Error())
	}

	return c.JSON(&stock)
}

//Create /(make) an stock
func AddStock(c *fiber.Ctx) error {
	stock := new(models.Stock)

	if err := c.BodyParser(stock); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Create(&stock)
	return c.JSON(&stock)
}

//Update / order and order
func UpdateStock(c *fiber.Ctx) error {
	id := c.Params("id")
	stock := new(models.Stock)

	database.DB.First(&stock, id)

	// if stock.BuyingPrice == 0 {
	// 	return c.Status(500).SendString("Stock not available")
	// }

	if err := c.BodyParser(stock); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&stock)
	return c.JSON(&stock)
}

//Delete an stock
func DeleteStock(c *fiber.Ctx) error {
	id := c.Params("id")
	var stock models.Stock
	database.DB.First(&stock, id)

	// if stock.BuyingPrice == 0 {
	// 	return c.Status(500).SendString("Stock not available")
	// }

	database.DB.Delete(&stock)
	return c.SendString("Stock deleted successfully")
}
