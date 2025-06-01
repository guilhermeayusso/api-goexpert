package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"githug.com/guilhermeayusso/api-goexpert/infra/database"
	"githug.com/guilhermeayusso/api-goexpert/internal/dto"
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

	http.ListenAndServe(":8000", nil)

}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)

	w.WriteHeader(http.StatusCreated)

}
