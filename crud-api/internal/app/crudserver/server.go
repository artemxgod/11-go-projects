package crudserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/movies", s.handleGetMovies()).Methods("GET")
	s.router.HandleFunc("/movies", s.handleGetMovie()).Methods("GET")
	s.router.HandleFunc("/movies", s.handleCreateMovie()).Methods("POST")
	s.router.HandleFunc("/movies", s.handleUpdateMovie()).Methods("PUT")
	s.router.HandleFunc("/movies", s.handleDeleteMovie()).Methods("DELETE")

}

func (s *server) handleGetMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {

	}
}

func (s *server) handleGetMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {

	}
}

func (s *server) handleCreateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {

	}
}

func (s *server) handleUpdateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {

	}
}

func (s *server) handleDeleteMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {

	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}