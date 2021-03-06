package handlers

import (
	"html/template"
	"net/http"
	"quizy/libs"
	"quizy/models"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateQuestionPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/question/question.html",
		"views/layout/_head.html",
		"views/layout/_usernavigation.html",
		"views/layout/_script.html",
	))
	db, quizID, userID := libs.GORM(), r.URL.Query().Get("id"), libs.GetAuth(r).ID
	var quiz models.Quiz
	var question []models.SelectedQuestion

	db.Select([]string{"Title", "ID"}).Where("id = ? AND user_id", quizID, userID).Find(&quiz)
	db.Model(&models.Question{}).Where("quiz_id = ?", quizID).Find(&question)

	t.ExecuteTemplate(w, "question", libs.M{
		"Quiz":     quiz,
		"CSRF":     libs.CSRFToken(r),
		"Question": question,
	})
}

func EditQuestion(w http.ResponseWriter, r *http.Request) {
	db, id := libs.GORM(), mux.Vars(r)["id"]
	var question models.SelectedQuestion
	db.Model(&models.Question{}).Where("id = ?", id).Find(&question)

	libs.JSON(w, "successfully get selected quiz", question, true)
}

//Actions

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	db, quizID, body := libs.GORM(), r.URL.Query().Get("id"), libs.RDecoder(r)
	conversedQuizID, _ := strconv.Atoi(quizID)
	question := models.Question{
		Question:    body["Question"].(string),
		OptionA:     body["OptionA"].(string),
		OptionB:     body["OptionB"].(string),
		OptionC:     body["OptionC"].(string),
		OptionD:     body["OptionD"].(string),
		RightAnswer: strings.ToUpper(body["RightAnswer"].(string)),
		QuizID:      uint(conversedQuizID),
	}

	db.Create(&question)
	libs.JSON(w, "Created a new question", question, true)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	db, QuestionID, body := libs.GORM(), mux.Vars(r)["id"], libs.RDecoder(r)
	var Question models.Question

	db.Model(Question).Where("id = ?", QuestionID).Updates(models.Question{
		Question:    body["Question"].(string),
		OptionA:     body["OptionA"].(string),
		OptionB:     body["OptionB"].(string),
		OptionC:     body["OptionC"].(string),
		OptionD:     body["OptionD"].(string),
		RightAnswer: strings.ToUpper(body["RightAnswer"].(string)),
	})
	libs.JSON(w, "Updated a new question", Question, true)
}

func DeleteQUestion(w http.ResponseWriter, r *http.Request) {
	db, QuestionID := libs.GORM(), mux.Vars(r)["id"]
	var Question models.Question

	db.Where("id = ?", QuestionID).Delete(&Question)
	libs.JSON(w, "Deleted a new question", Question, true)
}
