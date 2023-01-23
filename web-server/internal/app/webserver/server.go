package webserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
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
	s.router.HandleFunc("/", s.handleMainPage()).Methods("GET")
	s.router.HandleFunc("/hello", s.handleHelloPage()).Methods("GET")
	s.router.HandleFunc("/form", s.handleFormPage()).Methods("GET", "POST")

	fileserver := http.FileServer(http.Dir("./ui/static"))
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileserver))

}

func (s *server) handleMainPage() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./ui/html/index.html")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err = tmpl.Execute(w, nil); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleFormPage() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./ui/html/form.html")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err = tmpl.Execute(w, nil); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err = r.ParseForm(); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return 
		}

		s.router.HandleFunc("/form/new", func(w http.ResponseWriter, r *http.Request) {
			name, addr := r.FormValue("name"), r.FormValue("address")
			fmt.Fprintf(w, "Your name is: %s\nYour address is: %s ", name, addr)
		})

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleHelloPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			s.error(w, r, http.StatusBadRequest, errors.New("400: bad request"))
			return
		}
		fmt.Fprint(w, "Hello!")
			s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
