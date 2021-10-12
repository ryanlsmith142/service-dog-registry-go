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


func (m *DBModel) GetAllOrganizations()([]*Organization, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from organization where organizationID = $1`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var organizations []*Organization

	for rows.Next() {
		var organization Organization
		err := rows.Scan(
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
		organizations = append(organizations, &organization)
	}
	return organizations, nil
}

func (m *DBModel) UpdateOrganization(organization Organization) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update organizations set name = $1, createdBy = $2, street = $3, state = $4, zipcode = $5, phoneNumber = $6, email = $7  where id = $8`

	_, err := m.DB.ExecContext(
		ctx,
		stmt,
		organization.CreatedBy,
		organization.Address.Street,
		organization.Address.State,
		organization.Address.Zipcode,
		organization.PhoneNumber,
		organization.Email,
	)
	if err != nil {
		return err
	}

	return nil
}
