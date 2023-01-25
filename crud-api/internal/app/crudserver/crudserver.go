package crudserver

import "net/http"

func Start(cfg *Config) error {
	s := newServer()

	return http.ListenAndServe(cfg.BindAddr, s)
}