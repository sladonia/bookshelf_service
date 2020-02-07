package books

import (
	"bookshelf_service/src/datasources/postgress/bookshelfdb"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	bookshelfdb.InitDb(
		"localhost",
		"5433",
		"user",
		"password",
		"bookshelf_db",
	)
	if err := bookshelfdb.Client.Ping(); err != nil {
		fmt.Println("error connecting to test db", err)
		os.Exit(1)
	}
	bookshelfdb.Client.Exec("DELETE FROM author WHERE TRUE;")
	bookshelfdb.Client.Exec("DELETE FROM book WHERE TRUE;")
	bookshelfdb.Client.Exec("DELETE FROM book_genre WHERE TRUE;")
	bookshelfdb.Client.Exec("DELETE FROM genre WHERE TRUE;")
	os.Exit(m.Run())
}

func TestAuthor_Save(t *testing.T) {
	firstName := "Jack"
	lastName := "Black"
	autor := Author{
		FirstName: firstName,
		LastName:  lastName,
	}
	err := autor.Save()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, autor.Id)

	var dbFirstName string
	var dbLastName string
	row := bookshelfdb.Client.QueryRow("SELECT first_name, last_name FROM author WHERE id=$1;", autor.Id)
	err = row.Scan(&dbFirstName, &dbLastName)
	assert.Nil(t, err)
	assert.Equal(t, firstName, dbFirstName)
	assert.Equal(t, lastName, dbLastName)

	// trying to insert the same author second time
	err = autor.Save()
	assert.NotNil(t, err)
}
