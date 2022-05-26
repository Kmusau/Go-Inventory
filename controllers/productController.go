package controllers

import (
	"GoInventory/database"
	"GoInventory/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

//Read all products in the catalog
func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Preload("Order").Find(&products)
	return c.JSON(&products)
}

//function to find first. find only a single item with id
func FindProduct(id int, product *models.Product) error {
	database.DB.Find(&product, "id = ?", id)

	if product.ID == 0 {

		return errors.New("product does not exist")

	}
	return nil
}

//read a single product
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err = FindProduct(id, &product); err != nil {
		c.Status(400).JSON(err.Error())
	}

	return c.JSON(product)
}

//Add a product to the catalog
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Create(&product)
	return c.JSON(&product)
}

//Update / edit product name
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

//Delete a product
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
