package books

import (
	"bookshelf_service/src/datasources/postgress/bookshelfdb"
	"strings"
)

const (
	queryCreateAuthor = "INSERT INTO author(first_name, last_name) VALUES($1, $2) RETURNING id;"
)

type AuthorDaoInterface interface {
	Create(firstName, lastName string) error
}

func (a *Author) Save() error {
	row := bookshelfdb.Client.QueryRow(queryCreateAuthor, a.FirstName, a.LastName)

	var id int64
	err := row.Scan(&id)

	if err != nil {
		dbError := DatabaseError{
			Message: "error trying to create Author",
			Err:     err,
		}
		return dbError
	}

	a.Id = id
	return nil
}

func (a *Author) CleanAndValidate() error {
	a.LastName = strings.TrimSpace(a.LastName)
	a.FirstName = strings.TrimSpace(a.FirstName)
	if a.LastName == "" {
		return NewValidationError("last_name can not be empty")
	} else if a.FirstName == "" {
		return NewValidationError("first_name can not be empty")
	}
	return nil
}
