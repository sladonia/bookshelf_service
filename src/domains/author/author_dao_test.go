package author

import (
	"bookshelf_service/src/datasources/postgress/bookshelfdb"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func CleanData() {
	bookshelfdb.Client.Exec("DELETE FROM author WHERE TRUE;")
	bookshelfdb.Client.Exec("DELETE FROM book WHERE TRUE;")
	bookshelfdb.Client.Exec("DELETE FROM book_genre WHERE TRUE;")
	bookshelfdb.Client.Exec("DELETE FROM genre WHERE TRUE;")
}

func TestMain(m *testing.M) {
	bookshelfdb.InitDb(
		"localhost",
		"5433",
		"user",
		"password",
		"bookshelf_db",
		25,
		25,
		5,
	)
	if err := bookshelfdb.Client.Ping(); err != nil {
		fmt.Println("error connecting to test db", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestAuthor_Save(t *testing.T) {
	defer CleanData()
	firstName := "Jack"
	lastName := "Black"
	author := Author{
		FirstName: firstName,
		LastName:  lastName,
	}
	err := author.Save(bookshelfdb.Client)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, author.Id)

	var dbFirstName string
	var dbLastName string
	row := bookshelfdb.Client.QueryRow("SELECT first_name, last_name FROM author WHERE id=$1;", author.Id)
	err = row.Scan(&dbFirstName, &dbLastName)
	assert.Nil(t, err)
	assert.Equal(t, firstName, dbFirstName)
	assert.Equal(t, lastName, dbLastName)

	// trying to insert the same author second time
	err = author.Save(bookshelfdb.Client)
	assert.NotNil(t, err)
}

func TestAuthor_Remove(t *testing.T) {
	defer CleanData()
	unexistingAutor := Author{Id: 1}
	err := unexistingAutor.Delete(bookshelfdb.Client)
	assert.NotNil(t, err)

	firstName := "Jack"
	lastName := "Black"
	author := Author{
		FirstName: firstName,
		LastName:  lastName,
	}
	err = author.Save(bookshelfdb.Client)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, author.Id)
	err = author.Delete(bookshelfdb.Client)
	assert.Nil(t, err)
}

func TestAuthor_Update(t *testing.T) {
	defer CleanData()

	unexistingAuthor := Author{
		Id:        1,
		FirstName: "Ivan",
		LastName:  "Tyas",
	}
	err := unexistingAuthor.Update(bookshelfdb.Client)
	assert.NotNil(t, err)

	firstName := "Jack"
	lastName := "Black"
	author := Author{
		FirstName: firstName,
		LastName:  lastName,
	}
	err = author.Save(bookshelfdb.Client)
	assert.Nil(t, err)

	newFirstName := "Tony"
	newLastName := "Hawk"

	author.FirstName = newFirstName
	author.LastName = newLastName
	err = author.Update(bookshelfdb.Client)
	assert.Nil(t, err)

	var dbFirstName string
	var dbLastName string
	row := bookshelfdb.Client.QueryRow("SELECT first_name, last_name FROM author WHERE id=$1;", author.Id)
	err = row.Scan(&dbFirstName, &dbLastName)
	assert.Nil(t, err)
	assert.Equal(t, newFirstName, dbFirstName)
	assert.Equal(t, newLastName, dbLastName)
}

func TestAuthor_Retrieve(t *testing.T) {
	defer CleanData()
	firstName := "Jack"
	lastName := "Black"
	author := Author{
		FirstName: firstName,
		LastName:  lastName,
	}
	err := author.Save(bookshelfdb.Client)
	assert.Nil(t, err)

	retrievedAuthor := Author{Id: author.Id}
	err = retrievedAuthor.Retrieve(bookshelfdb.Client)
	assert.Nil(t, err)
	assert.Equal(t, firstName, retrievedAuthor.FirstName)
	assert.Equal(t, lastName, retrievedAuthor.LastName)
}
