package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type simpsonsQuote struct {
	Quote              string `json:"quote"`
	Character          string `json:"character"`
	Image              string `json:"image"`
	CharacterDirection string `json:"characterDirection"`
}

func seedQuotes() {
	quote = append(quote, quotes{Quote: "When something is important enough, you do it even if the odds are not in your favor.", Author: "Elon Musk"})
	quote = append(quote, quotes{Quote: "If you get up in the morning and think the future is going to be better, it is a bright day. Otherwise, it's not.", Author: "Elon Musk"})
	quote = append(quote, quotes{Quote: "Persistence is very important. You should not give up unless you are forced to give up.", Author: "Elon Musk"})
	quote = append(quote, quotes{Quote: "The first step is to establish that something is possible; then probability will occur.", Author: "Elon Musk"})
	quote = append(quote, quotes{Quote: "I could either watch it happen or be a part of it.", Author: "Elon Musk"})
}

func getRandom(length int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(length)
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func getSingle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	index := getRandom(len(quote))

	data := &quote[index]
	data.getFirstQuote()
	simpsonsQuote := getSimpsonsQuote()
	combineQuotes := data.combineQuote(simpsonsQuote[0])
	// (*data) += simpsonsQuote[0]["quote"]
	json.NewEncoder(w).Encode(combineQuotes)
}

func getSimpsonsQuote() []map[string]string {
	response, err := http.Get("https://thesimpsonsquoteapi.glitch.me/quotes")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
		return nil
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	// data := []simpsonsQuote{}
	data := []map[string]string{}
	error := json.Unmarshal(contents, &data)
	if error != nil {
		log.Fatal(error)
	}
	// fmt.Println(data[0]["quote"])
	return data
}

func (q *quotes) getFirstQuote() {
	quoteSlice := strings.Split((*q).Quote, " +")
	(*q).Quote = quoteSlice[0]

	authorSlice := strings.Split((*q).Author, " +")
	(*q).Author = authorSlice[0]

}

func (q *quotes) combineQuote(simpsons map[string]string) quotes {
	// s := &simpsons
	fmt.Println((*q).Quote)
	(*q).Quote += " + " + (simpsons)["quote"]
	(*q).Author += " + " + (simpsons)["character"]
	return *q
}
