package main

import (
	"../models"
	"encoding/json"
	"errors"
"github.com/julienschmidt/httprouter"
"net/http"
"strconv"

)

func(a *api) getUser(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.logger.Print(errors.New("invalid id parameter"))
		a.errorJSON(w, err)
		return
	}

	user, err := a.models.DB.GetUser(id)

	err = a.writeJSON(w, http.StatusOK, user, "user")
	if err != nil {
		a.errorJSON(w, err)
		return
	}
}

func(a *api) getAllUsersForOrganization(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.logger.Print(errors.New("invalid id parameter"))
		a.errorJSON(w, err)
		return
	}

	users, err := a.models.DB.GetAllDogsForOrganization(id)

	err = a.writeJSON(w, http.StatusOK, users, "users")
	if err != nil {
		a.errorJSON(w, err)
		return
	}
}

type UserPayload struct {
	ID string `json:"id""`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	OrganizationID string `json:"organizationID"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (a *api) editUser(w http.ResponseWriter, r *http.Request) {
	var payload UserPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		a.logger.Println(err)
		a.errorJSON(w, err)
		return
	}

	var user models.User

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		u, _ := a.models.DB.GetUser(id)
		user = *u
	}

	user.ID, _ = strconv.Atoi(payload.ID)
	user.FirstName = payload.FirstName
	user.LastName = payload.LastName
	user.Email = payload.Email
	user.Password = payload.Password
	user.OrganizationID = payload.OrganizationID

	if user.ID == 0 {
		//Insert Dog because dog doesn't exist yet.
	} else {
		err = a.models.DB.UpdateUser(user)
		if err != nil {
			a.errorJSON(w, err)
			return
		}
	}

	ok := JsonResp {
		OK: true,
	}

	err = a.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		a.errorJSON(w, err)
		return
	}

}
