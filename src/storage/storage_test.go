package storage

import (
	"testing"
)

var emptyStr = ""
var name1 = "SiteLi"
var name2 = "LiSite"
var pswd1 = "100"
var pswd2 = "200"


var post1 = TwitPosts {Contents: "hehe",  Date  : 1,  User  : "SiteLi" }
var post2 = TwitPosts {Contents: "haha",  Date  : 2,  User  : "SiteLi" }
var postList = Twitlist {post1, post2}
var user = User { UserName : name1, passWord : pswd1, Posts : postList, Following : []string{name1, name2} }

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
	// TODO : concurrency test

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
