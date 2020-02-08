package services

import (
	"bookshelf_service/src/domains"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	saveAuthorFunc   func(db *sql.DB) error
	deleteAuthorFunc func(db *sql.DB) error
	updateAuthorFunc func(db *sql.DB) error
)

type authorDaoMock struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (a *authorDaoMock) Save(db *sql.DB) error {
	return saveAuthorFunc(db)
}

func (a *authorDaoMock) Delete(db *sql.DB) error {
	return deleteAuthorFunc(db)
}

func (a *authorDaoMock) Update(db *sql.DB) error {
	return updateAuthorFunc(db)
}

func saveAuthorOk(db *sql.DB) error {
	return nil
}

func saveAuthorExists(db *sql.DB) error {
	return domains.DatabaseError{
		Message: "error trying to create Author",
		Err:     errors.New("user exists"),
	}
}

func TestAuthorService_Create(t *testing.T) {
	author := authorDaoMock{
		Id:        1,
		FirstName: "Jorge",
		LastName:  "Hamilgton",
	}

	// Test save author OK
	saveAuthorFunc = saveAuthorOk
	err := author.Save(nil)
	assert.Nil(t, err)

	// Test user exists
	saveAuthorFunc = saveAuthorExists
	err = author.Save(nil)
	assert.NotNil(t, err)
}
