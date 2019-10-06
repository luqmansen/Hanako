package controllers

import (
	"github.com/luqmansen/hanako/models"
	"github.com/luqmansen/hanako/utils"
	"net/http"
)

var GetAll = func(w http.ResponseWriter, r *http.Request){

	data := models.GetAll()
	resp := utils.Message(true, "Success")
	resp["data"] = data
	utils.Respond(w,resp)
}
