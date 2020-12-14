package handlers

import (
	"html/template"
	"net/http"
	"quizy/libs"
	"quizy/models"

	"github.com/gorilla/mux"
)

type BodyQuiz struct {
	Title      string
	IsPrivate  bool
	PrivateKey string
}

func GetAllQuiz(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/quiz/listquiz.html",
		"views/layout/_head.html",
		"views/layout/_usernavigation.html",
		"views/layout/_script.html",
	))
	db := libs.GORM()
	var quiz []models.SelectedQuiz

	db.Model(&models.Quiz{}).Where("user_id", libs.GetAuth(r).ID).Find(&quiz)

	t.ExecuteTemplate(w, "user_list_quiz", quiz)
}

func CreateQuizPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/quiz/makequiz.html",
		"views/layout/_usernavigation.html",
		"views/layout/_head.html",
		"views/layout/_script.html",
	))

	t.ExecuteTemplate(w, "user_make_quiz", libs.CSRFToken(r))
}

func EditQuiz(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	QuizID := mux.Vars(r)["id"]
	var quiz models.SelectedQuiz

	db.Model(&models.Quiz{}).Where("id = ?", QuizID).Find(&quiz)
	data := libs.M{
		"Data": quiz,
		"CSRF": libs.CSRFToken(r),
	}
	libs.JSON(w, "Edit data retrieved!", data, true)
}

//Action Handlers
func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	body := libs.RDecoder(r)

	db.Create(&models.Quiz{
		Title:      body["title"].(string),
		UserID:     libs.GetAuth(r).ID,
		IsPrivate:  body["isPrivate"].(bool),
		PrivateKey: body["privateKey"].(string),
	})
	libs.JSON(w, "Created new quiz", body, true)
}

func UpdateQuiz(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	quizID := mux.Vars(r)["id"]
	body := libs.RDecoder(r)
	var quiz models.Quiz

	db.Model(quiz).Where("id = ? AND user_id = ?", quizID, libs.GetAuth(r).ID).Updates(models.Quiz{
		Title:     body["name"].(string),
		IsPrivate: body["isPrivate"].(bool),
	})
	libs.JSON(w, "Update quiz", body, true)
}

func ChangeKey(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	body := libs.RDecoder(r)
	quizID := mux.Vars(r)["id"]
	var quiz models.Quiz

	if body["privateKey"].(string) != "" {
		body["privateKey"] = libs.Hashing(body["privateKey"].(string))
	}

	db.Model(quiz).Where("id = ? AND user_id = ?", quizID, libs.GetAuth(r).ID).Update("private_key", body["privateKey"])
	libs.JSON(w, "Quiz key changed", quiz, true)
}

func DeleteQuiz(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	var quizID = mux.Vars(r)["id"]
	var quiz models.Quiz

	db.Where("id = ? AND user_id = ?", quizID, libs.GetAuth(r).ID).Delete(&quiz)
	libs.JSON(w, "Quiz deleted", quiz, true)
}
