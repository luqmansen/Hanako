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
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/index.html")
	})

	api := r.PathPrefix("/api/v1").Subrouter()
	apiv2 := r.PathPrefix("/api/v2").Subrouter()
	//api.Use(app.JwtAuthentication)

	api.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
	api.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")

	api.Path("/anime/all").Queries("show", "{show}").HandlerFunc(controllers.GetAll).Methods("GET")
	api.Path("/anime/all").HandlerFunc(controllers.GetAll).Methods("GET")

	api.HandleFunc("/anime/{id}", controllers.GetById).Methods("GET")

	api.Path("/anime/search/q").Queries("title", "{title}").HandlerFunc(controllers.GetByTitle).Methods("GET")

	apiv2.Path("/anime/all").HandlerFunc(controllers.GetAllV2).Methods("GET")

	utils.Walk(api)
	//utils.Walk(apiv2)

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
