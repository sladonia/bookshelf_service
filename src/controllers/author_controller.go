package controllers

import (
	"bookshelf_service/src/domains/requests"
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
	authorRequest := requests.AuthorRequest{}
	if err = json.Unmarshal(requestBody, &authorRequest); err != nil {
		errorMsg := "invalid json body"
		log.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}
	if err := authorRequest.CleanAndValidate(); err != nil {
		errorMsg := "invalid json body"
		log.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusBadRequest)
		ErrorResponse(w, apiErr)
		return
	}

	author, err := services.AuthorService.Create(authorRequest.FirstName, authorRequest.LastName)
	if err != nil {
		errorMsg := "unable to crate author object"
		log.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		ErrorResponse(w, apiErr)
		return
	}
	response := responses.ResponseCreated{
		Message:   "author created",
		CreatedId: author.Id,
	}
	JsonResponse(w, http.StatusCreated, response)
}

func (a *authorController) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
