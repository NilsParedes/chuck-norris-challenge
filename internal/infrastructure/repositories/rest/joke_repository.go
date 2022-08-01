package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nilsparedes/chuck-norris-challenge/internal/core/domain"
)

const api = "https://api.chucknorris.io/jokes/random"

type repository struct {
}

func NewJokeRepository() *repository {
	return &repository{}
}

func (repository *repository) GetRandomJoke() (*domain.Joke, error) {

	resp, err := http.Get(api)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var joke domain.Joke
	_ = json.Unmarshal(body, &joke)

	return &joke, nil
}
