package app

import (
	"bookshelf_service/src/config"
	"bookshelf_service/src/datasources/postgress/bookshelfdb"
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

	bookshelfdb.InitDb(
		config.Config.BookshelfDb.Host,
		config.Config.BookshelfDb.Port,
		config.Config.BookshelfDb.User,
		config.Config.BookshelfDb.Password,
		config.Config.BookshelfDb.DbName,
	)

	r := mux.NewRouter()
	r.Use(logging.LoggingMw)
	mapUrls(r)

	srv := &http.Server{
		Addr:    config.Config.Port,
		Handler: r,
	}

	log.Infof("Start listening on port %s", config.Config.Port)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
