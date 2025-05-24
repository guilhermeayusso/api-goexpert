package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"githug.com/guilhermeayusso/api-goexpert/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user := entity.User{
		Name:     "John Doe",
		Email:    "john.doe@gmail.com",
		Password: "123456",
	}

	userDB := NewUser(db)
	err = userDB.Create(&user)
	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)

}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user := entity.User{
		Name:     "John Doe",
		Email:    "john.doe@gmail.com",
		Password: "123456",
	}

	userDB := NewUser(db)
	err = userDB.Create(&user)
	if err != nil {
		t.Error(err)
	}

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
