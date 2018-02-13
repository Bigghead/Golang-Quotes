package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type simpsonsQuote struct {
	Quote              string `json:"quote"`
	Character          string `json:"character"`
	Image              string `json:"image"`
	CharacterDirection string `json:"direction"`
}

func seedQuotes() {
	quotes = append(quotes, Quotes{Quote: "When something is important enough, you do it even if the odds are not in your favor.", Author: "Elon Musk"})
	quotes = append(quotes, Quotes{Quote: "If you get up in the morning and think the future is going to be better, it is a bright day. Otherwise, it's not.", Author: "Elon Musk"})
	quotes = append(quotes, Quotes{Quote: "Persistence is very important. You should not give up unless you are forced to give up.", Author: "Elon Musk"})
	quotes = append(quotes, Quotes{Quote: "The first step is to establish that something is possible; then probability will occur.", Author: "Elon Musk"})
	quotes = append(quotes, Quotes{Quote: "I could either watch it happen or be a part of it.", Author: "Elon Musk"})
}

func getRandom(length int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(length)
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func getSingle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	index := getRandom(len(quotes))
	json.NewEncoder(w).Encode(quotes[index])
}

func getSimpsonsQuote() {
	response, err := http.Get("https://thesimpsonsquoteapi.glitch.me/quotes")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", string(contents))
		// fmt.Printf("%s",)
		// fmt.Println(json.NewDecoder(response.Body).Decode(&data))
		// return json.NewDecoder(response.Body).Decode(&data)
	}
}
