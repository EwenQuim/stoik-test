package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger"

	carsControllers "stoik-leasing-cars/services/car/controller"
	carService "stoik-leasing-cars/services/car/service"
	userControllers "stoik-leasing-cars/services/user/controller"
	userService "stoik-leasing-cars/services/user/service"
)

const compressionLevel = 6

func NewGlobalController(st Store) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger, middleware.RedirectSlashes, middleware.CleanPath, middleware.Recoverer)

	r.Use(middleware.Heartbeat("/ping"))

	r.Use(middleware.Compress(compressionLevel, "text/html", "text/css", "application/json"))

	validator := validator.New()

	r.Group(func(router chi.Router) {
		router.Mount("/users", userControllers.Resources{
			Users: userService.Service{DB: st.DB}, Validator: validator,
		}.Routes())
		router.Mount("/cars", carsControllers.Resources{
			Cars: carService.Service{DB: st.DB}, Users: userService.Service{DB: st.DB}, Validator: validator,
		}.Routes())
	})

	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		b, err := os.ReadFile("./docs/swagger.json")
		if err != nil {
			panic(err)
		}

		w.Write(b)
	})

	if os.Getenv("ENV") == "docker" || os.Getenv("ENV") == "dev" {
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
		))
	}

	return r
}
