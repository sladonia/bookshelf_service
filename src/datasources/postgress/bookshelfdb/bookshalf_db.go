package bookshelfdb

import (
	"bookshelf_service/src/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sladonia/log"
)

var (
	BookshelfDb *sql.DB
)

func InitDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config.BookshelfDb.Host,
		config.Config.BookshelfDb.Port,
		config.Config.BookshelfDb.User,
		config.Config.BookshelfDb.Password,
		config.Config.BookshelfDb.DbName)

	fmt.Printf("%+v\n", config.Config)
	fmt.Printf("%s\n", psqlInfo)

	var err error
	BookshelfDb, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = BookshelfDb.Ping(); err != nil {
		panic(err)
	}

	log.Debugf("%s database is configured", config.Config.BookshelfDb.DbName)
}
