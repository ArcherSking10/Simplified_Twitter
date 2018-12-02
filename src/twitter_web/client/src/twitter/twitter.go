package twitter

import (
	"auth/cookie"
	"fmt"
	// pb "google.golang.org/grpc/examples/Simplified_Twitter/src/twitter_web/TwitterPage"
	"html/template"
	"net/http"
	"rpcFunction"
	"storage"
	"time"
)

func Twitter(w http.ResponseWriter, r *http.Request) {
	uName := cookie.GetUserName(r)
	if uName != "" {
		fmt.Println("----------------> Test rpc Start")
		curUser := rpcFunction.RpcGetUser(uName)
		fmt.Println("----------------> Test rpc End")
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("template/post.html")
			if err != nil {
				fmt.Fprintf(w, "Error : %v\n", err)
				return
			}
			t.Execute(w, curUser)
		case "POST":
			r.ParseForm()
			var curTwit = storage.TwitPosts{}
			curTwit.Contents = r.Form.Get("contents")
			// // If the post contents are empty not post
			if curTwit.Contents != "" {
				curTwit.Date = time.Now().Unix()
				curTwit.User = uName
				curUser.Posts = append(curUser.Posts, curTwit) // TODO

				rpcFunction.RpcUpdateUser(uName, curUser)
				fmt.Println("Posts", curUser.Posts)
			}
			http.Redirect(w, r, "/profile", 302)
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
