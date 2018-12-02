package storage

import (
	pb "Simplified_Twitter/src/rpc/proto"
	"sort"
)

type User struct {
	UserName  string
	PassWord  string
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

// Sort Posts according to Date
func Sort(posts Twitlist) Twitlist {
	var ret Twitlist
	sort.Sort(posts)
	for _, val := range posts {
		ret = append(ret, val)
	}
	return ret
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
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

func Deletes(a []string, x string) []string {
	var ret []string
	for _, n := range a {
		if x != n {
			ret = append(ret, n)
		}
	}
	return ret
}

func PbTypeTo(user *pb.User) User {
	var posts Twitlist
	for _, post := range user.Posts {
		var tmpPost = TwitPosts{Contents: post.Contents, Date: post.Date, User: post.User}
		posts = append(posts, tmpPost)
	}
	var tmp = User{UserName: user.UserName, PassWord: user.PassWord, Posts: posts, Following: user.Following}
	return tmp
}

func ToPbType(tmp User) *pb.User {
	var posts []*pb.TwitPosts
	for _, post := range tmp.Posts {
		var tmpPost = &pb.TwitPosts{Contents: post.Contents, Date: post.Date, User: post.User}
		posts = append(posts, tmpPost)
	}
	var user = &pb.User{UserName: tmp.UserName, PassWord: tmp.PassWord, Posts: posts, Following: tmp.Following}
	return user
}
