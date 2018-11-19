package storage

import (
	"testing"
)

// var WebDB storage.DB

// type TwitPosts storage.TwitPosts
// type Twitlist storage.Twitlist
// type User storage.User

var emptyStr = ""
var name1 = "SiteLi"
var name2 = "LiSite"
var name3 = "KS"
var name4 = "SK"

var pswd1 = "100"
var pswd2 = "200"
var pswd3 = "3"
var pswd4 = "4"

var post1 = TwitPosts{Contents: "hehe", Date: 1, User: "SiteLi"}
var post2 = TwitPosts{Contents: "haha", Date: 2, User: "SiteLi"}
var postList = Twitlist{post1, post2}
var user = User{UserName: name1, passWord: pswd1, Posts: postList, Following: []string{name1, name2}}

// times := 10
// done := make(chan int, times)

func TestAddUser(t *testing.T) {
	// edge case test
	if flg := WebDB.AddUser(name1, pswd1, pswd1); flg != true {
		t.Errorf("Expected to pass registration ")
	}
	if flg := WebDB.AddUser(name2, pswd2, pswd2); flg != true {
		t.Errorf("Expected to pass registration ")
	}
	if flg := WebDB.AddUser(name1, pswd1, pswd1); flg != false {
		t.Errorf("Same user name cannot register twice")
	}
	if flg := WebDB.AddUser(emptyStr, pswd2, pswd1); flg != false {
		t.Errorf("User name field is empty")
	}
	if flg := WebDB.AddUser(name2, emptyStr, pswd1); flg != false {
		t.Errorf("Password field is empty")
	}
	if flg := WebDB.AddUser(name2, pswd1, pswd2); flg != false {
		t.Errorf("Password1 and password2 don't match")
	}
	if flg := WebDB.AddUser(name2, pswd2, pswd1); flg != false {
		t.Errorf("Password1 and password2 don't match")
	}
	if flg := WebDB.AddUser(name3, pswd3, pswd3); flg != true {
		t.Errorf("Expected to pass registration")
	}
	if flg := WebDB.AddUser(name4, pswd4, pswd4); flg != true {
		t.Errorf("Expected to pass registration")
	}
	// TODO : concurrency test
	// for i:= 0; i< times; i+= {
	// 	go AddUser(name1 + str(i), pswd1, pswd1)

	// }
}

func TestHasUser(t *testing.T) {
	// edge case test
	if flg := WebDB.HasUser(name1, pswd1); flg != true {
		t.Errorf("Expected user exists")
	}
	if flg := WebDB.HasUser(name1, pswd2); flg != false {
		t.Errorf("Expected user exists")
	}
	if flg := WebDB.HasUser(emptyStr, pswd1); flg != false {
		t.Errorf("Expected user exists")
	}
	// TODO : concurrency test
}

func TestUpdateUser(t *testing.T) {
	if flg := WebDB.UpdateUser(name1, user); flg != true {
		t.Errorf("Should update correctly")
	}
}

func TestFollowUser(t *testing.T) {
	if flg := WebDB.FollowUser(name1, name2); flg != false {
		t.Errorf("Expected Follwer already followed!")
	}
	if flg := WebDB.FollowUser(name1, name3); flg != true {
		t.Errorf("Expected Followed successfully")
	}
	if flg := WebDB.FollowUser(name2, name4); flg != true {
		t.Errorf("Expected Followed successfully!")
	}
}

func TestUnFollowUser(t *testing.T) {
	if flg := WebDB.UnFollowUser(name1, name2); flg != true {
		t.Errorf("Expected Unfollowed successfully!")
	}
	if flg := WebDB.UnFollowUser(name3, name4); flg != false {
		t.Errorf("Expected did not followed !")
	}
}

func TestGetTwitterPage(t *testing.T) {
	if pg := WebDB.GetTwitterPage(name1); len(pg.Posts) != 2 {
		t.Errorf("Expected Load correct pages !")
	}
}
