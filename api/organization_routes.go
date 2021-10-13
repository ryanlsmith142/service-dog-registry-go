package main

import (
	"../models"
	"encoding/json"
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
type OrganizationPayload struct {
	ID string `json:"id""`
	Name string `json:"name"`
	CreatedBy string `json:"createdBy"`
	Address struct {
		Street string `json:"street"`
		State string `json:"state"`
		Zipcode string `json:"zipcode"`
	}
	PhoneNumber string `json:"phoneNumber"`
	Email string `json:"email"`
}

func (a *api) editOrganization(w http.ResponseWriter, r *http.Request) {
	var payload OrganizationPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		a.logger.Println(err)
		a.errorJSON(w, err)
		return
	}

	var organization models.Organization

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		o, _ := a.models.DB.GetOrganization(id)
		organization = *o
	}

	organization.ID, _ = strconv.Atoi(payload.ID)
	organization.Name = payload.Name
	organization.CreatedBy = payload.CreatedBy
	organization.Address.Street = payload.Address.Street
	organization.Address.Zipcode = payload.Address.Zipcode
	organization.Address.State = payload.Address.State
	organization.PhoneNumber = payload.PhoneNumber
	organization.Email = payload.Email

	if organization.ID == 0 {
		//Insert Dog because dog doesn't exist yet.
	} else {
		err = a.models.DB.UpdateOrganization(organization)
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
