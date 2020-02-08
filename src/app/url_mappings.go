package app

import (
	"bookshelf_service/src/controllers"
	"github.com/gorilla/mux"
)

func mapUrls(r *mux.Router) {
	r.HandleFunc("/", controllers.RootController.Get)

	// author routs
	r.HandleFunc("/author", controllers.AuthorController.Create).Methods("POST")
	r.HandleFunc("/author/{id:[0-9]+}", controllers.AuthorController.Update).Methods("POST")
	r.HandleFunc("/author/{id:[0-9]+}", controllers.AuthorController.Delete).Methods("DELETE")
	r.HandleFunc("/author/{id:[0-9]+}", controllers.AuthorController.Get).Methods("GET")
}
