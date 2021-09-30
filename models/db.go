package models

import "database/sql"

type DBModel struct {
	DB *sql.DB
}