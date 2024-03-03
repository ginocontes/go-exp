package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cumulus.org/becumulus/backend/dto"
	"cumulus.org/becumulus/backend/service"
)

func RegisterHandlers() {
	// http.HandleFunc("/getUsers", handler)
	http.HandleFunc("/createUser", createUserHandler)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received")
	byt, _ := io.ReadAll(r.Body)
	var newUser dto.UserDTO
	json.Unmarshal(byt, &newUser) // unmarshal the content and then
	//save to database
	service.CreateUser(newUser)
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	//write json user
// 	gino := []model.User{
// 		{"Gino", "hello", "world"},
// 		{"Karen", "hello", "world"},
// 	}
// 	m, err := json.Marshal(gino)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	log.Default().Println(gino)
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(m)
// }
