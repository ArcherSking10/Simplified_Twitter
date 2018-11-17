package handler

import (
	"auth/cookie"
	// "auth/storage"
	// "fmt"
	// "html/template"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie.ClearSession(w)
	http.Redirect(w, r, "/", 302)
}
