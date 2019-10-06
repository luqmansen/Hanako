package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luqmansen/hanako/controllers"
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
