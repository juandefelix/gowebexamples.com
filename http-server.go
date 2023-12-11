package main

import (
	"fmt"
	"net/http"
)

func main() {
	// the folder is relative to the current directory
	fs := http.FileServer(http.Dir("static/"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<script src=\"http://localhost/static/alert.js\"></script><h1>Hola Manola</h1>")
	})
	http.Handle("/static/alert.js", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
