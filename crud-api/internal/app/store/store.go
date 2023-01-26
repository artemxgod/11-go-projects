package store

import "github.com/artemxgod/11-go-projects/crudserver/internal/app/model"

type Store interface {
	Movie() MovieRepository
}

type MovieRepository interface {
	Create(*model.Movie)
	// Update(int)
	Delete(int) error
	Find(int) (*model.Movie, error)
	Get() (model.Movies)
}