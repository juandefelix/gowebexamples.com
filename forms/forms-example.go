package main

import (
	"html/template"
	"net/http"
	"log"
)

type ContactDetails struct {
	Email string
	Subject string
	Message string
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL.Path)
		f(w,r)
	}
}

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := ContactDetails{
		Email: r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	log.Println("Form values: ", details.Email, details.Subject, details.Message)

	tmpl.Execute(w, struct{ Success bool }{true})
}

func main() {
	// The program will try to find the template in the same directory where you run the `go run command,`
	// so make sure that you run this file from the same directory in your terminal

	http.HandleFunc("/", logging(rootPathHandler))

	http.ListenAndServe(":8080", nil)
}
