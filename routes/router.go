package routes

import (
	"net/http"
	"quizy/handlers"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	//use middleware
	CSRF := csrf.Protect([]byte("handsomedog"), csrf.Secure(false))
	r.Use(CSRF)

	//static files
	file := http.Dir("assets")
	prefix := http.StripPrefix("/assets/", http.FileServer(file))
	r.PathPrefix("/assets/").Handler(prefix)

	//Landing and auth pages routes
	r.HandleFunc("/", handlers.IndexPage).Methods("GET")
	r.HandleFunc("/login", Guest(handlers.LoginPage)).Methods("GET")
	r.HandleFunc("/register", Guest(handlers.RegisterPage)).Methods("GET")

	//auth action
	r.HandleFunc("/register", Guest(handlers.Register)).Methods("POST")
	r.HandleFunc("/login", Guest(handlers.Login)).Methods("POST")
	r.HandleFunc("/logout", Private(handlers.Logout)).Methods("GET")

	//user home
	r.HandleFunc("/user/home", Private(handlers.Home)).Methods("GET")

	//quiz render page
	r.HandleFunc("/user/my-quiz", Private(handlers.GetAllQuiz)).Methods("GET")
	r.HandleFunc("/user/create-quiz", Private(handlers.CreateQuizPage)).Methods("GET")
	r.HandleFunc("/user/edit-quiz/{id}", Private(handlers.EditQuiz)).Methods("GET")

	//quiz user action
	r.HandleFunc("/user/create-quiz", Private(handlers.CreateQuiz)).Methods("POST")
	r.HandleFunc("/user/update-quiz/{id}", Private(handlers.UpdateQuiz)).Methods("POST")
	r.HandleFunc("/user/delete-quiz/{id}", Private(handlers.DeleteQuiz)).Methods("POST")
	r.HandleFunc("/user/change-quiz-key/{id}", Private(handlers.ChangeKey)).Methods("POST")

	//Question render page
	r.HandleFunc("/user/create-question", Private(handlers.CreateQuestionPage)).Methods("GET")
	r.HandleFunc("/user/edit-question/{id}", Private(handlers.EditQuestion)).Methods("GET")

	//Question Action
	r.HandleFunc("/user/create-question", Private(handlers.CreateQuestion)).Methods("POST")
	r.HandleFunc("/user/update-question/{id}", Private(handlers.UpdateQuestion)).Methods("POST")
	r.HandleFunc("/user/delete-question/{id}", Private(handlers.DeleteQUestion)).Methods("POST")

	//Enrollment render page
	r.HandleFunc("/user/my-enrollment", Private(handlers.MyEnrollment)).Methods("GET")

	//Enrollment Action
	r.HandleFunc("/user/enrollment/{id}", Private(handlers.EnrollQuiz)).Methods("POST")
	r.HandleFunc("/user/private-enrollment/{id}", Private(handlers.EnrollPrivateQuiz)).Methods("POST")

	//Taking quiz
	r.HandleFunc("/user/take-quiz/{id}", Private(MustEnroll(handlers.TakeQuiz))).Methods("GET")
	r.HandleFunc("/user/submit-answer", Private(handlers.SubmitAnswer)).Methods("POST")

	return r
}
