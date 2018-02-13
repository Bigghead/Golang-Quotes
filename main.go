package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Quotes struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotes []Quotes

func main() {

	router := mux.NewRouter()

	seedQuotes()
	getSimpsonsQuote()

	router.HandleFunc("/", getQuote)
	router.HandleFunc("/quote", getSingle)
	log.Fatal(http.ListenAndServe(":8000", router))
}
