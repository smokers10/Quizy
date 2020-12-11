package libs

import (
	"github.com/gorilla/sessions"
)

//shorthand to map interface{} with index string
type M map[string]interface{}

//Normal sessions
var Store = sessions.NewCookieStore([]byte("handsome_dog"))

//logged user
type LoggedUser struct {
	ID    uint
	Email string
}

//For formating REST API response
type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
	Success bool        `json:"success"`
}
