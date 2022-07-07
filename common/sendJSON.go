package common

import (
	"encoding/json"
	"net/http"
)

// common.SendJSON sends a JSON response
func SendJSON(w http.ResponseWriter, e any, code ...int) {
	w.Header().Set("Content-Type", "application/json")
	if len(code) > 0 {
		w.WriteHeader(code[0])
	}

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
