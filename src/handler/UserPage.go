package handler

import (
	"auth/cookie"
	"auth/storage"
	"fmt"
	"html/template"
	"net/http"
	// "time"
)

func UserPage(w http.ResponseWriter, r *http.Request) {
	// First Get Login username
	uName := cookie.GetUserName(r)
	if uName != "" {
		curUser := storage.WebDB.GetUser(uName)
		fmt.Println("Username", curUser)
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("template/twitter.html")
			if err != nil {
				fmt.Fprintf(w, "Error : %v\n", err)
				return
			}
			pageContent := storage.WebDB.GetTwitterPage(uName)
			t.Execute(w, pageContent)
		case "POST":
			r.ParseForm()
			submitType := r.Form.Get("submit")
			fmt.Println(submitType)
			redirectUrl := r.URL.Path
			switch submitType {
			case "follow":
				person := r.Form.Get("unfollow")
				storage.WebDB.FollowUser(uName, person)
			case "unfollow":
				person := r.Form.Get("following")
				storage.WebDB.UnFollowUser(uName, person)
			case "twit":
				redirectUrl += "/post"
			case "logout":
				redirectUrl = "/logout"
				// 	cookie.ClearSession(w)
				// 	redirectUrl = "/"

			}
			http.Redirect(w, r, redirectUrl, 302)
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
