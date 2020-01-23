package books

//const (
//	UnknownGenre = iota
//	Fantasy
//	ScienceFiction
//	Horror
//	Western
//	Romance
//	Thriller
//	Mystery
//	Detective
//	Dystopia
//	Newspaper
//)

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	Id          int     `json:"id"`
	AuthorId    int     `json:"author_id"`
	Title       string  `json:"title"`
	Author      Author  `json:"author"`
	NumberPages int     `json:"number_pages"`
	Genres      []Genre `json:"genres"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
