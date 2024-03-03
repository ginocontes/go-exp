package main

import (
	"fmt"
	"log"
	"net/http"

	"cumulus.org/becumulus/backend/controller"
	"cumulus.org/becumulus/backend/db"
	"cumulus.org/becumulus/backend/dto"
	"cumulus.org/becumulus/backend/service"
)

func main() {
	fmt.Println("Hello world!")
	db.Initialize()
	controller.RegisterHandlers()

	// p := dto.ProjectDTO{Name: "firstProject"}
	db := db.DB
	db.Exec("delete from projects where TRUE")

	p1 := dto.ProjectDTO{Name: "FirstProject"}
	p2 := dto.ProjectDTO{Name: "SecondProject"}
	service.CreateProject(&p1)
	service.CreateProject(&p2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
