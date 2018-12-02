package main

import (
	"Simplified_Twitter/src/auth"
	"net/http"
	"Simplified_Twitter/src/profile"
	"Simplified_Twitter/src/twitter"
)

func main() {
	http.HandleFunc("/", auth.Login)
	http.HandleFunc("/profile", profile.Profile)
	http.HandleFunc("/profile/post", twitter.Twitter)
	http.HandleFunc("/logout", auth.Logout)
	http.ListenAndServe(":9090", nil)
}
