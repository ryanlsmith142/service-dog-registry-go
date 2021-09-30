package models

import "database/sql"

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models {
		DB: DBModel{DB: db},
	}
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

type User struct {
	ID int `json:"id""`
	Name string `json:"name"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Organization string `json:"organization"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Dog struct {
	ID int `json:"id""`
	Name string `json:"name"`
	WhelpDate string `json:"whelpDate"`
	Organization string `json:"organization"`
}

type Handler struct {
	ID int `json:"id""`
	Name string `json:"name"`
	CertificationDate string `json:"certifiedDate"`
	CertificationExpirationDate string `json:"certificationExpirationDate"`
	Notes string `json:"notes"`
	Organization string `json:"organization"`
}


