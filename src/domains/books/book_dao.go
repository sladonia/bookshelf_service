package books

type Book struct {
	Id          int64   `json:"id"`
	AuthorId    int     `json:"author_id"`
	Title       string  `json:"title"`
	Author      Author  `json:"author"`
	NumberPages int     `json:"number_pages"`
	Genres      []Genre `json:"genres"`
}
