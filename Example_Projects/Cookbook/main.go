package main
import (
    "fmt"
    "net/http"
    "Cookbook/views"
 
    "github.com/a-h/templ"
)
 
func main() {
    http.Handle("/", templ.Handler(views.Index()))

    // Start the server

    fmt.Println("Server is running on http://localhost:8080")

    http.ListenAndServe("localhost:8080", nil)
}

