package profile

import (
	"Simplified_Twitter/src/auth/cookie"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"Simplified_Twitter/src/rpc/client"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	// First Get Login username
	uName := cookie.GetUserName(r)
	if uName != "" {
		// fmt.Println("----------------> Test rpc Start")
		curUser := client.RpcGetUser(uName)
		// fmt.Println("----------------> Test rpc End")
		fmt.Println("Username", curUser.UserName)
		switch r.Method {
		case "GET":
			t, err := template.ParseFiles("./src/template/twitter.html")
			if err != nil {
				fmt.Fprintf(w, "Error : %v\n", err)
				return
			}
			fmt.Println("----------------> Test TwitterPage Start")
			// pageContent := storage.WebDB.GetTwitterPage(uName)
			pageContent := client.RpcGetTwitterPage(uName)
			log.Printf("........pagecontent1", pageContent.UserName)
			log.Printf("........pagecontent2", pageContent.Following)
			log.Printf("........pagecontent3", pageContent.UnFollowed)
			log.Printf("........pagecontent4", pageContent.Posts)

			t.Execute(w, pageContent)
		case "POST":
			r.ParseForm()
			submitType := r.Form.Get("submit")
			fmt.Println(submitType)
			redirectUrl := r.URL.Path
			switch submitType {
			case "follow":
				person := r.Form.Get("unfollow")
				// storage.WebDB.FollowUser(uName, person)
				client.RpcFollowUser(uName, person)
			case "unfollow":
				person := r.Form.Get("following")
				// storage.WebDB.UnFollowUser(uName, person)
				client.RpcUnFollowUser(uName, person)
			case "twit":
				redirectUrl += "/post"
			case "logout":
				redirectUrl = "/logout"
				cookie.ClearSession(w)
				redirectUrl = "/"

			}
			http.Redirect(w, r, redirectUrl, 302)
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
