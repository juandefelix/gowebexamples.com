package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc


func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Method(method string) Middleware{
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if (r.Method != method) {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}
			f(w, r)
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start))}()
			f(w, r)
		}
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}

// what's the signature of the handler function?
// 	- a function that accepts a ResponseWriter and a Request
// what's the signature of the Middleware?
// 	- a function that accetsa a handler function and returns another

