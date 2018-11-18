package main

import (
	// "fmt"
	"auth"
	"profile"
	"twitter"
	"net/http"
)

func main() {
	http.HandleFunc("/", auth.Login)
	http.HandleFunc("/profile", profile.Profile)
	http.HandleFunc("/profile/post", twitter.Twitter)
	http.HandleFunc("/logout", auth.Logout)
	http.ListenAndServe(":9090", nil)
}
