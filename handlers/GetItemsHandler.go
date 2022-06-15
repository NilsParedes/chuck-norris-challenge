package handlers

import (
	"encoding/json"
	"github.com/nilsparedes/chuck-norris-challenge/services"
	"net/http"
)

func GetChuckNorrisHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	data := services.FetchChuckNorris()
	json.NewEncoder(w).Encode(data)
}
