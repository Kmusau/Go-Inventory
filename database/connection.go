package database

import (
	"GoInventory/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DNS = "root:kelvin1234@tcp(127.0.0.1:3306)/goinventory?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func ConnectDB() {
	connection, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the Database")
	}

	DB = connection

	connection.AutoMigrate(&models.Product{}, &models.Order{})
}
