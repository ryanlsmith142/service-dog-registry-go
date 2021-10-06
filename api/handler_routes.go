package main

import (
"errors"
"github.com/julienschmidt/httprouter"
"net/http"
"strconv"
)

func(a *api) getHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.logger.Print(errors.New("invalid id parameter"))
		a.errorJSON(w, err)
		return
	}

	handler, err := a.models.DB.GetHandler(id)

	err = a.writeJSON(w, http.StatusOK, handler, "handler")
	if err != nil {
		a.errorJSON(w, err)
		return
	}
}

func(a *api) getAllHandlersForOrganization(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.logger.Print(errors.New("invalid id parameter"))
		a.errorJSON(w, err)
		return
	}

	handlers, err := a.models.DB.GetAllHandlersForOrganization(id)

	err = a.writeJSON(w, http.StatusOK, handlers, "handlers")
	if err != nil {
		a.errorJSON(w, err)
		return
	}
}
