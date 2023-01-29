package main

import (
	"log"
	"net/http"

	"github.com/artemxgod/11-go-projects/book-management/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	if err := http.ListenAndServe(":9010", r); err != nil {
		log.Fatal(err)
	}
}