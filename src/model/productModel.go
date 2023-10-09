package model

import ("gorm.io/gorm")

type Product struct{
	gorm.Model
	Productname  string
	Price int
	Quantity int 
	Description string
}

