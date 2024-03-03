package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"cumulus.org/becumulus/backend/db"
	"cumulus.org/becumulus/backend/dto"
	"cumulus.org/becumulus/backend/model"
	"cumulus.org/becumulus/backend/service"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func InizializeMux() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		// handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v3/",
		Index,
	},

	Route{
		"AllUsers",
		"GET",
		"/api/v3/users",
		allUsers,
	},
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db := db.DB
	var users []model.User

	db.Find(&users)
	fmt.Println(users)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I'm here!")
	vars := mux.Vars(r)
	user := vars["user"]
	service.AllProjects(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	bodyReader := r.Body
	reqBody, err := io.ReadAll(bodyReader)
	if err != nil {
		log.Fatal("Error parsing request")
	}
	var newUser dto.UserDTO
	err = json.Unmarshal(reqBody, &newUser)
	if err != nil {
		log.Fatal(err)
	}
	service.CreateUser(&newUser)
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("Get all users")
	default:
		createUser(w, r)
	}
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
