package storage

var WebDB DB

type User struct {
	UserName  string
	passWord  string
	Posts     []string
	following []string
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
	curUser := User{uName, pWord1, []string{}, []string{}}
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
	_, exist := db.UsersInfo[uName]
	return exist
}

func init() {
	m := make(map[string]User)
	WebDB = DB{UsersInfo: m}
}
