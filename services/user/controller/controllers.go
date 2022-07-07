package users

import (
	"net/http"

	"stoik-leasing-cars/common"
	"stoik-leasing-cars/services/user"

	"github.com/go-chi/chi/v5"
)

// createUser
// @Summary create a new user
// @Produce application/json
// @Success 200 {array} model.User
// @Failure 400
// @Tags Users
// @Router /users [post]
func (rs Resources) createUser(w http.ResponseWriter, r *http.Request) {
	newUser, err := common.RequestBody[user.User](w, r, rs.Validator)
	if err != nil {
		return
	}

	userCreated, err := rs.Users.Create(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	common.SendJSON(w, userCreated, http.StatusCreated)
}

// getAllUsers
// @Summary get the list of all users
// @Produce application/json
// @Success 200 {array} model.User
// @Failure 400
// @Tags Users
// @Router /users/all [get]
func (rs Resources) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := rs.Users.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	common.SendJSON(w, users)
}

// getAllUsers
// @Summary get the list of all users
// @Produce application/json
// @Success 200 {array} model.User
// @Failure 400
// @Tags Users
// @Router /users/all [get]
func (rs Resources) getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	users, err := rs.Users.ByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	common.SendJSON(w, users)
}
