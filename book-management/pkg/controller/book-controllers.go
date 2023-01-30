package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/artemxgod/11-go-projects/book-management/pkg/model"
	"github.com/artemxgod/11-go-projects/book-management/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook model.Book

func GetBooks() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		newbook := model.GetAllBooks()

		if err := json.NewEncoder(w).Encode(newbook); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func CreateBook() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		b := &model.Book{}
		if err := utils.ParseBody(r, b); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		b.CreateBook()

		if err := json.NewEncoder(w).Encode(b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetBookByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		p_id, err := strconv.ParseInt(params["id"], 0, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		searchBook, _ := model.GetBookByID(p_id)
		if err := json.NewEncoder(w).Encode(searchBook); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func UpdateBook() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		b := &model.Book{}
		if err := utils.ParseBody(r, b); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		params := mux.Vars(r)
		p_id, err := strconv.ParseInt(params["id"], 0, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		searchBook, db := model.GetBookByID(p_id)

		if b.Name != "" {
			searchBook.Name = b.Name
		}

		if b.Author != "" {
			searchBook.Author = b.Author
		}

		if b.Publication != "" {
			searchBook.Publication = b.Publication
		}

		db.Save(&searchBook)

		if err := json.NewEncoder(w).Encode(searchBook); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

}

func DeleteBook() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		p_id, err := strconv.ParseInt(params["id"], 0, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		searchBook := model.DeleteBook(p_id)
		if err := json.NewEncoder(w).Encode(searchBook); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetBooksByAuthor() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)

		authorName := params["name"]

		res := model.GetBooksByAuthor(authorName)

		if len(res) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
