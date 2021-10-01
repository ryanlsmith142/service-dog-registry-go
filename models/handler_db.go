package models

import (
	"context"
	"time"
)

func (m *DBModel) GetHandlers(id int) (*Handler, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from handler where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var handler Handler

	err := row.Scan (
		&handler.ID,
		&handler.FirstName,
		&handler.LastName,
		&handler.CertificationDate,
		&handler.CertificationExpirationDate,
		&handler.OrganizationID,
	)
	if err != nil {
		return nil, err
	}

	return &handler, nil
}