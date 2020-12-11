package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"quizy/libs"
	"quizy/models"
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
	QuizID := r.URL.Query().Get("id")
	var quiz models.SelectedQuiz

	db.Model(&models.Quiz{}).Where("id", QuizID).Find(&quiz)

	data := libs.M{
		"Data": quiz,
		"CSRF": libs.CSRFToken(r),
	}

	libs.JSON(w, "Edit data retrieved!", data, true)
}

//Action Handlers
func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()

	var body BodyQuiz
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if body.PrivateKey != "" {
		body.PrivateKey = libs.Hashing(body.PrivateKey)
	}

	quiz := models.Quiz{
		Title:      body.Title,
		UserID:     libs.GetAuth(r).ID,
		IsPrivate:  body.IsPrivate,
		PrivateKey: body.PrivateKey,
	}

	db.Create(&quiz)

	libs.JSON(w, "Created new quiz", body, true)
}

func UpdateQuiz(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	var body BodyQuiz
	var quizID = r.URL.Query().Get("id")
	var quiz models.Quiz

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	db.Model(quiz).Where("id = ?", quizID).Where("use_id = ?", libs.GetAuth(r).ID).Updates(models.Quiz{
		Title:     body.Title,
		IsPrivate: body.IsPrivate,
	})

	libs.JSON(w, "Update quiz", body, true)
}

func ChangeKey(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	var body BodyQuiz
	var quizID = r.URL.Query().Get("id")
	var quiz models.Quiz
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	if body.PrivateKey != "" {
		body.PrivateKey = libs.Hashing(body.PrivateKey)
	}

	db.Model(quiz).Where("id = ?", quizID).Where("use_id = ?", libs.GetAuth(r).ID).Update("private_key", body.PrivateKey)

	libs.JSON(w, "Quiz key changed", body, true)
}

func DeleteQuiz(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	var quizID = r.URL.Query().Get("id")
	var quiz models.Quiz

	db.Where("id = ?", quizID).Where("use_id = ?", libs.GetAuth(r).ID).Delete(&quiz)

	libs.JSON(w, "Quiz deleted", quiz, true)
}
