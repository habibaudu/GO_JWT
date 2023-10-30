package controllers

import (
	   "github.com/gin-gonic/gin"
       "go-jwt/initializers"
        "go-jwt/model"
        "net/http"	
)


func ProductSales(c *gin.Context){
	var productsales struct {
		Products int
		Sales  int
		Quantity int
		Total int
	}
    
	if c.Bind(&productsales) != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to read body",
		})

		return
	}

	// Create the Product

	product := model.ProductSales{Products:productsales.Products,
		                     Sales:productsales.Sales,Quantity:products.Quantity,
							 Description:products.Description}

    result := initializers.DB.Create(&product)

	if result.Error != nil{
        c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to create product",
		})

		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":product,
	})
}
