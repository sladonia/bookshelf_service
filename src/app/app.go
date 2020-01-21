package app

import (
	"bookshelf_service/src/config"
	"bookshelf_service/src/controllers/books_controllers"
	"bookshelf_service/src/middlewares/logging"
	"github.com/gorilla/mux"
	"github.com/sladonia/log"
	"net/http"
)

func StartApp() {
	if err := config.Load("config.yml"); err != nil {
		panic(err)
	}
	if err := log.InitLogger(config.Config.ServiceName, config.Config.LogLevel); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Use(logging.LoggingMw)
	r.HandleFunc("/book", books_controllers.GetBook).Methods("GET")
	r.HandleFunc("/book", books_controllers.CreateBook).Methods("POST")
	log.Infof("Start listening on port %s", config.Config.Port)
	if err := http.ListenAndServe(config.Config.Port, r); err != nil {
		panic(err)
	}
}
