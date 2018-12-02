package main

import (
	"auth"
	"net/http"
	"profile"
	"twitter"
)

func main() {
	http.HandleFunc("/", auth.Login)
	http.HandleFunc("/profile", profile.Profile)
	http.HandleFunc("/profile/post", twitter.Twitter)
	http.HandleFunc("/logout", auth.Logout)
	http.ListenAndServe(":9090", nil)
}
