package server

import (
	"net/http"

	"github.com/nilsparedes/chuck-norris-challenge/internal/core/services"
	"github.com/nilsparedes/chuck-norris-challenge/internal/infrastructure/repositories/rest"
)

func Bootstrap(engine *http.ServeMux) {

	repository := rest.NewChuckNorrisRepository()
	service := services.NewService(repository)
	getChuckNorrisHandler := GetChuckNorrisHandler(service)

	engine.HandleFunc("/api/items", getChuckNorrisHandler)
}
