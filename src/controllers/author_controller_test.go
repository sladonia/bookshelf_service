package controllers

import (
	"bookshelf_service/src/domains"
	"bookshelf_service/src/domains/author"
	"bookshelf_service/src/domains/responses"
	"bookshelf_service/src/logger"
	"bookshelf_service/src/services"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	createAuthorFunction func(author author.Author) (*author.Author, error)
	deleteAuthorFunction func(authorId int64) (*author.Author, error)
	getAuthorFunction    func(aithorId int64) (*author.Author, error)
	updateAuthorFunction func(author author.Author) (*author.Author, error)
)

type authorServiceMock struct{}

func (a *authorServiceMock) Create(author author.Author) (*author.Author, error) {
	return createAuthorFunction(author)
}

func (a *authorServiceMock) Delete(authorId int64) (*author.Author, error) {
	return deleteAuthorFunction(authorId)
}

func (a *authorServiceMock) Get(authorId int64) (*author.Author, error) {
	return getAuthorFunction(authorId)
}

func (a *authorServiceMock) Update(author author.Author) (*author.Author, error) {
	return updateAuthorFunction(author)
}

func createAuthorSuccess(author author.Author) (*author.Author, error) {
	return &author, nil
}

func createAuthorErrorAuhorExists(author author.Author) (*author.Author, error) {
	return nil, domains.DatabaseError{
		Message: "error trying to create Author",
		Err:     errors.New("author exists"),
	}
}

func TestMain(m *testing.M) {
	logger.InitLogger("bookshelf_service", "fatal")
	services.AuthorService = &authorServiceMock{}
	os.Exit(m.Run())
}

func TestAuthorController_CreateSuccess(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/author", strings.NewReader(`{"first_name":"James","last_name":"Blunt"}`))

	// test success
	createAuthorFunction = createAuthorSuccess
	AuthorController.Create(response, request)
	var result responses.ResponseCreated
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.Equal(t, "author created", result.Message)
}

func TestAuthorController_CreateErrorAuthorExists(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/author", strings.NewReader(`{"first_name":"James","last_name":"Blunt"}`))

	// author exists
	createAuthorFunction = createAuthorErrorAuhorExists
	AuthorController.Create(response, request)
	apiError, err := NewApiErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiError)
}
