# GO Project

## Store-Manager
Store Manager is a web application that helps store owners manage sales and product inventory records. This application is meant for use in a single store.

## Technologies Used

   - Golang For Backend


## Install Dependencies
- go get -u gorm.io/gorm

- go get -u gorm.io/driver/postgres

- go get -u github.com/gin-gonic/gin

- go get -u golang.org/x/crypto/bcrypt

- go get -u github.com/golang-jwt/jwt/v4

- go get github.com/joho/godotenv

- go get github.com/githubnemo/CompileDaemon

- go install github.com/githubnemo/CompileDaemon


## Environment Variable
- PORT = ***
- SECRET = ********
- DB = "****"


## expected Endpoints
  - POST: /api/v1/login
  - POST: /api/v1/product
  - POST: /api/v1/signup
  - POST: /api/v1/products

## Fire up Postman
  Enter the following routes
  - POST : localhost:8080/api/v1/signup 
        Enter password and username click send to signup

 - POST : localhost:8080/api/v1/login  
        Enter password and username click send to login 

 - POST : localhost:8080/api/v1/product
        To creat a single record

 - GET : localhost:8080/api/v1/products
        Get all products

 - GET : localhost:8080/api/v1/product/:id
        Get a product by id