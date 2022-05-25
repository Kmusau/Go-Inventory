package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Order    []Order
}