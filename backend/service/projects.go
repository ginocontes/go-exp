package service

import (
	"cumulus.org/becumulus/backend/db"
	"cumulus.org/becumulus/backend/dto"
	"cumulus.org/becumulus/backend/model"
)

func CreateProject(pDTO *dto.ProjectDTO) {
	db := db.DB

	var user model.User

	db.First(&user, "username = ?", "ginocontes")
	p := model.Project{
		Name:   pDTO.Name,
		UserID: user.Id,
	}
	db.Create(&p)
}
