package crudserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/artemxgod/11-go-projects/crudserver/internal/app/model"
	"github.com/artemxgod/11-go-projects/crudserver/internal/app/store"
	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(p_store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  p_store,
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {

	s.router.HandleFunc("/movies", s.handleGetMovies()).Methods("GET")
	s.router.HandleFunc("/movies/{id}", s.handleGetMovie()).Methods("GET")
	s.router.HandleFunc("/movies", s.handleCreateMovie()).Methods("POST")
	s.router.HandleFunc("/movies/{id}", s.handleUpdateMovie()).Methods("PUT")
	s.router.HandleFunc("/movies/{id}", s.handleDeleteMovie()).Methods("DELETE")

}

func (s *server) handleGetMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		s.respond(w, r, http.StatusOK, s.store.Movie().Get())
	}
}

func (s *server) handleGetMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		params := mux.Vars(r)
		p_id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		mov, err := s.store.Movie().Find(p_id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, mov)
	}
}

func (s *server) handleCreateMovie() http.HandlerFunc {
	type request struct {
		Isbn     string          `json:"isbn"`
		Title    string          `json:"title"`
		Director *model.Director `json:"director"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		mov := &model.Movie{
			ID:       len(s.store.Movie().Get()) + 1,
			Isbn:     req.Isbn,
			Title:    req.Title,
			Director: req.Director,
		}

		if err := s.store.Movie().Create(mov); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, mov)
	}
}

func (s *server) handleUpdateMovie() http.HandlerFunc {
	type request struct {
		Isbn     string          `json:"isbn"`
		Title    string          `json:"title"`
		Director *model.Director `json:"director"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// parsing {id} from route
		params := mux.Vars(r)
		p_id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Movie().Delete(p_id); err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		mov := &model.Movie{
			ID:       p_id,
			Isbn:     req.Isbn,
			Title:    req.Title,
			Director: req.Director,
		}

		if err := s.store.Movie().Create(mov); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, mov)
	}
}

func (s *server) handleDeleteMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		p_id, err := strconv.Atoi(params["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Movie().Delete(p_id); err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, s.store.Movie().Get())
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
