package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, world.")

	handlerRootStaticFilms := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Test"},
				{Title: "The Godfather2", Director: "Test2"},
				{Title: "The Godfather3", Director: "Test3"},
			},
		}
		tmpl.Execute(w, films)
	}

	http.HandleFunc("/", handlerRootStaticFilms)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
