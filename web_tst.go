package main

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	// "strings"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	if r.URL.Path == "/login" {
		login(w, r)
		return
	}
	http.NotFound(w, r)
	fmt.Println("http not found")
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm a customized router!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method : ", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			fmt.Fprintf(w, "Error : %v\n", err)
			return
		}
		t.Execute(w, nil)
	} else { // POST
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error: %v\n", err)
			return
		}
		uName := r.Form.Get("username")
		pWord := r.Form.Get("password")
		fmt.Println("username : ", uName)
		fmt.Println("password : ", pWord)
		if len(uName) == 0 || len(pWord) == 0 {
			fmt.Fprintf(w, "Fields Empty")
		} else {
			fmt.Fprintf(w, "login success!")
		}
	}
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
