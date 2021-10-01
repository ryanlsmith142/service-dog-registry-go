package models

import (
	"context"
	"time"
)

func (m *DBModel) GetHandler(id int) (*Handler, error){
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

func (m *DBModel) GetAllHandlersForOrganization(id int)([]*Handler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from handler where organizationID = $1`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var handlers []*Handler

	for rows.Next() {
		var handler Handler
		err := rows.Scan(
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
		handlers = append(handlers, &handler)
	}
	return handlers, nil
}