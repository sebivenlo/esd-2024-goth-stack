package main
 
import (
    "fmt"
    "net/http"
    "workshop/views"
 
    "github.com/a-h/templ"
)
 
func main() {
    http.Handle("/", templ.Handler(views.Index()))
    fmt.Println("Server is running on port 8080")
    
    http.ListenAndServe(":8082", nil)
}