package bookshelfdb

import (
	"bookshelf_service/src/config"
	"bookshelf_service/src/logger"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

var (
	Client *sql.DB
)

func InitDb(host, port, user, password, dbName string, maxConn, maxIdleConn, connLifetime int) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	var err error
	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	Client.SetMaxOpenConns(maxConn)
	Client.SetMaxIdleConns(maxIdleConn)
	Client.SetConnMaxLifetime(time.Duration(connLifetime) * time.Minute)
	logger.Logger.Debugf("%s database is configured", config.Config.BookshelfDb.DbName)
}
