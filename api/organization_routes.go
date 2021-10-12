package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
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
	ID string `json:"id"`
	Name string `json:"name"`

	OrganizationID string `json:"OrganizationID"`
}
type Organization struct {
	ID int `json:"id""`
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


func (a *api) editDog(w http.ResponseWriter, r *http.Request) {
	var payload DogPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		a.logger.Println(err)
		a.errorJSON(w, err)
		return
	}

	var dog models.Dog

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		d, _ := a.models.DB.GetDog(id)
		dog = *d
	}

	dog.ID, _ = strconv.Atoi(payload.ID)
	dog.Name = payload.Name
	dog.WhelpDate, _ = time.Parse("2006-01-02", payload.WhelpDate)
	dog.OrganizationID = payload.OrganizationID

	if dog.ID == 0 {
		//Insert Dog because dog doesn't exist yet.
	} else {
		err = a.models.DB.UpdateDog(dog)
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
