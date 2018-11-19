# Simplified Twitter

## Description
This a simplified version of twitter.  
It is a academic project built for 2018 fall distributed system course.
It is divided into three parts:
- [x] Build simple web application with database in memory
- [ ] Split off backend into a seperate service
- [ ] Bind the service with a distributed system

## Main Features
- Login & Register
- Edit & Create Post
- Follow & Unfollow Users
- View Posts from Followees and Myself 

## Instructions To Run
**1. Install thrid-party packages**   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*go get github.com/gorilla/securecookie*  
**2. Clone the project into "/your/path"**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*git clone ...*   
**3. Go into the src directory and run it**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*cd /your/path/src*  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*go run main.go*

## Project Structure
```bash
├── README.md
└── src
    ├── main.go    // go run main.go
    ├── auth       // "login & logout" module
    │   ├── cookie
    │   │   └── cookie.go
    │   └── *.go
    ├── profile    // "presonal profile" module
    │   └── *.go
    ├── twitter    // "twitter page" module
    │   └── *.go
    ├── storage    // database , its functions and test
    │   └── *.go
    └── template   // html and css templates
        ├── *.html
        └── static
            ├── css
            └── js
```

## Team
- Site Li
- Kuang Sheng
- Xinyu Ma
