package handlers

import (
	"net/http"
	"quizy/libs"
	"quizy/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	body := libs.RDecoder(r)
	var user models.User
	db.Where("email = ?", body["email"].(string)).First(&user)

	if user.ID == 0 {
		libs.JSON(w, "User not registered", nil, false)
		return
	}
	if isMatch := libs.Compare(user.Password, body["password"].(string)); !isMatch {
		libs.JSON(w, "Wrong password entry", nil, false)
		return
	}

	session, _ := libs.Store.Get(r, "auth")
	session.Values["userID"] = user.ID
	session.Values["email"] = user.Email
	if err := session.Save(r, w); err != nil {
		panic(err)
	}

	libs.JSON(w, "Login success", nil, true)
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := libs.GORM()
	body := libs.RDecoder(r)

	db.Create(&models.User{
		Name:     body["name"].(string),
		Email:    body["email"].(string),
		Password: libs.Hashing(body["password"].(string)),
	})

	libs.JSON(w, "Register success", nil, true)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := libs.Store.Get(r, "auth")
	session.Options.MaxAge = -1
	if e := session.Save(r, w); e != nil {
		panic(e.Error())
	}

	http.Redirect(w, r, "/login", 302)
}
