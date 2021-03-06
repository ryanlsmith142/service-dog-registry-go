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

func (m *DBModel) GetAllUsersForOrganization(id int)([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select * from user where organizationID = $1`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.OrganizationID,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (m *DBModel) UpdateUser(user User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update users set firstName = $1, lastName = $2, organizationID = $3, email = $4 where id = $5`

	_, err := m.DB.ExecContext(ctx, stmt,user.FirstName, user.LastName, user.OrganizationID, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteUser(user User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from users where id = $1`

	_, err := m.DB.ExecContext(ctx, stmt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

