package routes

import (
	"net/http"
	"quizy/libs"
	"quizy/models"

	"github.com/gorilla/mux"
)

func Private(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, e := libs.Store.Get(r, "auth")
		if e != nil {
			panic(e.Error())
		}
		if len(session.Values) == 0 {
			http.Redirect(w, r, "/login", 302)
			return
		}

		handler.ServeHTTP(w, r)
	}
}

func Guest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, e := libs.Store.Get(r, "auth")
		if e != nil {
			panic(e.Error())
		}
		if len(session.Values) != 0 {
			http.Redirect(w, r, "/user/home", 302)
			return
		}

		handler.ServeHTTP(w, r)
	}
}

func MustEnroll(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get user id
		session, e := libs.Store.Get(r, "auth")
		db := libs.GORM()
		if e != nil {
			panic(e.Error())
		}
		userID := session.Values["userID"]
		quizID := mux.Vars(r)["id"]
		var enrollment models.EnrollmentOnCheck
		db.Model(&models.Enrollment{}).Where("user_id = ? AND quiz_id = ?", userID, quizID).First(&enrollment)
		//checking process
		if enrollment.ID == 0 {
			http.Redirect(w, r, "/user/home", 302)
		}
		handler.ServeHTTP(w, r)
	}
}
