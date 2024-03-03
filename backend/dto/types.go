package dto

import "fmt"

type UserDTO struct {
	Username string
	Password string
	Email    string
}

func (u UserDTO) String() string {
	return fmt.Sprint(u.Username, u.Password, u.Email)
}

type ProjectDTO struct {
	Name string
}
