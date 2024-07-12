package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorhugonf/go-expert/07-APIs/internal/entity"
)

func TestCreateUser(t *testing.T) {
	db := DBOpen(t)
	db.AutoMigrate(&entity.User{})
	userDB := NewUser(db)
	user, _ := entity.NewUser("Victor", "v@v.com", "123")

	err := userDB.Create(user)

	var userFound entity.User
	assert.Nil(t, err)
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db := DBOpen(t)
	db.AutoMigrate(&entity.User{})
	userDB := NewUser(db)
	user, _ := entity.NewUser("Victor", "v@v.com", "123")
	err := userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
