package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type quotes struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quote []quotes

func main() {

	router := mux.NewRouter()

	seedQuotes()

	router.HandleFunc("/", getQuote)
	router.HandleFunc("/quote", getSingle)
	log.Fatal(http.ListenAndServe(":8000", router))
}
