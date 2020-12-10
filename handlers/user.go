package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/home.html",
		"views/layout/_usernavigation.html",
		"views/layout/_head.html",
		"views/layout/_script.html",
	))

	t.ExecuteTemplate(w, "user_home", nil)
}
