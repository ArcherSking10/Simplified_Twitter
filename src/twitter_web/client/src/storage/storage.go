package storage

type User struct {
	UserName  string
	passWord  string
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

// Get Rid of the arrtribute of time
// Just leave username + contents
