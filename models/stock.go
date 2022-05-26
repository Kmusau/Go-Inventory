package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model

	Quantity    int     `json:"quantity"`
	BuyingPrice float32 `json:"buying_price"`
	ProductID   uint    `json:"product_id"`
}
