package storage

import (
	"fmt"
)

var WebDB DB

type User struct {
	UserName  string
	passWord  string
	Posts     []string
	Following []string
}

type TwitterPage struct {
	UserName     string
	UnFollowed []string
	Following  []string
	Posts      []string
}

type DB struct {
	UsersInfo map[string]User
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
	curUser := User{uName, pWord1, []string{}, []string{uName}}
	// Use uName as key put curUser inside
	db.UsersInfo[uName] = curUser
	return true
}

func (db *DB) UpdateUser(uName string, usr User) bool {
	db.UsersInfo[uName] = usr
	return true
}

func (db *DB) HasUser(uName string, pWord string) bool {
	if uName == "" || pWord == "" {
		return false
	}
	// Check Whether User in usersInfo
	user, exist := db.UsersInfo[uName]
	if exist && user.passWord == pWord{
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
            ret = append(ret, x)
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
	if ! Contains(user.Following, otherName) {
		return false
	}
	user.Following = Deletes(user.Following, otherName)
	db.UpdateUser(uName, user)
	return true
}

func (db *DB) GetTwitterPage(uName string) TwitterPage {
	user, _ := db.UsersInfo[uName]
	UserName := user.UserName
	Following := user.Following
	var UnFollowed []string
	var Posts []string
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
	pg := TwitterPage{UserName : UserName, Following : Following, UnFollowed : UnFollowed, Posts : Posts}
	return pg
}

func init() {
	m := make(map[string]User)
	WebDB = DB{UsersInfo: m}
}
