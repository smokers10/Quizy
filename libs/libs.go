package libs

import (
	"encoding/json"
	"net/http"
	"quizy/database"

	"gorm.io/gorm"

	"github.com/gorilla/csrf"
	"golang.org/x/crypto/bcrypt"
)

//generate csrf token
func CSRFToken(r *http.Request) M {
	token := M{
		csrf.TemplateTag: csrf.TemplateField(r),
	}

	return token
}

//create hashed string
func Hashing(plaintext string) string {
	plainToByte := []byte(plaintext)
	hashed, e := bcrypt.GenerateFromPassword(plainToByte, 10)
	if e != nil {
		panic(e.Error())
	}
	return string(hashed)
}

//comparing beetwen hashed value and plaintext, will return true if match false if not match
func Compare(hashed string, plaintext string) bool {
	hashedToByte := []byte(hashed)
	plainToByte := []byte(plaintext)

	compared := bcrypt.CompareHashAndPassword(hashedToByte, plainToByte)
	if compared != nil {
		return false
	}
	return true
}

//get logged user data from session
func GetAuth(r *http.Request) LoggedUser {
	session, _ := Store.Get(r, "auth")
	logged := LoggedUser{
		ID:    session.Values["userID"].(uint),
		Email: session.Values["email"].(string),
	}

	return logged
}

//response json for normal operation like CRUD
func JSON(w http.ResponseWriter, message string, data interface{}, success bool) {
	w.Header().Set("Content-Type", "application/json")
	response_data := Response{
		Message: message,
		Data:    data,
		Success: success,
	}

	res, e := json.Marshal(response_data)
	if e != nil {
		panic(e.Error())
	}

	w.Write(res)
}

//response json for authentication operation
func AuthJSON(w http.ResponseWriter, message string, data interface{}, token string, success bool) {
	w.Header().Set("Content-Type", "application/json")
	response_data := Response{
		Message: message,
		Data:    data,
		Token:   token,
		Success: success,
	}

	res, e := json.Marshal(response_data)
	if e != nil {
		panic(e.Error())
	}

	w.Write(res)
}

//To simplefy gorm callng on a handler with one function
func GORM() *gorm.DB {
	gorm := database.Connect()
	return gorm
}

//to decode body
func RDecoder(r *http.Request) map[string]interface{} {
	type EmptyInterface interface{}
	var body EmptyInterface
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	return body.(map[string]interface{})
}
