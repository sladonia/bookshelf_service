package controllers

import (
	"net/http"
)

var (
	BookstoreController BooksControllerInterface = &booksController{}
)

type BooksControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type booksController struct{}

func (b *booksController) Create(w http.ResponseWriter, r *http.Request) {
	err := NewNotImplementedApiError("api not implemented")
	ErrorResponse(w, err)
}

func (b *booksController) Get(w http.ResponseWriter, r *http.Request) {
	type DummyResponse struct {
		BookName string `json:"status"`
	}
	resp := DummyResponse{BookName: "The Goblin Book"}
	JsonResponse(w, http.StatusOK, resp)
}

func (b *booksController) Search(w http.ResponseWriter, r *http.Request) {
	err := NewNotImplementedApiError("api not implemented")
	ErrorResponse(w, err)
}
