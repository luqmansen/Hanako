package controllers

import (
	models_mongo "github.com/luqmansen/hanako/models/models-mongo"
	u "github.com/luqmansen/hanako/utils"
	"net/http"
)

var dao = models_mongo.AnimeDAO{}

var GetAllV2 = func(w http.ResponseWriter, r *http.Request){

	data, err:= dao.FindAll()
	if err != nil {
		u.Message(http.StatusInternalServerError, err.Error())
	}
	if data == nil {
		u.Respond(w, u.Message(http.StatusNoContent, "Not Found"))
	} else {
		resp := u.Message(http.StatusOK, "Success")
		resp["data"] = data
		u.Respond(w, resp)
	}

}

var GetByTitleV2 = func(w http.ResponseWriter, r *http.Request){

	v := r.URL.Query()

	queryMap := map[string]interface{}{
		"title" : v.Get("title"),
		"type"	: v.Get("type"),
	}

	data, err:= dao.FindByQuery(queryMap)
	if err != nil {
		u.Message(http.StatusInternalServerError, err.Error())
	}
	if data == nil {
		u.Respond(w, u.Message(http.StatusNoContent, "Not Found"))
	} else {
		resp := u.Message(http.StatusOK, "Success")
		resp["data"] = data
		u.Respond(w, resp)
	}

}
