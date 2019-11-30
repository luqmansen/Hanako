package controllers

import (
	"encoding/json"
	"github.com/luqmansen/hanako/api/models/postgres"
	u "github.com/luqmansen/hanako/api/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &postgres.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid Request"))
		return
	}
	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &postgres.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid Request"))
		return
	}
	resp := postgres.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
