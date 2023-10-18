package controllers

import (
	   "github.com/gin-gonic/gin"
       "go-jwt/initializers"
        "go-jwt/model"
        "net/http"	
)



func Sales(c *gin.Context){
	var sales struct {
		Attendant string
		Totalprice int
	}
    
	if c.Bind(&sales) != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to read body",
		})

		return
	}

	// Create a sale

	sale := model.Sales{Attendant:sales.Attendant,Totalprice:sales.Totalprice}

    result := initializers.DB.Create(&sale)

	if result.Error != nil{
        c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to create sale",
		})

		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":sale,
	})
}
