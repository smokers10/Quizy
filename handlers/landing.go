package handlers

import (
	"html/template"
	"net/http"
	"quizy/libs"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/index.html",
		"views/layout/_head.html",
		"views/layout/_navigation.html",
		"views/layout/_script.html",
	))

	t.ExecuteTemplate(w, "index", nil)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/register.html",
		"views/layout/_head.html",
		"views/layout/_script.html",
	))

	t.ExecuteTemplate(w, "register", libs.CSRFToken(r))
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/login.html",
		"views/layout/_head.html",
		"views/layout/_script.html",
	))

	t.ExecuteTemplate(w, "login", libs.CSRFToken(r))
}
