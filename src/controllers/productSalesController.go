package controllers

import (
	"go-jwt/initializers"
	"go-jwt/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ProductSales(c *gin.Context) {

	var productsale struct {
		Quantity int
		Total    int
		ProductID uint
	}

	if c.Bind(&productsale) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Create the Product Sales

	productsales := model.ProductSales{ProductID: productsale.ProductID,
		Quantity: productsale.Quantity, Total: productsale.Total}

	result := initializers.DB.Create(&productsales)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create product sale",
			"message":productsales,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": productsales,
	})
}
