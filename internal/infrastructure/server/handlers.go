package server

import (
	"encoding/json"
	"net/http"

	"github.com/nilsparedes/chuck-norris-challenge/internal/core/ports"
)

func GetChuckNorrisHandler(service ports.ChuckNorrisService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, _ := service.Fetch()
		json.NewEncoder(w).Encode(data)
	}
}
