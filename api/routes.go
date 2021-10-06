package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *api) newRouter() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", a.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/user/:id", a.getUser)
	router.HandlerFunc(http.MethodGet, "/v1/organizationUsers/:id", a.getAllUsersForOrganization)

	router.HandlerFunc(http.MethodGet, "/v1/dog/:id", a.getDog)
	router.HandlerFunc(http.MethodGet, "/v1/organizationDogs/:id", a.getAllDogsForOrganization)

	router.HandlerFunc(http.MethodGet, "/v1/handler/:id", a.getHandler)
	router.HandlerFunc(http.MethodGet, "/v1/organizationHandlers/:id", a.getAllHandlersForOrganization)

	router.HandlerFunc(http.MethodGet, "/v1/organization/:id", a.getOrganization)
	router.HandlerFunc(http.MethodGet, "/v1/organizations", a.getAllOrganizations)



	return router
}

func (a *api) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := ApiStatus{
		Status: "Available",
		Environment: a.config.env,
		Version: version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
