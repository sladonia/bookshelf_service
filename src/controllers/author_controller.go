package controllers

import (
	"bookshelf_service/src/domains/books"
	"bookshelf_service/src/domains/responses"
	"bookshelf_service/src/services"
	"encoding/json"
	"github.com/sladonia/log"
	"io/ioutil"
	"net/http"
)

var (
	AuthorController AuthorControllerInterface = &authorController{}
)

type AuthorControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type authorController struct {
}

func (a *authorController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMsg := "invalid body"
		log.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}
	defer r.Body.Close()
	var author books.Author
	err = json.Unmarshal(requestBody, &author)
	if err != nil {
		errorMsg := "invalid json body"
		log.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}

	result, err := services.AuthorService.Create(author)
	if err != nil {
		errorMsg := "unable to crate author object"
		log.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		ErrorResponse(w, apiErr)
		return
	}
	response := responses.ResponseCreated{
		Message:   "author created",
		CreatedId: result.Id,
	}
	JsonResponse(w, http.StatusCreated, response)
}

func (a *authorController) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
