package model

import "testing"

func TestMovie(t *testing.T) *Movie {
	t.Helper()
	return &Movie{
		ID: 1,
		Isbn:  "1A234B48",
		Title: "The Game",
		Director: &Director{
			FirstName: "Verum",
			LastName:  "Virtus",
		},
	}
}