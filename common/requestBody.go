package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

const maxBodySize = 1024 * 1024 * 10 // 10MB

// Body deserializes the request body into the given type.
// If the request body is empty, it returns an error and write it to the response.
//
func RequestBody[T any](w http.ResponseWriter, r *http.Request, v *validator.Validate) (T, error) {
	var t T

	// Deserialize the request body
	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return t, errors.New(msg)
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&t)
	if err != nil {
		errWrapped := fmt.Errorf("cannot decode request body to %T: %w", t, err)
		http.Error(w, errWrapped.Error(), http.StatusBadRequest)
		return t, errWrapped
	}

	// Validate input
	if v != nil {
		err = v.Struct(t)
		if err != nil {
			// this check is only needed when your code could produce an
			// invalid value for validation such as interface with nil value
			if _, exists := err.(*validator.InvalidValidationError); exists {
				log.Println("validation error:", err)
			}
			for _, err := range err.(validator.ValidationErrors) {
				log.Println("invalid payload:", err, err.Value())
			}

			http.Error(w, err.Error(), http.StatusBadRequest)
			return t, err
		}
	}

	return t, nil
}
