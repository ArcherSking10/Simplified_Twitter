package main

import (
	"auth"
	"auth/storage"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/" {
		// Go to main page
		indexPage(w, r)
	} else {
		uName := auth.GetUserName(r)
		// If uName is in COOKIE,which means the user is login
		// so the username will be returned
		if uName != "" {
			twitter(w, r)
		} else {
			// If user not in COOKIE, the user is not login
			// Then return to the main page
			redirectTarget := "/"
			http.Redirect(w, r, redirectTarget, http.StatusFound)
		}
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method : ", r.Method)
	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("template/index.html")
		if err != nil {
			fmt.Fprintf(w, "Error : %v\n", err)
			return
		}
		t.Execute(w, nil)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error: %v\n", err)
			return
		}
		redirectTarget := "/"
		uName := r.Form.Get("username")
		// Check which form has submitted
		switch submitType := r.Form.Get("submit"); submitType {
		case "Register":
			pWord1 := r.Form.Get("password1")
			pWord2 := r.Form.Get("password2")
			// Check registeration is valid or not
			if ok := storage.WebDB.AddUser(uName, pWord1, pWord2); ok {
				fmt.Println("Register success!")
			} else {
				fmt.Println("Register failed!")
			}
		case "Login":
			pWord := r.Form.Get("password")
			// Check login is valid or not
			if ok := storage.WebDB.HasUser(uName, pWord); ok {
				auth.SetSession(uName, w)
				redirectTarget = "/" + uName
			}
		}
		http.Redirect(w, r, redirectTarget, http.StatusFound)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func twitter(w http.ResponseWriter, r *http.Request) {
	// First Get Login username
	uName := auth.GetUserName(r)
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
			// Put the posts in the Login user's post
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
		case "logout":
			auth.ClearSession(w)
			redirectUrl = "/"
		}
		http.Redirect(w, r, redirectUrl, 302)
	}

}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
