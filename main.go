package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var quotes = []string{
	"The best way to predict the future is to invent it.",
	"Life is 10% what happens to us and 90% how we react to it.",
	"The only way to do great work is to love what you do.",
	"You miss 100% of the shots you don’t take.",
	"The only limit to our realization of tomorrow is our doubts of today.",
}

func getRandomQuote() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return quotes[r.Intn(len(quotes))]
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	quote := getRandomQuote()
	w.Write([]byte(quote))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, getRandomQuote())
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/quote", quoteHandler)
	http.ListenAndServe(":8080", nil)
}
