package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Movie struct {
	ID       int       `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Movies []*Movie

func (mov *Movie) Validate() error {
	if err := validation.ValidateStruct(mov.Director,
		validation.Field(
			&mov.Director.FirstName,
			validation.Required,
			is.Alpha,
		),
		validation.Field(
			&mov.Director.LastName,
			validation.Required,
			is.Alpha,
		),
	); err != nil {
		return err
	}
	if err := validation.ValidateStruct(mov,
		validation.Field(
			&mov.Isbn,
			validation.Required,
			is.Alphanumeric,
		),
	); err != nil {
		return err
	}

	return nil
}
