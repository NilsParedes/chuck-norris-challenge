package services

import (
	"sync"

	"github.com/nilsparedes/chuck-norris-challenge/internal/core/domain"
	"github.com/nilsparedes/chuck-norris-challenge/internal/core/helpers"
	"github.com/nilsparedes/chuck-norris-challenge/internal/core/ports"
)

type service struct {
	repository ports.ChuckNorrisRepository
}

func NewService(repository ports.ChuckNorrisRepository) *service {
	return &service{repository: repository}
}

var list []domain.Item

func (s *service) Fetch() ([]domain.Item, error) {

	list = nil
	wg := sync.WaitGroup{}

	for len(list) < 25 {
		for i := len(list); i < 25; i++ {
			wg.Add(1)
			go func() {
				item, _ := s.getRandomJoke(&wg)
				list = append(list, item)
			}()
		}
		wg.Wait()

		list = helpers.FilterUniqueItems(list)
	}

	return list, nil
}

func (s *service) getRandomJoke(wg *sync.WaitGroup) (domain.Item, error) {
	return s.repository.GetRandomJoke(wg)
}
