package main

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	// "os"
	"strings"
)

// Declare all global varibles here
var userdataDB map[string]string
var usernameDB map[string]bool
var userFollowerDB map[string]map[string]bool
var userstateDB map[string]bool

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// strings.Contains(s, substr)
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	} else if strings.Contains(r.URL.Path, "/login") {
		if r.URL.Path == "/login" {
			login(w, r)
			return
		} else {
			userInterface(w, r)
			return
		}
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
	t, err := template.ParseFiles("default.html")
	if err != nil {
		fmt.Fprintf(w, "Error : %v\n", err)
		return
	}
	type data struct {
		Title string
		User  string
	}
	d := data{Title: "I'm a customized router!", User: "Client"}
	// t.Execute(w, "I'm a customized router!")
	// t.Execute(w, "123")
	t.Execute(w, d)
}

func userInterface(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	userName := path[len(path)-1]
	_, userExist := usernameDB[userName]
	_, userLogin := userstateDB[userName]
	if userExist {
		if userLogin {
			if r.Method == "GET" {

				type user struct {
					User string
				}
				userinfo := user{User: userName}
				t, err := template.ParseFiles("userlogin.html")
				if err != nil {
					fmt.Fprintf(w, "Error : %v\n", err)
					return
				}
				t.Execute(w, userinfo)
			} else { // Post
				// Get Form Value
				if err := r.ParseForm(); err != nil {
					fmt.Fprintf(w, "Error: %v\n", err)
					return
				}
				delete(userstateDB, userName)
				http.Redirect(w, r, "/login", http.StatusFound)
				return

			}
		} else {
			if r.Method == "GET" {
				type user struct {
					User string
					URL  string
				}
				userinfo := user{User: userName, URL: r.URL.Path}
				t, err := template.ParseFiles("user_visited.html")
				if err != nil {
					fmt.Fprintf(w, "Error : %v\n", err)
					return
				}
				t.Execute(w, userinfo)
			} else { // Post
				// Get Form Value
				if err := r.ParseForm(); err != nil {
					fmt.Fprintf(w, "Error: %v\n", err)
					return
				}
				val := r.Form.Get("relation")
				if val == "unfollow" {
					fmt.Println("Unfollow", "\n")
				} else {
					fmt.Println("follow", "\n")
				}
				// fmt.Println("Unfollow", r.Form.Get("unfollow"), "\n")
				fmt.Fprintf(w, "Follow or Unfollow")
			}
		}
	} else {
		http.NotFound(w, r)
		fmt.Println("http not found")
		return
	}

}

// func validUserName(uname string) bool {
// 	if m, _ := regexp.MatchString("^[a-zA-Z]+$", uname); !m {
// 		return false
// 	}
// }

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method : ", r.Method)
	// Judge request type
	if r.Method == "GET" {
		// Use certain file as template
		t, err := template.ParseFiles("register.html")
		// Check template name
		fmt.Println(t.Name())
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
			fmt.Print("username :", uName, " password :", pWord, "\n")
			// t.Excute(w, uName, pWord)
			// fmt.Fprintf(w, "Register sucessfully !")
			fmt.Println("Started Redirect !", "\n")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method : ", r.Method)
	var curURL string = r.URL.Path
	// fmt.Println("current URL", curURL, "\n")
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		fmt.Println(t.Name())
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
				userstateDB[uName] = true
				http.Redirect(w, r, curURL+"/"+uName, http.StatusFound)
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
	userstateDB = make(map[string]bool)
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
