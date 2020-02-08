package services

import (
	"bookshelf_service/src/datasources/postgress/bookshelfdb"
	"bookshelf_service/src/domains/books"
)

var (
	AuthorService AuthorServiceInterface = &authorService{}
)

type AuthorServiceInterface interface {
	Create(author books.Author) (*books.Author, error)
}

type authorService struct{}

func (a *authorService) Create(author books.Author) (*books.Author, error) {
	if err := author.Save(bookshelfdb.Client); err != nil {
		return nil, err
	}
	return &author, nil
}
