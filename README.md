# Simplified Twitter

## Description
This a simplified version of twitter.  
It is a academic project built for 2018 fall distributed system course.
It is divided into three parts:
- [x] Build simple web application with database in memory
- [x] Split off backend into a seperate service
- [ ] Bind the service with a distributed system

## Main Features
- Login & Register
- Edit & Create Post
- Follow & Unfollow Users
- View Posts from Followees and Myself 

## Instructions To Run
**1. Clone the project into "/your/path"**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*git clone ...*   
**2. Go into the src directory and run it**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*cd /your/path/src*  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*go run src/server/server.go*  --> set up rpc server  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*go run src/main.go*           --> set up web server  

## Project Structure
```bash
├── README.md
└── src
    ├── main.go
    ├── auth       // "login & logout" module
    │   └── cookie
    ├── profile    // "presonal profile" module
    ├── twitter    // "twitter page" module
    ├── storage    // definition of various struct
    ├── rpc        // gRPC module
    │   ├── client
    │   ├── server
    │   └── proto
    ├── vendor     // dependencies
    └── template   // html and css templates
        ├── *.html
        └── static
            ├── css
            └── js
```

## Team
- Site Li (sl6890)
- Kuang Sheng (ks4504, but not enrolled)
