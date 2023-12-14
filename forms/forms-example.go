package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
)

type ContactDetails struct {
	Email string
	Subject string
	Message string
}

func main() {
	// The program will try to find the template in the same directory where you run the `go run command,`
	// so make sure that you run this file from the same directory in your terminal
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email: r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		fmt.Println("here")
		log.Print("%v %v %v", details.Email, details.Subject, details.Message)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
