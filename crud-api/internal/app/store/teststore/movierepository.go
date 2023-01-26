package teststore

import (
	"errors"

	"github.com/artemxgod/11-go-projects/crudserver/internal/app/model"
)

type movieRepository struct {
	movies model.Movies
}

func (r *movieRepository) Get() model.Movies {
	return r.movies
}

func (r *movieRepository) Create(mov *model.Movie) error {
	if err := mov.Validate(); err != nil {
		return err
	}

	r.movies = append(r.movies, mov)
	return nil
}

func (r *movieRepository) Delete(ID int) error {
	if err := checkID(len(r.movies), ID); err != nil {
		return err
	} else {
		for idx, movie := range r.movies {
			if movie.ID == ID {
				r.movies = append(r.movies[:idx], r.movies[idx+1:]...)
				break
			}
		}
		return nil
	}
}

func (r *movieRepository) Find(ID int) (*model.Movie, error) {
	if err := checkID(len(r.movies), ID); err != nil {
		return &model.Movie{}, err
	} else {
		for idx, movie := range r.movies {
			if movie.ID == ID {
				return r.movies[idx], nil
			}
		}
	}
	return &model.Movie{}, errors.New("record not found")
}

func checkID(mov_len int, ID int) error {
	if ID > mov_len || ID < 0 {
		return errors.New("BAD ID NUMBER")
	}
	return nil
}
