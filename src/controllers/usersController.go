package controllers

import ("github.com/gin-gonic/gin"
       "go-jwt/initializers"
        "go-jwt/model"
        "net/http"
		"golang.org/x/crypto/bcrypt"
)
func Signup(c *gin.Context){
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to read body",
		})

		return
	}

	hash,err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to Hash password",
		})
		return
	}

	// Create the user

	user := model.User{Email:body.Email,Password:string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to create user",
		})

		return
	}

	// Respond

	c.JSON(http.StatusOK,gin.H{

	})

}
