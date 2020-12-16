package handlers

import (
	"net/http"
	"quizy/libs"
	"quizy/models"
	"strconv"

	"github.com/gorilla/mux"
)

func SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	db, quizID, body := libs.GORM(), mux.Vars(r)["id"], libs.ArrayRDecoder(r)
	quizIDFormated, _ := strconv.Atoi(quizID)
	var enrollment models.Enrollment
	var enrollmentCheck models.Enrollment

	db.Where("user_id = ? AND quiz_id AND is_submited", libs.GetAuth(r).ID, quizIDFormated).First(&enrollmentCheck)
	if enrollmentCheck.ID != 0 {
		libs.JSON(w, "This quiz already submitted", nil, true)
		return
	}

	for _, v := range body {
		questionID := v.(map[string]interface{})["name"].(string)
		answer := v.(map[string]interface{})["value"]
		if questionID == "gorilla.csrf.Token" {
			continue
		}
		formatedQuestionID, _ := strconv.Atoi(questionID)
		db := libs.GORM()
		var question models.SelectedQuestionOnCheck
		db.Model(&models.Question{}).Where("id = ?", questionID).First(&question)

		if question.RightAnswer == answer.(string) {
			db.Create(&models.Answer{
				QuizId:     uint(quizIDFormated),
				QuestionID: uint(formatedQuestionID),
				Answer:     answer.(string),
				IsRight:    true,
				UserID:     libs.GetAuth(r).ID,
			})
		} else {
			db.Create(&models.Answer{
				QuizId:     uint(quizIDFormated),
				QuestionID: uint(formatedQuestionID),
				Answer:     answer.(string),
				IsRight:    false,
				UserID:     libs.GetAuth(r).ID,
			})
		}
	}

	db.Model(enrollment).Where("user_id = ? AND quiz_id = ?", libs.GetAuth(r).ID, quizIDFormated).Updates(models.Enrollment{
		IsSubmited: true,
	})
	libs.JSON(w, "Submited successfully", nil, true)
}
