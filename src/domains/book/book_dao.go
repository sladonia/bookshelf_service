package book

import (
	"bookshelf_service/src/domains/author"
	"bookshelf_service/src/domains/genre"
)

type Book struct {
	Id          int64         `json:"id"`
	AuthorId    int           `json:"author_id"`
	Title       string        `json:"title"`
	Author      author.Author `json:"author"`
	NumberPages int           `json:"number_pages"`
	Genres      []genre.Genre `json:"genres"`
}
