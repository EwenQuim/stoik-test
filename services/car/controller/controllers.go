package cars

import (
	"net/http"

	"stoik-leasing-cars/common"
	"stoik-leasing-cars/services/car"

	"github.com/go-chi/chi/v5"
)

// createCar
// @Summary create a new car
// @Produce application/json
// @Success 200 {array} model.Car
// @Failure 400
// @Tags Cars
// @Router /cars [post]
func (rs Resources) createCar(w http.ResponseWriter, r *http.Request) {
	newCar, err := common.RequestBody[car.Car](w, r, rs.Validator)
	if err != nil {
		return
	}

	carCreated, err := rs.Cars.Create(newCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	common.SendJSON(w, carCreated, http.StatusCreated)
}

// getAllCars
// @Summary get the list of all cars
// @Produce application/json
// @Success 200 {array} model.Car
// @Failure 400
// @Tags Cars
// @Router /cars/all [get]
func (rs Resources) getAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := rs.Cars.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	common.SendJSON(w, cars)
}

// getAllCars
// @Summary get the list of all cars
// @Produce application/json
// @Success 200 {array} model.Car
// @Failure 400
// @Tags Cars
// @Router /cars/all [get]
func (rs Resources) getCar(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	cars, err := rs.Cars.ByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	common.SendJSON(w, cars)
}
