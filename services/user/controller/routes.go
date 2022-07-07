package users

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	userService "stoik-leasing-cars/services/user/service"
)

type Resources struct {
	Users     userService.Service
	Validator *validator.Validate
}

func (rs Resources) Routes() chi.Router {
	r := chi.NewRouter()
	// User routes

	r.Post("/", rs.createUser)
	r.Get("/all", rs.getAllUsers)
	r.Get("/{id}", rs.getUser)

	return r
}
