package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luqmansen/hanako/app"
	"github.com/luqmansen/hanako/controllers"
	"net/http"
	"os"
)

func main(){

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/anime/all", controllers.GetAll).Methods("GET")


	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Println(err)
	}

}