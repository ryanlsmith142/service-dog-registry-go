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

func (m *DBModel) UpdateHandler(handler Handler) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update handlers set firstname = $1, lastname = $2, certificationDate = $3 certificationExpirationDate = $4 organizationID = $5 where id = $6`

	_, err := m.DB.ExecContext(
		ctx,
		stmt,
		handler.FirstName,
		handler.LastName,
		handler.CertificationDate,
		handler.CertificationExpirationDate,
		handler.OrganizationID,
		handler.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteHandler(handler Handler) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from handlers where id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, handler.ID)
	if err != nil {
		return err
	}

	return nil
}
