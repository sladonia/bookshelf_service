package books_controllers

import (
	"github.com/sladonia/log"
	"io"
	"net/http"
)

func GetBook(resp http.ResponseWriter, req *http.Request) {
	log.Debug("getting the book")
	//resp.Header().Add("Content-Type", "application/json")
	resp.Write([]byte(`{"status":"OK"}`))
}

func CreateBook(resp http.ResponseWriter, req *http.Request) {
	log.Debug("getting the book")
	io.WriteString(resp, "creating a book")
}
