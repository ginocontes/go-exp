package model

type User struct {
	Id       uint
	Username string
	Password string
	Email    string
	Projects []Project //one user can have many projects
}

type Project struct {
	Name   string
	UserID uint
}
