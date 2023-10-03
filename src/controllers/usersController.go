package controllers

import ("github.com/gin-gonic/gin"
	"os"
	"time"
       "go-jwt/initializers"
        "go-jwt/model"
        "net/http"
		"golang.org/x/crypto/bcrypt"
		"github.com/golang-jwt/jwt/v4"
		
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

	c.JSON(http.StatusOK,gin.H{

	})

}


func Login(c *gin.Context){
	var  body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to read Body",
		})

		return
	}
    

	var user model.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid Email or Password",
		})

		return
	}
	
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Failed to create Token",
		})
	}
    
	// send Token back as a cookie

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",tokenString,3600 * 24 * 30, "","", false,true)

	c.JSON(http.StatusOK,gin.H{
		
		// "token":tokenString,
	})

}
