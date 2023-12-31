package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}


type TodoPageData struct {
	PageTitle string
	Todos []Todo
}


func main() {
	// The program will try to find the template in the same directory where you run the `go run command,`
	// so make sure that you run this file from the same directory in your terminal
	tmpl := template.Must(template.ParseFiles("layout.html"))
	data := TodoPageData{
		PageTitle: "Today TODOS",
		Todos: []Todo{
			{Title: "Practice Guitar"},
			{Title: "Read Effective Go"},
			{Title: "Exercise"},
			{Title: "Do a tutorial from gowebexamples.com", Done: true},
		},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
