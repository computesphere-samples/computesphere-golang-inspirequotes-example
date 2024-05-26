package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var quotes = []string{
	"Code is like humor. When you have to explain it, it’s bad. – Cory House",
	"First, solve the problem. Then, write the code. – John Johnson",
	"Experience is the name everyone gives to their mistakes. – Oscar Wilde",
	"In order to be irreplaceable, one must always be different. – Coco Chanel",
	"Java is to JavaScript what car is to Carpet. – Chris Heilmann",
	"Knowledge is power. – Francis Bacon",
	"Sometimes it pays to stay in bed on Monday, rather than spending the rest of the week debugging Monday’s code. – Dan Salomon",
	"Code never lies, comments sometimes do. – Ron Jeffries",
	"Perfection is achieved not when there is nothing more to add, but rather when there is nothing more to take away. – Antoine de Saint-Exupéry",
	"It’s not a bug – it’s an undocumented feature. – Anonymous",
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
