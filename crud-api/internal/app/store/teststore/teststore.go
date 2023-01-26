package teststore

import (
	"github.com/artemxgod/11-go-projects/crudserver/internal/app/model"
	"github.com/artemxgod/11-go-projects/crudserver/internal/app/store"
)

type Store struct {
	movieRepository *movieRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Movie() store.MovieRepository {
	if s.movieRepository != nil {
		return s.movieRepository
	}

	s.movieRepository = &movieRepository{
		movies: model.Movies{},
	}

	return s.movieRepository
}
