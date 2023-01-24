package webserver

import (
	"net/http"

	"github.com/artemxgod/11-go-projects/web-server/internal/app/store/mapstore"
	"github.com/gorilla/sessions"
)

func Start(cfg *Config) error {
	db := mapstore.New()

	newss := sessions.NewCookieStore([]byte("secret"))

	s := newServer(db, newss)

	return http.ListenAndServe(cfg.BindAddr, s)
}