package handlers

import (
	"html/template"
	"net/http"
	"quizy/libs"
	"quizy/models"

	"github.com/gorilla/mux"
)

func TakeQuiz(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/takequiz.html",
		"views/layout/_usernavigation.html",
		"views/layout/_head.html",
		"views/layout/_script.html",
	))
	db := libs.GORM()
	quizID := mux.Vars(r)["id"]
	var Quiz models.SelectedQuiz
	var Questions []models.SelectedQuestionOnTake

	db.Model(&models.Quiz{}).Where("id = ?", quizID).First(&Quiz)
	db.Model(&models.Question{}).Where("quiz_id = ?", quizID).Find(&Questions)

	t.ExecuteTemplate(w, "taking_quiz", libs.M{
		"Quiz":     Quiz,
		"Question": Questions,
	})
}

func SubmitAnswer(w http.ResponseWriter, r *http.Request) {

}
