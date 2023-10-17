package controllers

import (
	   "github.com/gin-gonic/gin"
       "go-jwt/initializers"
        "go-jwt/model"
        "net/http"	
)

type Products model.Product


func Product(c *gin.Context){
	var products struct {
		Productname string
		Price int
		Quantity int
		Description string
	}
    
	if c.Bind(&products) != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to read body",
		})

		return
	}

	// Create the Product

	product := model.Product{Productname:products.Productname,
		                     Price:products.Price,Quantity:products.Quantity,
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

func Delete_A_Product(c *gin.Context){
	var product []Products

	id := c.Param("id")
	results := initializers.DB.Delete(&product,id)

	if results.Error != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Product Not Deleted",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Message":"Product Deleted",
	})
}

func Update_A_Product(c *gin.Context){
	
    id := c.Param("id")

	var product struct {
		Productname string
		Price int
		Quantity int
		Description string
	}
    
	if c.Bind(&product) != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to read body",
		})

		return
	}


	products := model.Product{Productname:product.Productname,
		Price:product.Price,Quantity:product.Quantity,
		Description:product.Description}


	results := initializers.DB.Model(&Products{}).Where("id = ?", id).Updates(&products)
	
	if results.Error != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Product not successfully updated",
		})

		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Message":"Product Updated successfully",
	})

}