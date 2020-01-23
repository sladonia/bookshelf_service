package app

import (
	"bookshelf_service/src/controllers"
	"github.com/gorilla/mux"
)

func mapUrls(r *mux.Router) {
	r.HandleFunc("/", controllers.RootController.Get)

	// book routs
	r.HandleFunc("/book/{id:[0-9]+}", controllers.BooksController.Get).Methods("GET")
	r.HandleFunc("/book", controllers.BooksController.Create).Methods("POST")
	r.HandleFunc("/book/search", controllers.BooksController.Search).Methods("GET")

	// author routs
	r.HandleFunc("/author", controllers.AuthorController.Create).Methods("POST")
}
