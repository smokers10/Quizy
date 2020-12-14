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
	var enrolledQuiz []string
	//get enrolled quiz
	db.Raw("SELECT quiz_id FROM enrollments WHERE user_id = ?", libs.GetAuth(r).ID).Scan(&enrolledQuiz)

	if len(enrolledQuiz) != 0 {
		db.Joins("User").Where("user_id != ? AND quizzes.id NOT IN ?", libs.GetAuth(r).ID, enrolledQuiz).Find(&quiz).Limit(10)
		t.ExecuteTemplate(w, "user_home", libs.M{
			"Quiz": quiz,
			"CSRF": libs.CSRFToken(r),
		})
	} else {
		db.Joins("User").Where("user_id != ?", libs.GetAuth(r).ID).Find(&quiz).Limit(5)
		t.ExecuteTemplate(w, "user_home", libs.M{
			"Quiz": quiz,
			"CSRF": libs.CSRFToken(r),
		})
	}
}
