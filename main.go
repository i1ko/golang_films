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

	handlerFormDataAppend := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)

		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", handlerRootStaticFilms)
	http.HandleFunc("/add-film/", handlerFormDataAppend)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
