package service

import (
	"cumulus.org/becumulus/backend/db"
	"cumulus.org/becumulus/backend/dto"
	"cumulus.org/becumulus/backend/model"
)

func CreateUser(u dto.UserDTO) {
	db := db.GetDB()
	newUser := model.User{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
	db.Create(&newUser)
}
