package main
 
import (
    "net/http"
    "workshop/views"
 
    "github.com/a-h/templ"
)
 
func main() {
    http.Handle("/", templ.Handler(views.Index()))

    http.ListenAndServe("localhost:8080", nil)
}