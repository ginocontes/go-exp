package service

import (
	"fmt"

	"cumulus.org/becumulus/backend/db"
	"cumulus.org/becumulus/backend/dto"
	"cumulus.org/becumulus/backend/model"
)

// todo implement error handling
func CreateUser(u *dto.UserDTO) {
	db := db.GetDB()
	newUser := model.User{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}

	var existsUser model.User
	result := db.First(&existsUser, "username = ?", newUser.Username)
	if result.Error == nil {
		fmt.Println("Username already exists!")
		return
	}
	db.Create(&newUser)
}
