package main

import (
	"../models"
	"encoding/json"
	"errors"
"github.com/julienschmidt/httprouter"
"net/http"
"strconv"
	"time"
)

type HandlerPayload struct {
	ID string `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	CertificationDate string `json:"CertificationDate"`
	CertificationExpirationDate string `json:"CertificationExpirationDate"`
	OrganizationID string `json:"OrganizationID"`
}

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

func (a *api) editHandler(w http.ResponseWriter, r *http.Request) {
	var payload HandlerPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		a.logger.Println(err)
		a.errorJSON(w, err)
		return
	}

	var handler models.Handler

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		h, _ := a.models.DB.GetHandler(id)
		handler = *h
	}

	handler.ID, _ = strconv.Atoi(payload.ID)
	handler.FirstName = payload.FirstName
	handler.LastName = payload.LastName
	handler.CertificationDate, _ = time.Parse("2006-01-02", payload.CertificationDate)
	handler.CertificationExpirationDate, _ = time.Parse("2006-01-02", payload.CertificationExpirationDate)
	handler.OrganizationID = payload.OrganizationID

	if handler.ID == 0 {
		//Insert Dog because dog doesn't exist yet.
	} else {
		err = a.models.DB.UpdateHandler(handler)
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

