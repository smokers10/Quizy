package handlers

import (
	"net/http"
	"quizy/database"
	"quizy/libs"
	"quizy/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	gorm := database.Connect()
	db, _ := gorm.DB()
	defer db.Close()

	var user models.User

	email := r.FormValue("email")
	password := r.FormValue("password")

	gorm.Where("email", email).First(&user)
	if user.ID == 0 {
		http.Redirect(w, r, "/login", 302)
	}

	isMatch := libs.Compare(user.Password, password)
	if !isMatch {
		http.Redirect(w, r, "/login", 302)
	}

	session, _ := libs.Store.Get(r, "auth")
	session.Values["userID"] = user.ID
	session.Values["email"] = user.Email
	if e := session.Save(r, w); e != nil {
		panic(e)
	}

	http.Redirect(w, r, "/user/home", 302)
}

func Register(w http.ResponseWriter, r *http.Request) {
	gorm := database.Connect()
	db, _ := gorm.DB()
	defer db.Close()

	gorm.Create(&models.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: libs.Hashing(r.FormValue("password")),
	})

	http.Redirect(w, r, "/login", 302)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := libs.Store.Get(r, "auth")
	session.Options.MaxAge = -1
	if e := session.Save(r, w); e != nil {
		panic(e.Error())
	}

	http.Redirect(w, r, "/login", 302)
}
