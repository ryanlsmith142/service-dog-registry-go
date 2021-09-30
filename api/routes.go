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
