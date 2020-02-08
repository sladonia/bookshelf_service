package services

import (
	"bookshelf_service/src/datasources/postgress/bookshelfdb"
	"bookshelf_service/src/domains/author"
)

var (
	AuthorService AuthorServiceInterface = &authorService{}
)

type AuthorServiceInterface interface {
	Create(author author.Author) (*author.Author, error)
	Delete(authorId int64) (*author.Author, error)
	Get(authorId int64) (*author.Author, error)
	Update(author author.Author) (*author.Author, error)
}

type authorService struct{}

func (a *authorService) Create(author author.Author) (*author.Author, error) {
	if err := author.Save(bookshelfdb.Client); err != nil {
		return nil, err
	}
	return &author, nil
}

func (a *authorService) Delete(authorId int64) (*author.Author, error) {
	aut := author.Author{Id: authorId}
	if err := aut.Delete(bookshelfdb.Client); err != nil {
		return nil, err
	}
	return &aut, nil
}

func (a *authorService) Get(authorId int64) (*author.Author, error) {
	aut := author.Author{Id: authorId}
	if err := aut.Retrieve(bookshelfdb.Client); err != nil {
		return nil, err
	}
	return &aut, nil
}

func (a *authorService) Update(author author.Author) (*author.Author, error) {
	if err := author.Update(bookshelfdb.Client); err != nil {
		return nil, err
	}
	return &author, nil
}
