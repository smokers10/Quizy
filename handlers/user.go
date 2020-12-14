package handlers

import (
	"html/template"
	"net/http"
	"quizy/libs"
	"quizy/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/home.html",
		"views/layout/_usernavigation.html",
		"views/layout/_head.html",
		"views/layout/_script.html",
	))
	db := libs.GORM()
	var quiz []models.Quiz
	db.Joins("User").Where("user_id != ?", libs.GetAuth(r).ID).Find(&quiz).Limit(10)
	t.ExecuteTemplate(w, "user_home", libs.M{
		"Quiz": quiz,
		"CSRF": libs.CSRFToken(r),
	})
}
