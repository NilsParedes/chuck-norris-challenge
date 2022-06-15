package main

import (
	"github.com/nilsparedes/chuck-norris-challenge/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/items", handlers.GetChuckNorrisHandler)
	_ = http.ListenAndServe(":8081", mux)
}
