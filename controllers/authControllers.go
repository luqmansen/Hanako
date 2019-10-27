package controllers

import (
	"encoding/json"
	"github.com/luqmansen/hanako/models/models-postegres"
	u "github.com/luqmansen/hanako/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models_postegres.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid Request"))
		return
	}
	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models_postegres.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid Request"))
		return
	}
	resp := models_postegres.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
