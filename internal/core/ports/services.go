package ports

import "github.com/nilsparedes/chuck-norris-challenge/internal/core/domain"

type ChuckNorrisService interface {
	Fetch() ([]domain.Item, error)
}
