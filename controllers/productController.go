package controllers

import (
	"GoInventory/database"
	"GoInventory/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Preload("Order").Find(&products)
	return c.JSON(&products)
}
func FindProduct(id int, product *models.Product) error {
	database.DB.Find(&product, "id = ?", id)

	if product.ID == 0 {

		return errors.New("product does not exist")

	}
	return nil
}
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err = FindProduct(id, &product); err != nil {
		c.Status(400).JSON(err.Error())
	}

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Create(&product)
	return c.JSON(&product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Product)

	database.DB.First(&product, id)

	if product.Name == "" {
		return c.Status(500).SendString("Product not available")
	}

	if err := c.BodyParser(product); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&product)
	return c.JSON(&product)
}

func Deleteproduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	database.DB.First(&product, id)

	if product.Name == "" {
		return c.Status(500).SendString("Product not available")
	}
	database.DB.Delete(&product)
	return c.SendString("Product deleted successfully")
}
