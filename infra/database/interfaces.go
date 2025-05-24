package database

import (
	"githug.com/guilhermeayusso/api-goexpert/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
