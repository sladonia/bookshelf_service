package books

import (
	"bookshelf_service/src/datasources/postgress/bookshelfdb"
	"fmt"
	"testing"
)

func TestAuthor_Save(t *testing.T) {
	bookshelfdb.InitDb(
		"localhost",
		"5432",
		"user",
		"password",
		"bookshelf_db",
	)

	author := Author{
		FirstName: "Jhon",
		LastName:  "Doe",
	}
	err := author.Save()
	fmt.Println(err)

	fmt.Println(author)
}
