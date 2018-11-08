package main

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	// "strings"
)

var userdata map[string]string

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	} else if r.URL.Path == "/login" {
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

// func validUserName(uname string) bool {
// 	if m, _ := regexp.MatchString("^[a-zA-Z]+$", uname); !m {
// 		return false
// 	}
// }

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

		// Check if it is empty
		if len(uName) == 0 || len(pWord) == 0 /*|| validUserName(uName)*/ {
			fmt.Fprintf(w, "Wrong username format !")
		}
		// Check username and password
		_, ok := userdata[uName]

		if ok {
			if userdata[uName] != pWord {
				fmt.Fprintf(w, "Wrong password, please try again !")
			} else {
				fmt.Fprintf(w, "Login success!")
			}
		} else {
			userdata[uName] = pWord
			fmt.Fprintf(w, "Login success!")
		}

	}
}

func main() {
	userdata = make(map[string]string)
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
