package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func(a *api) getOrganization(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.logger.Print(errors.New("invalid id parameter"))
		a.errorJSON(w, err)
		return
	}

	organization, err := a.models.DB.GetOrganization(id)

	err = a.writeJSON(w, http.StatusOK, organization, "organization")
	if err != nil {
		a.errorJSON(w, err)
		return
	}
}

func(a *api) getAllOrganizations(w http.ResponseWriter, r *http.Request) {

	organizations, err := a.models.DB.GetAllOrganizations()

	err = a.writeJSON(w, http.StatusOK, organizations, "organizations")
	if err != nil {
		a.errorJSON(w, err)
		return
	}
}
