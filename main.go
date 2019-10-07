package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luqmansen/hanako/controllers"
	"github.com/luqmansen/hanako/utils"
	"net/http"
	"os"
)


func main() {

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	//api.Use(app.JwtAuthentication)

	api.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
	api.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")

	api.HandleFunc("/anime/all", controllers.GetAll).Methods("GET")
	api.HandleFunc("/anime/{id}", controllers.GetById).Methods("GET")
	api.Path("/anime/search/q").Queries("title", "{title}").HandlerFunc(controllers.GetByTitle).Methods("GET")

	utils.Walk(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println(err)
	}

}
