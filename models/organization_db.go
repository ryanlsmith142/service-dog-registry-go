package models

import (
	"context"
	"time"
)

func (m *DBModel) GetOrganization(id int) (*Organization, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from organization where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var organization Organization

	err := row.Scan (
		&organization.ID,
		&organization.Name,
		&organization.CreatedBy,
		&organization.Address.Street,
		&organization.Address.State,
		&organization.Address.Zipcode,
		&organization.PhoneNumber,
		&organization.Email,
	)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}
