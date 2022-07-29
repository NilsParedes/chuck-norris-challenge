package ports

import (
	"sync"

	"github.com/nilsparedes/chuck-norris-challenge/internal/core/domain"
)

type ChuckNorrisRepository interface {
	GetRandomJoke(wg *sync.WaitGroup) (domain.Item, error)
}
