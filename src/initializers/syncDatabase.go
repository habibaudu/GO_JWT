package initializers

import (
	"go-jwt/model"
)
type Product model.Product
func SyncDatabase() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.Sales{})
	DB.AutoMigrate(&model.ProductSales{})
	//DB.Migrator().DropTable(&model.ProductSales{})
	//DB.Migrator().DropTable(&model.Product{})
}

// Retrieve user list with eager loading credit cards
func GetAll() ([]Product, error) {
    var products []Product
    err := DB.Model(&model.Product{}).Preload("ProductSales").Find(&products).Error
    return products, err
}
