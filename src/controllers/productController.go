package controllers

import (
	   "github.com/gin-gonic/gin"
       "go-jwt/initializers"
        "go-jwt/model"
        "net/http"		
)

type Products model.Product


func Product(c *gin.Context){
	var body struct {
		Productname string
		Price int
		Quantity int
		Description string
	}
    
	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to read body",
		})

		return
	}

	// Create the Product

	product := model.Product{Productname:body.Productname,
		                     Price:body.Price,Quantity:body.Quantity,
							 Description:body.Description}

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

func Get_Products(c *gin.Context){
    var products []Products
	

	results := initializers.DB.Find(&products)

	if results.Error != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Failed to get Products",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":products,
		"Count":results.RowsAffected,
	})

}

func Get_A_Product(c *gin.Context){
	var product []Products

	id := c.Param("id")

	results := initializers.DB.First(&product,id)

	if results.Error != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Product Not Found",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Message": product,
	})
}
