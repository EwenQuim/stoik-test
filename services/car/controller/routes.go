package cars

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	carService "stoik-leasing-cars/services/car/service"
	userService "stoik-leasing-cars/services/user/service"
)

type Resources struct {
	Cars      carService.Service
	Users     userService.Service
	Validator *validator.Validate
}

func (rs Resources) Routes() chi.Router {
	r := chi.NewRouter()
	// Car routes

	r.Post("/", rs.createCar)
	r.Get("/all", rs.getAllCars)
	r.Get("/{id}", rs.getCar)

	r.Post("/{id}/rent", rs.rentCar)

	return r
}
