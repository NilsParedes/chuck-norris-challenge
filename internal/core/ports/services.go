package ports

import "github.com/nilsparedes/chuck-norris-challenge/internal/core/domain"

type JokeService interface {
	// Testeando
	// Testing
	Fetch() ([]domain.Joke, error)
}
