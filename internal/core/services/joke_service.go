package services

import (
	"sync"

	"github.com/nilsparedes/chuck-norris-challenge/internal/core/domain"
	"github.com/nilsparedes/chuck-norris-challenge/internal/core/helpers"
	"github.com/nilsparedes/chuck-norris-challenge/internal/core/ports"
)

type service struct {
	repository ports.JokeRepository
}

func NewJokeService(repository ports.JokeRepository) *service {
	return &service{repository: repository}
}

func (s *service) Fetch() ([]domain.Joke, error) {

	var (
		list []domain.Joke
		wg   sync.WaitGroup
	)

	for len(list) < 25 {
		for i := len(list); i < 25; i++ {
			wg.Add(1)
			go func() {
				item, _ := s.getRandomJoke()
				list = append(list, *item)
				defer wg.Done()
			}()
		}
		wg.Wait()

		list = helpers.FilterUniqueItems(list)
	}

	return list, nil
}

func (s *service) getRandomJoke() (*domain.Joke, error) {
	return s.repository.GetRandomJoke()
}
