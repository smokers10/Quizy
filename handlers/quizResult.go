package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"quizy/libs"
	"quizy/models"

	"github.com/gorilla/mux"
)

func CalculateQuizresult(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/quiz/result.html",
		"views/layout/_usernavigation.html",
		"views/layout/_head.html",
		"views/layout/_script.html",
	))
	db, quizID, userID := libs.GORM(), mux.Vars(r)["id"], libs.GetAuth(r).ID
	var Answer []models.Answer

	db.Joins("Question").Where("answers.quiz_id  = ? AND answers.user_id = ?", quizID, userID).Find(&Answer)
	fmt.Println(Answer)
	t.ExecuteTemplate(w, "quiz_result", Answer)
}
