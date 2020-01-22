package app

import (
	"bookshelf_service/src/config"
	"bookshelf_service/src/middlewares/logging"
	"github.com/gorilla/mux"
	"github.com/sladonia/log"
	"net/http"
	"time"
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
	mapUrls(r)

	srv := &http.Server{
		Addr:         config.Config.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(config.Config.ReadTimeout) * time.Microsecond,
		WriteTimeout: time.Duration(config.Config.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.Config.IdleTimeout) * time.Second,
	}

	log.Infof("Start listening on port %s", config.Config.Port)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
