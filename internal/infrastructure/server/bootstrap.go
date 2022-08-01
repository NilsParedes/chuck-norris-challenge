package server

import (
	"net/http"

	"github.com/nilsparedes/chuck-norris-challenge/internal/core/services"
	"github.com/nilsparedes/chuck-norris-challenge/internal/infrastructure/repositories/rest"
)

func Bootstrap(engine *http.ServeMux) {

	repository := rest.NewJokeRepository()
	service := services.NewJokeService(repository)
	getJokesHandler := GetJokesHandler(service)

	engine.HandleFunc("/api/jokes", getJokesHandler)
}
