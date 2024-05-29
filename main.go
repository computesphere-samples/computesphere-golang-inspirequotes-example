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
	"Code never lies, comments sometimes do. – Ron Jeffries",
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
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}
