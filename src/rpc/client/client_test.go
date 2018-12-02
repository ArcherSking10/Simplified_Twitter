package client

import (
	"testing"
	// "fmt"
	"Simplified_Twitter/src/storage"
	"strconv"
	"sync"
)

// var WebDB storage.DB

// type TwitPosts storage.TwitPosts
// type Twitlist storage.Twitlist
// type User storage.User

var concurrencyFactor = 100

var emptyStr = ""
var name1 = "SiteLi"
var name2 = "LiSite"
var name3 = "KS"
var name4 = "SK"

var pswd1 = "100"
var pswd2 = "200"
var pswd3 = "3"
var pswd4 = "4"

var post1 = storage.TwitPosts{Contents: "hehe", Date: 1, User: "SiteLi"}
var post2 = storage.TwitPosts{Contents: "haha", Date: 2, User: "SiteLi"}
var postList = storage.Twitlist{post1, post2}
var user = storage.User{UserName: name1, PassWord: pswd1, Posts: postList, Following: []string{name1, name2}}




/*
concurrency test:
go run() {}

t.run():
*/

func TestAddUser(t *testing.T) {
	t.Run("EdgeCaseTest", func(t *testing.T) {
		if flg := RpcAddUser(name1, pswd1, pswd1); flg != true {
			t.Errorf("Expected to pass registration ")
		}
		if flg := RpcAddUser(name2, pswd2, pswd2); flg != true {
			t.Errorf("Expected to pass registration ")
		}
		if flg := RpcAddUser(name1, pswd1, pswd1); flg != false {
			t.Errorf("Same user name cannot register twice")
		}
		if flg := RpcAddUser(emptyStr, pswd2, pswd1); flg != false {
			t.Errorf("User name field is empty")
		}
		if flg := RpcAddUser(name2, emptyStr, pswd1); flg != false {
			t.Errorf("Password field is empty")
		}
		if flg := RpcAddUser(name2, pswd1, pswd2); flg != false {
			t.Errorf("Password1 and password2 don't match")
		}
		if flg := RpcAddUser(name2, pswd2, pswd1); flg != false {
			t.Errorf("Password1 and password2 don't match")
		}
		if flg := RpcAddUser(name3, pswd3, pswd3); flg != true {
			t.Errorf("Expected to pass registration")
		}
		if flg := RpcAddUser(name4, pswd4, pswd4); flg != true {
			t.Errorf("Expected to pass registration")
		}
	})
	t.Run("ConcurrencyTest", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(concurrencyFactor)
		for i := 0; i < concurrencyFactor; i++ {
			var testName = name1 + strconv.Itoa(i)
			var testPassword = pswd1 + strconv.Itoa(i)
			go func(name string, password string) {
				defer wg.Done()
				RpcAddUser(name, password, password)
			}(testName, testPassword)
		}
		wg.Wait()
	})
}

func TestHasUser(t *testing.T) {
	t.Run("EdgeCaseTest", func(t *testing.T) {
		if flg := RpcHasUser(name1, pswd1); flg != true {
			t.Errorf("Expected user exists")
		}
		if flg := RpcHasUser(name1, pswd2); flg != false {
			t.Errorf("Expected user exists")
		}
		if flg := RpcHasUser(emptyStr, pswd1); flg != false {
			t.Errorf("Expected user exists")
		}
	})
	t.Run("ConcurrencyTest", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(concurrencyFactor)
		for i := 0; i < concurrencyFactor; i++ {
			var testName = name1 + strconv.Itoa(i)
			var expect = true
			var testPassword = pswd1 + strconv.Itoa(i)
			if i > concurrencyFactor / 2 {
				expect = false
				testPassword = testPassword + strconv.Itoa(i)
			}
			go func(name string, password string, expect bool) {
				defer wg.Done()
				if flg := RpcHasUser(name, password); flg != expect {
					t.Errorf("User exists /not exists")
				}
			}(testName, testPassword, expect)
		}
		wg.Wait()
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("NormalCaseTest", func(t *testing.T) {
		if flg := RpcUpdateUser(name1, user); flg != true {
			t.Errorf("Should update correctly")
		}
	})
	t.Run("ConcurrencyTest", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(concurrencyFactor)
		for i := 0; i < concurrencyFactor; i++ {
			go func(name string, user storage.User) {
				defer wg.Done()
				if flg := RpcUpdateUser(name1, user); flg != true {
					t.Errorf("Should update correctly")
				}
			}(name1, user)
		}
		wg.Wait()
	})
}

func TestFollowUser(t *testing.T) {
	t.Run("NormalCaseTest", func(t *testing.T) {
		if flg := RpcFollowUser(name1, name2); flg != false {
			t.Errorf("Expected Follwer already followed!")
		}
		if flg := RpcFollowUser(name1, name3); flg != true {
			t.Errorf("Expected Followed successfully")
		}
		if flg := RpcFollowUser(name2, name4); flg != true {
			t.Errorf("Expected Followed successfully!")
		}
	})
	t.Run("ConcurrencyTest", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(concurrencyFactor)
		for i := 0; i < concurrencyFactor; i++ {
			go func(name1 string, name2 string) {
				defer wg.Done()
				if flg := RpcFollowUser(name1, name2); flg != false {
					t.Errorf("Follwer already followed!")
				}
			}(name1, name2)
		}
		wg.Wait()
	})
}

func TestUnFollowUser(t *testing.T) {
	t.Run("NormalCaseTest", func(t *testing.T) {
		if flg := RpcUnFollowUser(name1, name2); flg != true {
			t.Errorf("Expected return successfully!")
		}
		if flg := RpcUnFollowUser(name3, name4); flg != false {
			t.Errorf("Expected did not followed !")
		}
	})
	t.Run("ConcurrencyTest", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(concurrencyFactor)
		for i := 0; i < concurrencyFactor; i++ {
			go func(name1 string, name2 string) {
				defer wg.Done()
				if flg := RpcUnFollowUser(name1, name2); flg != false {
					t.Errorf("Follwer already followed!")
				}
			}(name1, name2)
		}
		wg.Wait()
	})
}

func TestGetTwitterPage(t *testing.T) {
	t.Run("NormalCaseTest", func(t *testing.T) {
		if pg := RpcGetTwitterPage(name1); len(pg.Posts) != 2 {
			t.Errorf("Expected Load correct pages !")
		}
	})
	t.Run("ConcurrencyTest", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(concurrencyFactor)
		for i := 0; i < concurrencyFactor; i++ {
			go func(name1 string) {
				defer wg.Done()
				if pg := RpcGetTwitterPage(name1); len(pg.Posts) != 2 {
					t.Errorf("Expected Load correct pages !")
				}
			}(name1)
		}
		wg.Wait()
	})
}
