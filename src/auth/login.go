package auth

import (
	"Simplified_Twitter/src/auth/cookie"
	"fmt"
	"html/template"
	"net/http"
	"Simplified_Twitter/src/rpc/client"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method : ", r.Method)
	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("./src/template/index.html")
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
			if ok := client.RpcAddUser(uName, pWord1, pWord2); ok {
				fmt.Println("Register success!")
			} else {
				fmt.Println("Register failed!")
			}
		case "Login":
			pWord := r.Form.Get("password")
			// Check login is valid or not
			if ok := client.RpcHasUser(uName, pWord); ok {
				cookie.SetSession(uName, w)
				redirectTarget += "profile"
			}
		}
		http.Redirect(w, r, redirectTarget, http.StatusFound)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
