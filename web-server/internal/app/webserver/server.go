package webserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/artemxgod/11-go-projects/web-server/internal/app/model"
	"github.com/artemxgod/11-go-projects/web-server/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

const (
	sessionName = "web-server"
	ctxKeyUser ctxKey = iota
)

type ctxKey int8

var (
	errNotAuthenticated = errors.New("not authenticated")
)

type server struct {
	router       *mux.Router
	store        store.Store
	sessionStore sessions.Store
}

func newServer(p_store store.Store, p_sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		store:        p_store,
		sessionStore: p_sessionStore,
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.handleMainPage()).Methods("GET")
	s.router.HandleFunc("/form", s.handleFormPage()).Methods("GET", "POST")

	fileserver := http.FileServer(http.Dir("./ui/static"))
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileserver))

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.Handle("/hello", s.handleHelloPage()).Methods("GET")

}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// getting cashed session by name
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		// checking if user session exists
		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errors.New("no id found"))
			return
		}

		// check if user is in database
		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		// if user was find, we are going on
		// we are sending respond with contect so next time we dont have to check this user
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
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
			tmpl, err = template.ParseFiles("./ui/html/userinfo.html")
			if err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}

			name, addr := r.FormValue("name"), r.FormValue("address")
			u := &model.User{
				Name:    name,
				Address: addr,
			}
			if err := s.store.User().Create(u); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
			}

			session, err := s.sessionStore.Get(r, sessionName)
			if err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}

			session.Values["user_id"] = u.ID

			if err := s.sessionStore.Save(r, w, session); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}

			// have to be AFTER SAVING SESSION
			if err = tmpl.Execute(w, u); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}

			s.respond(w, r, http.StatusCreated, u)

			// fmt.Fprintf(w, "Your name is: %s\nYour address is: %s ", name, addr)
		})
	}
}

func (s *server) handleHelloPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			s.error(w, r, http.StatusBadRequest, errors.New("400: bad request"))
			return
		}
		u := r.Context().Value(ctxKeyUser).(*model.User)

		fmt.Fprintf(w, "Hello, %s!", u.Name)
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
