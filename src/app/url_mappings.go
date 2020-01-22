package app

import (
	"bookshelf_service/src/controllers/books_controllers"
	"github.com/gorilla/mux"
)

func mapUrls(r *mux.Router) {
	r.HandleFunc("/book", books_controllers.GetBook).Methods("GET")
	r.HandleFunc("/book", books_controllers.CreateBook).Methods("POST")
}
