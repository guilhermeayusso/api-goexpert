package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

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

	userDB := database.NewUser(db)
	UserHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Product routes
	r.Post("/products", ProductHandler.CreateProduct)
	r.Get("/products/{id}", ProductHandler.GetProduct)
	r.Get("/products", ProductHandler.GetAllProducts)
	r.Put("/products/{id}", ProductHandler.UpdateProduct)
	r.Delete("/products/{id}", ProductHandler.DeleteProduct)

	// User routes
	r.Post("/users", UserHandler.CreateUser)

	// Start the server
	http.ListenAndServe(":8000", r)

}
