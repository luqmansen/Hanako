package controllers

import (
	"fmt"
	models_mongo "github.com/luqmansen/hanako/api/models/mongo"
	"github.com/luqmansen/hanako/api/responses"
	"net/http"
)

var dao = models_mongo.AnimeDAO{}

var GetAllV2 = func(w http.ResponseWriter, r *http.Request) {

	data, err := dao.FindAll()
	if err != nil {
		fmt.Println(err)
		responses.JSON(w, http.StatusInternalServerError, err.Error())
	}
	resp := make(map[string]interface{})
	resp["data"] = data
	responses.JSON(w, http.StatusOK, resp)


}

var GetByTitleV2 = func(w http.ResponseWriter, r *http.Request) {

	v := r.URL.Query()

	queryMap := map[string]interface{}{
		"title": v.Get("title"),
		"type":  v.Get("type"),
	}

	data, err := dao.FindByQuery(queryMap)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err.Error())
	}
	//TODO refactor this condition, controllers shouldn't check nil data
	if data == nil {
		responses.JSON(w, http.StatusNoContent, "Not Found")
	} else {
		resp := make(map[string]interface{})
		resp["data"] = data
		responses.JSON(w, http.StatusOK, resp)
	}

}
