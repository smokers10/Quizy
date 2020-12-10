package routes

import (
	"net/http"
	"quizy/libs"
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
