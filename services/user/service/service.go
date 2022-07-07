package service

import (
	"database/sql"

	"stoik-leasing-cars/services/user"
)

type Service struct {
	DB *sql.DB
}

func (s Service) Create(u user.User) (user.User, error) {
	_, err := s.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s Service) All() ([]user.User, error) {
	rows, err := s.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (s Service) ByID(id string) (user.User, error) {
	var u user.User
	err := s.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return u, err
	}
	return u, nil
}
