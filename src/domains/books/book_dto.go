package books

type Genre struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	Id          int64   `json:"id"`
	AuthorId    int     `json:"author_id"`
	Title       string  `json:"title"`
	Author      Author  `json:"author"`
	NumberPages int     `json:"number_pages"`
	Genres      []Genre `json:"genres"`
}

type Author struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
