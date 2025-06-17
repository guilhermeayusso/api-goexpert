package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"

	"githug.com/guilhermeayusso/api-goexpert/infra/database"
	"githug.com/guilhermeayusso/api-goexpert/infra/webserver/handlers"
	"githug.com/guilhermeayusso/api-goexpert/internal/entity"
	"githug.com/guilhermeayusso/api-goexpert/pkg/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := config.LoadConfig(".")
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
	UserHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExperesIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		// Product routes
		r.Post("/", ProductHandler.CreateProduct)
		r.Get("/{id}", ProductHandler.GetProduct)
		r.Get("/", ProductHandler.GetAllProducts)
		r.Put("/{id}", ProductHandler.UpdateProduct)
		r.Delete("/{id}", ProductHandler.DeleteProduct)
	})

	// User routes
	r.Post("/users", UserHandler.CreateUser)
	r.Post("/users/generate_token", UserHandler.GetJWT)

	// Start the server
	http.ListenAndServe(":8000", r)

}
