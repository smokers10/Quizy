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

	//route list
	r.HandleFunc("/", handlers.IndexPage).Methods("GET")
	r.HandleFunc("/login", Guest(handlers.LoginPage)).Methods("GET")
	r.HandleFunc("/register", Guest(handlers.RegisterPage)).Methods("GET")

	//auth action
	r.HandleFunc("/register", Guest(handlers.Register)).Methods("POST")
	r.HandleFunc("/login", Guest(handlers.Login)).Methods("POST")
	r.HandleFunc("/logout", Private(handlers.Logout)).Methods("GET")

	//user routes
	r.HandleFunc("/user/home", Private(handlers.Home)).Methods("GET")

	//quiz render page
	r.HandleFunc("/user/my-quiz", Private(handlers.GetAllQuiz)).Methods("GET")
	r.HandleFunc("/user/create-quiz", Private(handlers.CreateQuizPage)).Methods("GET")
	r.HandleFunc("/user/edit-quiz", Private(handlers.EditQuiz)).Methods("GET")

	//quiz user action
	r.HandleFunc("/user/create-quiz", Private(handlers.CreateQuiz)).Methods("POST")
	r.HandleFunc("/user/update-quiz", Private(handlers.UpdateQuiz)).Methods("POST")
	r.HandleFunc("/user/delete-quiz", Private(handlers.DeleteQuiz)).Methods("POST")
	r.HandleFunc("/user/change-quiz-key", Private(handlers.ChangeKey)).Methods("POST")

	//Question render page
	r.HandleFunc("/user/create-question", Private(handlers.CreateQuestionPage)).Methods("GET")
	r.HandleFunc("/user/edit-question", Private(handlers.EditQuestion)).Methods("GET")

	//Question Action
	r.HandleFunc("/user/create-question", Private(handlers.CreateQuestion)).Methods("POST")
	r.HandleFunc("/user/update-question", Private(handlers.UpdateQuestion)).Methods("POST")
	r.HandleFunc("/user/delete-question", Private(handlers.DeleteQUestion)).Methods("POST")

	return r
}
