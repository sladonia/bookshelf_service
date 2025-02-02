package app

import (
	"bookshelf_service/src/config"
	"bookshelf_service/src/datasources/postgress/bookshelfdb"
	"bookshelf_service/src/logger"
	"bookshelf_service/src/middlewares/logging"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartApp() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	if err := logger.InitLogger(config.Config.ServiceName, config.Config.LogLevel); err != nil {
		panic(err)
	}

	dbConfig := config.Config.BookshelfDb
	bookshelfdb.InitDb(
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DbName,
		dbConfig.MaxOpenConnections,
		dbConfig.MaxIdleConnections,
		dbConfig.ConnectionMaxLifetime,
	)

	r := mux.NewRouter()
	r.Use(logging.LoggingMw)
	mapUrls(r)

	srv := &http.Server{
		Addr:    config.Config.Port,
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Logger.Infof("Start listening on port %s", config.Config.Port)
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	<-done
	logger.Logger.Info("shutting down gracefully")
	logger.Logger.Sync()
}
