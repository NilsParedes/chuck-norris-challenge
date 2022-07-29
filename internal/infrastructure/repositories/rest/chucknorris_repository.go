package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/nilsparedes/chuck-norris-challenge/internal/core/domain"
)

const api = "https://api.chucknorris.io/jokes/random"

type repository struct {
}

func NewChuckNorrisRepository() *repository {
	return &repository{}
}

func (repository *repository) GetRandomJoke(wg *sync.WaitGroup) (domain.Item, error) {

	defer wg.Done()

	resp, err := http.Get(api)

	if err != nil {
		log.Println("Response error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Read error in content")
	}

	var item domain.Item
	_ = json.Unmarshal(body, &item)

	return item, nil
}
