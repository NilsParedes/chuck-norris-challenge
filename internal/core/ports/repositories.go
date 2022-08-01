package ports

import (
	"github.com/nilsparedes/chuck-norris-challenge/internal/core/domain"
)

type JokeRepository interface {
	GetRandomJoke() (*domain.Joke, error)
}
