package models

import (
	"context"
	"time"
)

func (m *DBModel) GetDog(id int) (*Dog, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from dog where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var dog Dog

	err := row.Scan (
		&dog.ID,
		&dog.Name,
		&dog.WhelpDate,
		&dog.OrganizationID,
	)
	if err != nil {
		return nil, err
	}

	return &dog, nil

}

func (m *DBModel) GetAllDogsForOrganization(id int)([]*Dog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from dog where organizationID = $1`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dogs []*Dog

	for rows.Next() {
		var dog Dog
		err := rows.Scan(
			&dog.ID,
			&dog.Name,
			&dog.WhelpDate,
			&dog.OrganizationID,
		)
		if err != nil {
			return nil, err
		}
		dogs = append(dogs, &dog)
	}
	return dogs, nil

}

func (m *DBModel) UpdateDog(dog Dog) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update dogs set name = $1, whelpDate = $2, organizationID = $3 where id = $4`

	_, err := m.DB.ExecContext(ctx, stmt,dog.Name, dog.WhelpDate, dog.OrganizationID, dog.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteDog(dog Dog) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from dogs where id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, dog.ID)
	if err != nil {
		return err
	}

	return nil
}