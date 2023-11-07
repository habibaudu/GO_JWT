package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Productname  string
	Price        int
	Quantity     int
	Description  string
	ProductSales []ProductSales `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Sales struct {
	gorm.Model
	Attendant  string `gorm:"unique"`
	Totalprice int
}

type ProductSales struct {
	gorm.Model
	Total     int
	Quantity  int
	ProductID uint
}
