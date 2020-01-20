package app

import (
	"bookshelf_service/src/app/config"
	"bookshelf_service/src/controllers/books_controllers"
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

	http.HandleFunc("/book", books_controllers.GetBook)
	log.Debugf("Start listening on port %s", config.Config.Port)
	if err := http.ListenAndServe(config.Config.Port, nil); err != nil {
		panic(err)
	}
}
