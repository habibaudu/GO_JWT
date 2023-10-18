package initializers

import (
	"go-jwt/model"
)

func SyncDatabase(){
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.Sales{})
}