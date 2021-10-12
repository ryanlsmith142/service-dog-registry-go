package models

import (
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

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
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	OrganizationID string `json:"organizationID"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Dog struct {
	ID int `json:"id""`
	Name string `json:"name"`
	WhelpDate time.Time `json:"whelpDate"`
	OrganizationID string `json:"organizationID"`
}

type Handler struct {
	ID int `json:"id""`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	CertificationDate time.Time `json:"certificationDate"`
	CertificationExpirationDate time.Time `json:"certificationExpirationDate"`
	OrganizationID string `json:"organizationID"`
}


