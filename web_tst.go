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

// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
// 	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
// 	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("method:", r.Method) //获取请求的方法
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("login.html")
// 		log.Println(t.Execute(w, nil))
// 	} else {
// 		//请求的是登录数据，那么执行登录的逻辑判断
// 		r.ParseForm()
// 		fmt.Println("username:", r.Form["username"])
// 		fmt.Println("password:", r.Form["password"])
// 	}
// }

// func main() {
// 	http.HandleFunc("/", sayhelloName)       //设置访问的路由
// 	http.HandleFunc("/login", login)         //设置访问的路由
// 	err := http.ListenAndServe(":9090", nil) //设置监听的端口
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
