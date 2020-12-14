package handlers

import (
	"html/template"
	"net/http"
	"quizy/libs"
	"quizy/models"
	"strconv"

	"github.com/gorilla/mux"
)

func MyEnrollment(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"views/my-enrollment.html",
		"views/layout/_usernavigation.html",
		"views/layout/_head.html",
		"views/layout/_script.html",
	))
	db := libs.GORM()
	var enrollments []models.Enrollment
	db.Joins("Quiz").Where("enrollments.user_id = ?", libs.GetAuth(r).ID).Find(&enrollments)

	t.ExecuteTemplate(w, "my-enrollment", enrollments)
}

func EnrollQuiz(w http.ResponseWriter, r *http.Request) {
	db, userID := libs.GORM(), libs.GetAuth(r).ID
	quizID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var quiz models.SelectedQuizEnrollment

	db.Model(&models.Quiz{}).Where("id = ?", quizID, userID).First(&quiz)
	if quiz.ID == 0 {
		libs.JSON(w, "Quiz not found", nil, false)
		return
	}
	if quiz.IsPrivate {
		libs.JSON(w, "Quiz is private, please call creator to get key", nil, false)
		return
	}

	enrollment := models.Enrollment{
		UserID: userID,
		QuizID: uint(quizID),
	}
	db.Create(&enrollment)

	libs.JSON(w, "Successfully Enrolled", quiz.ID, true)
	return
}

func EnrollPrivateQuiz(w http.ResponseWriter, r *http.Request) {
	db, userID, body := libs.GORM(), libs.GetAuth(r).ID, libs.RDecoder(r)
	quizID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var quiz models.SelectedQuizEnrollment

	db.Model(&models.Quiz{}).Where("id = ?", quizID, userID).First(&quiz)
	if quiz.ID == 0 {
		libs.JSON(w, "Quiz not found", nil, false)
		return
	}
	if !quiz.IsPrivate {
		libs.JSON(w, "Quiz is public access!", nil, false)
		return
	}
	if isMatch := libs.Compare(quiz.PrivateKey, body["privatekey"].(string)); !isMatch {
		libs.JSON(w, "Wrong key, ask creator for right key", nil, false)
		return
	}

	enrollment := models.Enrollment{
		UserID: userID,
		QuizID: uint(quizID),
	}
	db.Create(&enrollment)

	libs.JSON(w, "Successfully enrolled", quiz.ID, true)
	return
}

func Unenrollment(w http.ResponseWriter, r *http.Request) {
	db, enrollmentID, userID := libs.GORM(), mux.Vars(r)["id"], libs.GetAuth(r).ID
	var Enrollment models.Enrollment

	db.Where("id = ? AND user_id", enrollmentID, userID).Delete(&Enrollment)
}
