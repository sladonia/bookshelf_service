package app

import (
	"log"
	"net/http"
)

func StartApp() {
	mux := &http.ServeMux{}

	addr := "0.0.0.0:8080"
	log.Printf("start listening on: %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		panic(err)
	}
}
