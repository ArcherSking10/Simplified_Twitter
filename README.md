# Simplified Twitter

## Description
This a simplified version of twitter.  
It is a academic project built for 2018 fall distributed system course.
It is divided into three parts:
- [x] Build simple web application with database in memory
- [x] Split off backend into a seperate service
- [x] Bind the service with a distributed system

## Main Features
- Login & Register
- Edit & Create Post
- Follow & Unfollow Users
- View Posts from Followees and Myself 

## Instructions To Run
**1. Clone the project**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*git clone ...*   
**2. Download dependencies**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*git submodule update --init --recursive*  
**3. Start http server**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*go run src/main.go*  
**4. Start raft server**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;cd src/rpc/server/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;go run server.go -storageDir /tmp/node1 -nodeName node1 -rpcPort :9091 -raftPort :12000 -isLeader=true  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;go run server.go -storageDir /tmp/node2 -nodeName node2 -rpcPort :9092 -raftPort :13000 -isLeader=false  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;go run server.go -storageDir /tmp/node3 -nodeName node3 -rpcPort :9093 -raftPort :14000 -isLeader=false  
**5. Play with it**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; http://localhost:9090


## Project Structure
```bash
├── README.md
└── src
    ├── main.go
    ├── auth               // authentication module
    ├── handle             // handle join request from server
    ├── profile            // "presonal profile" module
    ├── twitter            // "twitter page" module
    ├── storage            // definition of various struct
    ├── rpc                // rpc module
    │   ├── client         // request from client
    │   ├── proto          // message proto
    │   └── server         // response from raft server
    ├── vendor             // dependencies
    └── template           // html and css templates
```

## Team
- Site Li (sl6890)
- Xinyu Ma (xm546)
- Kuang Sheng (ks4504, but not enrolled)
