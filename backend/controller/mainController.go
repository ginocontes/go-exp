package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cumulus.org/becumulus/backend/dto"
	"cumulus.org/becumulus/backend/service"
	"github.com/gorilla/mux"
)

func InizializeMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/createUser", createUserHandler)
	r.HandleFunc("/users/{user}", getUsersHandler)
	http.Handle("/", r)
	return r
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["user"])
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received")
	byt, _ := io.ReadAll(r.Body)
	var newUser dto.UserDTO
	json.Unmarshal(byt, &newUser) // unmarshal the content and then
	//save to database
	service.CreateUser(&newUser)
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
