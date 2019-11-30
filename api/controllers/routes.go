package controllers

import (
	"github.com/luqmansen/hanako/api/middlewares"
	"github.com/luqmansen/hanako/api/utils"
	"net/http"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/index.html")
	})

	api := s.Router.PathPrefix("/api/v1").Subrouter()
	apiv2 := s.Router.PathPrefix("/api/v2").Subrouter()
	//api.Use(app.JwtAuthentication)

	api.HandleFunc("/user/new", CreateAccount).Methods("POST")
	api.HandleFunc("/user/login", Authenticate).Methods("POST")

	api.Path("/anime/all").Queries("show", "{show}").HandlerFunc(middlewares.SetMiddlewareJSON(GetAll)).Methods("GET")
	api.Path("/anime/all").HandlerFunc(middlewares.SetMiddlewareJSON(GetAll)).Methods("GET")
	api.HandleFunc("/anime/{id}", middlewares.SetMiddlewareJSON(GetById)).Methods("GET")
	api.Path("/anime/search/q").Queries("title", "{title}").HandlerFunc(middlewares.SetMiddlewareJSON(GetByTitle)).Methods("GET")

	apiv2.Path("/anime/all").HandlerFunc(middlewares.SetMiddlewareJSON(GetAllV2)).Methods("GET")
	apiv2.Path("/anime/search/q").Queries("title", "{title}").HandlerFunc(middlewares.SetMiddlewareJSON(GetByTitleV2)).Methods("GET")
	apiv2.Path("/anime/search/q").Queries("type", "{type}").HandlerFunc(middlewares.SetMiddlewareJSON(GetByTitleV2)).Methods("GET")
	apiv2.Path("/anime/search/q").Queries("title", "{title}", "type", "{type}").HandlerFunc(middlewares.SetMiddlewareJSON(GetByTitleV2)).Methods("GET")

	utils.Walk(api)
}
