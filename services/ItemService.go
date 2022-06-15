package services

import (
	"encoding/json"
	"github.com/nilsparedes/chuck-norris-challenge/entities"
	"github.com/nilsparedes/chuck-norris-challenge/helpers"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var list []entities.Item

func FetchChuckNorris() []entities.Item {

	list = nil
	wg := sync.WaitGroup{}

	for len(list) < 25 {
		for i := len(list); i < 25; i++ {
			wg.Add(1)
			go addRandomChuckNorrisToList(&wg)
		}
		wg.Wait()

		list = helpers.FilterUniqueItems(list)
	}

	return list
}

func addRandomChuckNorrisToList(wg *sync.WaitGroup) {

	defer wg.Done()

	const api = "https://api.chucknorris.io/jokes/random"

	resp, err := http.Get(api)

	if err != nil {
		log.Println("Response error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Read error in content")
	}

	var item entities.Item
	_ = json.Unmarshal(body, &item)

	list = append(list, item)
}
