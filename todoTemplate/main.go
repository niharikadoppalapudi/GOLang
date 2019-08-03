package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		fmt.Println("Template File not found.")
	}
	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
