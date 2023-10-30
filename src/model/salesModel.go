package model

import ("gorm.io/gorm")

type Sales struct {
	gorm.Model
	Attendant string `gorm:"unique"`
	Totalprice int
	ProductSalesID uint
  }
