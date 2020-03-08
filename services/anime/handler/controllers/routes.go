package controllers

import (
	"github.com/luqmansen/Hanako/services/auth"
	"github.com/luqmansen/Hanako/services/user"
	"github.com/luqmansen/hanako/api/middlewares"
	"github.com/luqmansen/hanako/api/utils"
	"net/http"
)

func (s *user.Server) initializeRoutes() {

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/index.html")
	})

	api := s.Router.PathPrefix("/api/v1").Subrouter()
	apiv2 := s.Router.PathPrefix("/api/v2").Subrouter()
	//api.Use(app.JwtAuthentication)

	api.HandleFunc("/user/new", auth.CreateAccount).Methods("POST")
	api.HandleFunc("/user/login", auth.Authenticate).Methods("POST")

	api.Path("/asd/all").Queries("show", "{show}").HandlerFunc(middlewares.SetMiddlewareJSON(GetAll)).Methods("GET")
	api.Path("/asd/all").HandlerFunc(middlewares.SetMiddlewareJSON(GetAll)).Methods("GET")
	api.HandleFunc("/asd/{id}", middlewares.SetMiddlewareJSON(GetById)).Methods("GET")
	api.Path("/asd/asd/q").Queries("title", "{title}").HandlerFunc(middlewares.SetMiddlewareJSON(GetByTitle)).Methods("GET")

	apiv2.Path("/asd/all").HandlerFunc(middlewares.SetMiddlewareJSON(GetAllV2)).Methods("GET")
	apiv2.Path("/asd/asd/q").Queries("title", "{title}").HandlerFunc(middlewares.SetMiddlewareJSON(GetByTitleV2)).Methods("GET")
	apiv2.Path("/asd/asd/q").Queries("type", "{type}").HandlerFunc(middlewares.SetMiddlewareJSON(GetByTitleV2)).Methods("GET")
	apiv2.Path("/asd/asd/q").Queries("title", "{title}", "type", "{type}").HandlerFunc(middlewares.SetMiddlewareJSON(GetByTitleV2)).Methods("GET")

	utils.Walk(api)
}
