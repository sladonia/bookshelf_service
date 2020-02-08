package controllers

import (
	"bookshelf_service/src/domains/author"
	"bookshelf_service/src/domains/responses"
	"bookshelf_service/src/logger"
	"bookshelf_service/src/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	AuthorController AuthorControllerInterface = &authorController{}
)

type AuthorControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type authorController struct {
}

func (a *authorController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMsg := "invalid body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}
	defer r.Body.Close()
	var aut author.Author
	err = json.Unmarshal(requestBody, &aut)
	if err != nil {
		errorMsg := "invalid json body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}

	result, err := services.AuthorService.Create(aut)
	if err != nil {
		errorMsg := "unable to crate author"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
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
	vars := mux.Vars(r)
	authorId, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		errorMsg := "invalid request"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}

	result, err := services.AuthorService.Get(authorId)
	if err != nil {
		errorMsg := "unable to get author"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		ErrorResponse(w, apiErr)
		return
	}
	JsonResponse(w, http.StatusCreated, result)
}

func (a *authorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		errorMsg := "invalid request"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}

	result, err := services.AuthorService.Delete(authorId)
	if err != nil {
		errorMsg := fmt.Sprintf("unable to delete author id=%d", authorId)
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		ErrorResponse(w, apiErr)
		return
	}
	response := responses.ResponseDeleted{
		Message:   "author deleted",
		DeletedId: result.Id,
	}
	JsonResponse(w, http.StatusAccepted, response)
}

func (a *authorController) Update(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMsg := "invalid body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}
	defer r.Body.Close()
	var aut author.Author
	err = json.Unmarshal(requestBody, &aut)
	if err != nil {
		errorMsg := "invalid json body"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}
	vars := mux.Vars(r)
	authorId, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		errorMsg := "invalid request"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewBadRequestApiError(errorMsg)
		ErrorResponse(w, apiErr)
		return
	}

	aut.Id = authorId
	result, err := services.AuthorService.Update(aut)
	if err != nil {
		errorMsg := "unable to update"
		logger.Logger.Infow(errorMsg, "err", err.Error(), "path", r.URL.Path)
		apiErr := NewApiError(errorMsg, err.Error(), http.StatusConflict)
		ErrorResponse(w, apiErr)
		return
	}
	response := responses.ResponseCreated{
		Message:   "author updated",
		CreatedId: result.Id,
	}
	JsonResponse(w, http.StatusAccepted, response)
}
