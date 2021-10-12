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

type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

type DogPayload struct {
	ID string `json:"id"`
	Name string `json:"Name"`
	WhelpDate string `json:"WhelpDate"`
	OrganizationID string `json:"OrganizationID"`
}

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

	ok := jsonResp {
		OK: true,
	}

	err = a.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		a.errorJSON(w, err)
		return
	}

}


