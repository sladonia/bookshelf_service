package books

import (
	"bookshelf_service/src/domains"
	"database/sql"
	"fmt"
	"strings"
)

const (
	queryCreateAuthor = "INSERT INTO author(first_name, last_name) VALUES($1, $2) RETURNING id;"
	queryDeleteAuthor = "DELETE FROM author WHERE id = $1;"
	querySelectAuthor = "SELECT id, first_name, last_name FROM author WHERE id = $1;"
	queryUpdateAuthor = "UPDATE author SET first_name=$1, last_name=$2 WHERE id=$3;"
)

type Author struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AuthorInterface interface {
	Save() error
	Remove(id int64) error
}

func (a *Author) ValidateData() error {
	a.FirstName = strings.TrimSpace(a.FirstName)
	if a.FirstName == "" {
		return domains.DatabaseError{
			Message: "first name can not be empty",
			Err:     nil,
		}
	}
	a.LastName = strings.TrimSpace(a.LastName)
	if a.LastName == "" {
		return domains.DatabaseError{
			Message: "last name can not be empty",
			Err:     nil,
		}
	}
	return nil
}

func (a *Author) Save(db *sql.DB) error {
	err := a.ValidateData()
	if err != nil {
		return err
	}

	row := db.QueryRow(queryCreateAuthor, a.FirstName, a.LastName)

	var id int64
	err = row.Scan(&id)

	if err != nil {
		dbError := domains.DatabaseError{
			Message: "error trying to create Author",
			Err:     err,
		}
		return dbError
	}

	a.Id = id
	return nil
}

func (a *Author) Delete(db *sql.DB) error {
	result, err := db.Exec(queryDeleteAuthor, a.Id)
	number_deleted, err := result.RowsAffected()
	if err != nil {
		return err
	} else if number_deleted == 0 {
		return domains.DatabaseError{
			Message: fmt.Sprintf("no author with id=%d found", a.Id),
			Err:     err,
		}
	}
	return err
}

func (a *Author) Update(db *sql.DB) error {
	row := db.QueryRow(querySelectAuthor, a.Id)

	var id int64
	var firstName string
	var lastName string
	err := row.Scan(&id, &firstName, &lastName)
	if err != nil {
		return domains.DatabaseError{
			Message: fmt.Sprintf("no author with id=%d found", a.Id),
			Err:     err,
		}
	}

	if a.FirstName == "" {
		a.FirstName = firstName
	}
	if a.LastName == "" {
		a.LastName = lastName
	}

	result, err := db.Exec(queryUpdateAuthor, a.FirstName, a.LastName, a.Id)
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
