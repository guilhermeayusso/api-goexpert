package main

import (
	"fmt"
	"net/http"

	"githug.com/guilhermeayusso/api-goexpert/infra/database"
	"githug.com/guilhermeayusso/api-goexpert/infra/webserver/handlers"
	"githug.com/guilhermeayusso/api-goexpert/internal/entity"
	"githug.com/guilhermeayusso/api-goexpert/pkg/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	ProductHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", ProductHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)

}
