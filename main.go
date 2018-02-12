package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Quotes struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotes []Quotes

func main() {

	router := mux.NewRouter()

	quotes = append(quotes, Quotes{Quote: "When something is important enough, you do it even if the odds are not in your favor.", Author: "Elon Musk"})
	quotes = append(quotes, Quotes{Quote: "If you get up in the morning and think the future is going to be better, it is a bright day. Otherwise, it's not.", Author: "Elon Musk"})
	quotes = append(quotes, Quotes{Quote: "Persistence is very important. You should not give up unless you are forced to give up.", Author: "Elon Musk"})
	quotes = append(quotes, Quotes{Quote: "The first step is to establish that something is possible; then probability will occur.", Author: "Elon Musk"})
	quotes = append(quotes, Quotes{Quote: "I could either watch it happen or be a part of it.", Author: "Elon Musk"})

	fmt.Println(getRandom(len(quotes)))
	router.HandleFunc("/", getQuote)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getRandom(length int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(length)
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}
