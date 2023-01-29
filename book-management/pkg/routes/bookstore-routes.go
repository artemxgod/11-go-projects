package routes

import (
	ctrl "github.com/artemxgod/11-go-projects/book-management/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", ctrl.GetBooks()).Methods("GET")
	router.HandleFunc("/book/", ctrl.CreateBook()).Methods("POST")
	router.HandleFunc("/book/{id}", ctrl.GetBookByID()).Methods("GET")
	router.HandleFunc("/book/{id}", ctrl.UpdateBook()).Methods("PUT")
	router.HandleFunc("/book/{id}", ctrl.DeleteBook()).Methods("DELETE")

}