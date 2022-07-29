package main

import (
	"log"
	"net/http"

	"github.com/nilsparedes/chuck-norris-challenge/internal/infrastructure/server"
)

func main() {
	mux := http.NewServeMux()
	server.Bootstrap(mux)
	log.Fatal(http.ListenAndServe(":8081", mux))
}
