package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func(a *api) getDog(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.logger.Print(errors.New("invalid id parameter"))
		a.errorJSON(w, err)
		return
	}

	dog, err := a.models.DB.GetDog(id)

	err = a.writeJSON(w, http.StatusOK, dog, "dog")
	if err != nil {
		a.errorJSON(w, err)
		return
	}
}

func(a *api) getAllDogsForOrganization(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.logger.Print(errors.New("invalid id parameter"))
		a.errorJSON(w, err)
		return
	}

	dogs, err := a.models.DB.GetAllDogsForOrganization(id)

	err = a.writeJSON(w, http.StatusOK, dogs, "dogs")
	if err != nil {
		a.errorJSON(w, err)
		return
	}
}


