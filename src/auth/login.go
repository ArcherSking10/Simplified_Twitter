package auth

import (
	"auth/cookie"
	"storage"
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method : ", r.Method)
	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("./template/index.html")
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
				cookie.SetSession(uName, w)
				redirectTarget += "profile"
			}
		}
		http.Redirect(w, r, redirectTarget, http.StatusFound)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
