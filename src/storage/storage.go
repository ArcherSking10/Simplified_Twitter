package storage

import (
	"fmt"
	"reflect"
	"sort"
)

var WebDB DB


type User struct {
	UserName string
	passWord string
	Posts     Twitlist
	Following []string
}

// Define this type for sort
type Twitlist []TwitPosts

// Using time to define post order
// Username to who do the posts

type TwitPosts struct {
	Contents string
	Date     int64
	User     string
}

type TwitterPage struct {
	UserName   string
	UnFollowed []string
	Following  []string
	Posts      []string
	// Posts Twitlist
}

type DB struct {
	UsersInfo map[string]User
}

// Sort Function needed these three Function
func (I Twitlist) Len() int {
	return len(I)
}
func (I Twitlist) Less(i, j int) bool {
	return I[i].Date < I[j].Date
}
func (I Twitlist) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

func (db *DB) GetUser(uName string) User {
	return db.UsersInfo[uName]
}

func (db *DB) AddUser(uName string, pWord1 string, pWord2 string) bool {
	if pWord1 != pWord2 {
		return false
	}
	if uName == "" || pWord1 == "" {
		return false
	}
	curUser := User{uName, pWord1, Twitlist{}, []string{uName}}
	if _, ok := db.UsersInfo[uName]; ok {
		return false
	}
	// Use uName as key put curUser inside
	db.UsersInfo[uName] = curUser
	return true
}

func (db *DB) UpdateUser(uName string, usr User) bool {
	if _, ok := db.UsersInfo[uName]; ok != true {
		return false
	}
	db.UsersInfo[uName] = usr
	return true
}

func (db *DB) HasUser(uName string, pWord string) bool {
	if uName == "" || pWord == "" {
		return false
	}
	// Check Whether User in usersInfo
	user, exist := db.UsersInfo[uName]
	if exist && user.passWord == pWord {
		return true
	}
	return false
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Deletes(a []string, x string) []string {
	var ret []string
	for _, n := range a {
		if x != n {
			ret = append(ret, n)
		}
	}
	return ret
}

func (db *DB) FollowUser(uName string, otherName string) bool {
	user, _ := db.UsersInfo[uName]
	if Contains(user.Following, otherName) {
		return false
	}
	user.Following = append(user.Following, otherName)
	db.UpdateUser(uName, user)
	return true
}

func (db *DB) UnFollowUser(uName string, otherName string) bool {
	user, _ := db.UsersInfo[uName]
	if !Contains(user.Following, otherName) {
		return false
	}
	user.Following = Deletes(user.Following, otherName)
	db.UpdateUser(uName, user)
	return true
}

// Get Rid of the arrtribute of time
// Just leave username + contents

func GetContents(arr Twitlist) []string {
	var ret []string
	for _, twit := range arr {
		tmp := twit.User + ": " + twit.Contents
		ret = append(ret, tmp)
	}
	return ret
}

func (db *DB) GetTwitterPage(uName string) TwitterPage {
	user, _ := db.UsersInfo[uName]
	UserName := user.UserName
	Following := user.Following
	var UnFollowed []string
	var Posts Twitlist
	// Get all Posts information
	for name, userInfo := range db.UsersInfo {
		if Contains(Following, name) {
			for _, post := range userInfo.Posts {
				Posts = append(Posts, post)
			}
		} else {
			UnFollowed = append(UnFollowed, name)
		}
	}
	fmt.Println(Posts)
	sort.Sort(Posts)
	fmt.Println("Type ------>", reflect.TypeOf(Posts))
	newPosts := GetContents(Posts)
	// Remove the user itself from following list (just not shown in screen but in memory)
	Following = Deletes(Following, uName)
	fmt.Println("-------->", Posts)
	pg := TwitterPage{UserName: UserName, Following: Following, UnFollowed: UnFollowed, Posts: newPosts}
	return pg
}

func init() {
	m := make(map[string]User)
	WebDB = DB{UsersInfo: m}
}
