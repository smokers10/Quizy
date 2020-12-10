package handlers

import (
	"html/template"
	"net/http"
	"quizy/libs"
	"quizy/models"
)

func CreateQuestionPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/question/question.html",
		"views/layout/_head.html",
		"views/layout/_usernavigation.html",
		"views/layout/_script.html",
	))
	db := libs.GORM()
	quizID := r.URL.Query().Get("id")
	userID := libs.GetAuth(r).ID
	var quiz models.Quiz

	db.Select([]string{"QuizName", "ID"}).Where("id = ? AND user_id", quizID, userID).Find(&quiz)

	t.ExecuteTemplate(w, "question", quiz)
}

func EditQuestion(w http.ResponseWriter, r *http.Request) {

}

//Action handler

func CreateQuestion(w http.ResponseWriter, r *http.Request) {

}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {

}

func DeleteQUestion(w http.ResponseWriter, r *http.Request) {

}
