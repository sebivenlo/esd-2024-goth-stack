package main

import (
	"net/http"
	"workshop/views"

	"github.com/a-h/templ"
)

var todos = []views.Todo{
	{Id: 1, Title: "Buy milk", Completed: false},
	{Id: 2, Title: "Buy eggs", Completed: true},
}

func main() {
	http.Handle("/", templ.Handler(views.Index()))
	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/todos/", handleTodo)

	http.ListenAndServe("localhost:8080", nil)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		views.TodoList(todos).Render(r.Context(), w)
	case http.MethodPost:
		title := r.FormValue("title")
		newTodo := views.Todo{
			Id:        len(todos) + 1,
			Title:     title,
			Completed: false,
		}
		todos = append([]views.Todo{newTodo}, todos...)
		views.TodoItemOob(newTodo, "afterbegin").Render(r.Context(), w)
	}
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		w.Write([]byte(""))
	}
}
