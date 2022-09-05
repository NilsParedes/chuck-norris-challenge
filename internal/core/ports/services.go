package ports

import "github.com/nilsparedes/chuck-norris-challenge/internal/core/domain"

type JokeService interface {
	Fetch() ([]domain.Joke, error)
}
