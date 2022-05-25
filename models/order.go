package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	Quantity     int     `json:"quantity"`
	SellingPrice float32 `json:"selling_price"`
	ProductID    uint    `json:"product_id"`
}
