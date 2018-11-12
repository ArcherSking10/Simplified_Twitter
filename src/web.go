package main

import (
	"auth"
	"auth/storage"
	"fmt"
	"html/template"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// Go to main page
		indexPage(w, r)
		return
	} else {
		uName := auth.GetUserName(r)
		// If uName is in COOKIE,which means the user is login
		// so the username will be returned
		if uName != "" {
			// TODO: go to twitter user page
			twitter(w, r)
			fmt.Println("User")
			return
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
		// Check which form is submited
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
	fmt.Println("twitter Page", uName, "\n")
	if uName != "" {
		// Get the corresponding User struct
		curUser := storage.WebDB.UsersInfo[uName]
		fmt.Println("Username", curUser)
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("template/twitter.html")
			if err != nil {
				fmt.Fprintf(w, "Error : %v\n", err)
				return
			}
			t.Execute(w, curUser)
		case "POST":
			r.ParseForm()
			// Put the posts in the Login user's post
			curUser.Posts = append(curUser.Posts, r.Form.Get("contents"))
			// Update the infomation in storage
			storage.WebDB.UsersInfo[uName] = curUser
			fmt.Println("Posts", curUser.Posts)
			http.Redirect(w, r, r.URL.Path, 302)
		}
	}

}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}