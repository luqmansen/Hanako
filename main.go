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
	api.Path("/anime").Queries("search", "{title}").HandlerFunc(controllers.GetByTitle).Methods("GET")
	api.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
	http.Handle("/", api)


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
