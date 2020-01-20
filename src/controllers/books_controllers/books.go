package books_controllers

import (
	"github.com/sladonia/log"
	"io"
	"net/http"
)

func GetBook(resp http.ResponseWriter, req *http.Request) {
	log.Debug("getting the book")
	io.WriteString(resp, "here is your book, please")
}
