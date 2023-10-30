package model

import (
	"gorm.io/gorm"
)

type ProductSales struct{
	gorm.Model
	Products  []Product
	Sales  []Sales
	Total int
	Quantity int 
}
