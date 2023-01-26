package crudserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/artemxgod/11-go-projects/crudserver/internal/app/model"
	"github.com/artemxgod/11-go-projects/crudserver/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_GetMovie(t *testing.T) {
	store := teststore.New()
	mov := model.TestMovie(t)
	err := store.Movie().Create(mov)
	assert.NoError(t, err)

	testCasesParams := []struct {
		name         string
		route        string
		expectedCode int
	} {
		{
			name: "Acceptable",
			route: "/movies/1",
			expectedCode: http.StatusOK, 
		},
		{
			name: "Bad route",
			route: "/movies/r",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad id",
			route: "/movies/3",
			expectedCode: http.StatusNotFound, 
		},
	}
	s := newServer(store)

	for _, tc := range testCasesParams {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.route, nil)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_CreateMovie(t *testing.T) {
	store := teststore.New()

	testCasesParams := []struct {
		name         string
		mov 		*model.Movie
		expectedCode int
	} {
		{
			name: "Acceptable",
			mov: &model.Movie{
				Isbn: "1Ab",
				Title: "movie",
				Director: &model.Director{
					FirstName: "ab",
					LastName: "ba",
				},
			},
			expectedCode: http.StatusCreated, 
		},
		{
			name: "not valid isbn",
			mov: &model.Movie{
				ID: 5,
				Isbn: "1.2",
				Title: "movie",
				Director: &model.Director{
					FirstName: "ab",
					LastName: "ba",
				},
			},
			expectedCode: http.StatusUnprocessableEntity, 
		},
		{
			name: "not valid name",
			mov: &model.Movie{
				ID: 66,
				Isbn: "12",
				Title: "movie",
				Director: &model.Director{
					FirstName: "ab123",
					LastName: "ba",
				},
			},
			expectedCode: http.StatusUnprocessableEntity, 
		},
	}
	s := newServer(store)

	for _, tc := range testCasesParams {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.mov)
			req, _ := http.NewRequest(http.MethodPost, "/movies", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
