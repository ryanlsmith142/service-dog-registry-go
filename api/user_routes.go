package main

import (
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
