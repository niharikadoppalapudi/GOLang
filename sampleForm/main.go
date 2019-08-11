package main

import (
	"net/http"
	"text/template"
)

type ContactDetails struct {
	Email   string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Message: r.FormValue("Hello world"),
		}
		_ = details
		tmpl.Execute(w, struct{ Success bool }{true})
	})
	http.ListenAndServe(":8002", nil)
}
