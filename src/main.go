package main

import (
	"Simplified_Twitter/src/auth"
	"Simplified_Twitter/src/handle"
	"Simplified_Twitter/src/profile"
	"Simplified_Twitter/src/twitter"
	"net/http"
)

func main() {
	http.HandleFunc("/", auth.Login)
	http.HandleFunc("/profile", profile.Profile)
	http.HandleFunc("/profile/post", twitter.Twitter)
	http.HandleFunc("/logout", auth.Logout)
	http.HandleFunc("/join", handle.Join)
	http.ListenAndServe(":9090", nil)
}
