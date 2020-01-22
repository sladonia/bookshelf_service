package app

import (
	"bookshelf_service/src/controllers"
	"github.com/gorilla/mux"
)

func mapUrls(r *mux.Router) {
	r.HandleFunc("/book/{id:[0-9]+}", controllers.BookstoreController.Get).Methods("GET")
	r.HandleFunc("/book", controllers.BookstoreController.Create).Methods("POST")
	r.HandleFunc("/book/search", controllers.BookstoreController.Search).Methods("GET")
}
