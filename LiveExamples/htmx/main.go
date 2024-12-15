package main

import (
	"net/http"
	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(Component()))

	http.Handle("/greeting", templ.Handler(Response()))

	http.ListenAndServe("localhost:8080", nil)
}