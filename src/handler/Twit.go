package handler

import (
	"auth/cookie"
	"auth/storage"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func Twit(w http.ResponseWriter, r *http.Request) {
	uName := cookie.GetUserName(r)
	fmt.Println("twit --------->")
	if uName != "" {
		curUser := storage.WebDB.GetUser(uName)
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("template/post.html")
			if err != nil {
				fmt.Fprintf(w, "Error : %v\n", err)
				return
			}
			t.Execute(w, curUser)
		case "POST":
			r.ParseForm()
			var curTwit storage.TwitPosts
			curTwit.Contents = r.Form.Get("contents")
			// If the post contents are empty not post
			if curTwit.Contents != "" {
				curTwit.Date = time.Now().Unix()
				curTwit.User = uName
				curUser.Posts = append(curUser.Posts, curTwit) // TODO
				// Update the infomation in storage
				storage.WebDB.UpdateUser(uName, curUser)
				// storage.WebDB.UsersInfo[uName] = curUser
				fmt.Println("Posts", curUser.Posts)
			}
			http.Redirect(w, r, "/profile", 302)
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
