package models

import (
	"context"
	"time"
)

func (m *DBModel) GetUser(id int) (*User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from user where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var user User

	err := row.Scan (
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.OrganizationID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
