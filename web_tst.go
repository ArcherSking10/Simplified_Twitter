package main

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	// "strings"
)

// Declare all global varibles here
var userdataDB map[string]string
var usernameDB map[string]bool

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	} else if r.URL.Path == "/login" {
		login(w, r)
		return
	} else if r.URL.Path == "/register" {
		register(w, r)
		return
	} else {
		http.NotFound(w, r)
		fmt.Println("http not found")
		return
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm a customized router!")
}

// func validUserName(uname string) bool {
// 	if m, _ := regexp.MatchString("^[a-zA-Z]+$", uname); !m {
// 		return false
// 	}
// }

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method : ", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("register.html")
		if err != nil {
			fmt.Fprintf(w, "Error : %v\n", err)
			return
		}
		t.Execute(w, nil)
	} else { // Post
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error: %v\n", err)
			return
		}
		uName := r.Form.Get("username")
		pWord := r.Form.Get("password")
		// Check whether it is empty
		if len(uName) == 0 || len(pWord) == 0 /*|| validUserName(uName)*/ {
			if len(uName) == 0 {
				fmt.Fprintf(w, "Please enter username !")
				return
			} else {
				fmt.Fprintf(w, "Please enter passwords !")
				return
			}
		}
		fmt.Println("username", uName)
		_, ok := usernameDB[uName]
		fmt.Println("user", ok)
		if ok {
			fmt.Fprintf(w, "Username existed !")
			return
		} else {
			usernameDB[uName] = true
			userdataDB[uName] = pWord
			fmt.Print("username :", uName, " password :", pWord)
			fmt.Fprintf(w, "Register sucessfully !")
			return
		}

	}
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

		// Check username and password
		_, ok := userdataDB[uName]

		if ok {
			if userdataDB[uName] != pWord {
				fmt.Fprintf(w, "Wrong password, please try again !")
				return
			} else {
				fmt.Fprintf(w, "Login success!")
				return
			}
		} else {
			fmt.Fprintf(w, "User not exist !")
			return
		}

	}
}

func main() {
	userdataDB = make(map[string]string)
	usernameDB = make(map[string]bool)
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
