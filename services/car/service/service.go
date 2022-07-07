package service

import (
	"database/sql"

	"stoik-leasing-cars/services/car"

	"github.com/google/uuid"
)

type Service struct {
	DB *sql.DB
}

func (s Service) Create(c car.Car) (car.Car, error) {
	c.ID = uuid.NewString()
	_, err := s.DB.Exec("INSERT INTO cars (id, name, price, year, color) VALUES ($1, $2, $3, $4, $5)", c.ID, c.Name, c.Price, c.Year, c.Color)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (s Service) All() ([]car.Car, error) {
	rows, err := s.DB.Query("SELECT * FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []car.Car
	for rows.Next() {
		var c car.Car
		err := rows.Scan(&c.ID, &c.Name, &c.Price, &c.Year, &c.Color)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}

func (s Service) ByID(id string) (car.Car, error) {
	var c car.Car
	err := s.DB.QueryRow("SELECT * FROM cars WHERE id = $1", id).Scan(&c.ID, &c.Name, &c.Price, &c.Year, &c.Color)
	if err != nil {
		return c, err
	}
	return c, nil
}
