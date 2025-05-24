package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@gmail.com", "123456789")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@gmail.com", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@gmail.com", "123456789")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("123456789"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.NotEqual(t, "123456789", user.Password)
}
