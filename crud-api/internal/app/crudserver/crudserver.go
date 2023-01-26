package crudserver

import (
	"net/http"

	"github.com/artemxgod/11-go-projects/crudserver/internal/app/store/teststore"
)

func Start(cfg *Config) error {
	db := teststore.New()

	s := newServer(db)

	return http.ListenAndServe(cfg.BindAddr, s)
}