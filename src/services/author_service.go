package services

import "bookshelf_service/src/domains/books"

var (
	AuthorService AuthorServiceInterface = &authorService{}
)

type AuthorServiceInterface interface {
	Create(firstName, lastName string) (*books.Author, error)
}

type authorService struct{}

func (a *authorService) Create(firstName, lastName string) (*books.Author, error) {
	author := &books.Author{
		FirstName: firstName,
		LastName:  lastName,
	}
	if err := author.CleanAndValidate(); err != nil {
		return nil, err
	}
	if err := author.Save(); err != nil {
		return nil, err
	}
	return author, nil
}
