package storage


var WebDB DB

type User struct {  
    userName   string
    passWord   string
    posts     []string
    following []string
}

type DB struct {
	usersInfo [] User
}

func (db *DB) AddUser(uName string, pWord1 string, pWord2 string) bool {
	if pWord1 != pWord2 {
		return false
	}
	if uName == "" || pWord1 == "" {
		return false
	}
	curUser := User {uName, pWord1, []string{}, []string{}}
	db.usersInfo = append(db.usersInfo, curUser)
	return true
}

func (db *DB) HasUser(uName string, pWord string) bool {
	if (uName == "" || pWord == "") {
		return false
	}
    for _, user := range db.usersInfo {
        if (uName == user.userName && pWord == user.passWord) {
        	return true
        }
    }
	return false
}

func init() {
	WebDB = DB {}
}