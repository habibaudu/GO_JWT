package main

import (
	"go-jwt/middleware"
	"go-jwt/controllers"
	"go-jwt/initializers"
	"github.com/gin-gonic/gin"
        )

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	router := gin.Default()

	router.POST("api/v1/signup", controllers.Signup)
	router.POST("api/v1/login",controllers.Login)
	router.POST("api/v1/product",middleware.RequireAuth,controllers.Product)
	router.GET("api/v1/validate",middleware.RequireAuth,controllers.Validate)
	router.GET("api/v1/products",middleware.RequireAuth,controllers.Get_Products)
	router.GET("api/v1/product/:id",middleware.RequireAuth,controllers.Get_A_Product)

	router.Run()
}