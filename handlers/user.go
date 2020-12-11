package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"quizy/libs"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/home.html",
		"views/layout/_usernavigation.html",
		"views/layout/_head.html",
		"views/layout/_script.html",
	))

	fmt.Println(libs.GetAuth(r))

	t.ExecuteTemplate(w, "user_home", nil)
}
