package main

import (
	"fmt"
	"html/template"
	"auth/storage"
	"net/http"
	"auth"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        indexPage(w, r)
    } else {
    	fmt.Println(r.URL.Path)
        uName := auth.GetUserName(r)
		if uName != "" {
			// TODO: go to twitter page
			fmt.Println("haha")
		} else {
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
		switch submitType := r.Form.Get("submit"); submitType {
		case "Register":
			pWord1 := r.Form.Get("password1")
			pWord2 := r.Form.Get("password2")
			if ok := storage.WebDB.AddUser(uName, pWord1, pWord2); ok {
				fmt.Println("Register success!")
			} else {
				fmt.Println("Register failed!")
			}
		case "Login":
			pWord := r.Form.Get("password")
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

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}