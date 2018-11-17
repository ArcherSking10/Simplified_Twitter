package main

import (
	// "fmt"
	"handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Auth)
	http.HandleFunc("/profile", handler.UserPage)
	http.HandleFunc("/profile/post", handler.Twit)
	http.HandleFunc("/logout", handler.Logout)
	http.ListenAndServe(":9090", nil)
}
