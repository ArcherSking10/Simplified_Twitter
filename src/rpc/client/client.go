package client

import (
	pb "Simplified_Twitter/src/rpc/proto"
	"Simplified_Twitter/src/storage"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"
)

var curLeader int = -1
var addresslist = []string{"localhost:9091", "localhost:9092", "localhost:9093"}

func RpcEstablish(addresslist []string) (*grpc.ClientConn, pb.WebClient) {
	// a = append(a, b...)
	if curLeader != -1 {
		addresslist = append(addresslist[curLeader+1:], addresslist[:curLeader]...)
	}
	for idx, address := range addresslist {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c := pb.NewWebClient(conn)
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		log.Println("---------> 6666666", c)
		log.Println("-----> sssss")
		port := strings.Split(address, ":")[1]
		fmt.Println(port)
		var add = &pb.IsLeaderRequest{Address: port}
		t, err := c.IsLeader(ctx, add)
		log.Println("--------> T", t)
		log.Println("--------> err", err)
		if err == nil {
			if t.T {
				return conn, c
			}
		} else {
			curLeader = idx
		}
		conn.Close()
	}
	log.Println("----> pointer")
	return nil, nil
}

// func RpcEstablish(addresslist []string) (*grpc.ClientConn, pb.WebClient) {
// 	for _, address := range addresslist {
// 		conn, err := grpc.Dial(address, grpc.WithInsecure())
// 		if err != nil {
// 			log.Fatalf("did not connect: %v", err)
// 		}
// 		c := pb.NewWebClient(conn)
// 		ctx, _ := context.WithTimeout(context.Background(), time.Second)
// 		log.Println("---------> 6666666", c)
// 		log.Println("-----> sssss")
// 		port := strings.Split(address, ":")[1]
// 		fmt.Println(port)
// 		var add = &pb.IsLeaderRequest{Address: port}
// 		t, err := c.IsLeader(ctx, add)
// 		log.Println("--------> T", t)
// 		log.Println("--------> err", err)
// 		if err == nil && t.T {
// 			curLeader =
// 			return conn, c
// 		}
// 		conn.Close()
// 	}
// 	log.Println("----> pointer")
// 	return nil, nil
// }

func RpcGetUser(uName string) storage.User {
	conn, c := RpcEstablish(addresslist)
	defer conn.Close()
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, err := c.GetUser(ctx, &pb.GetUserRequest{Uname: uName})
	if err != nil {
		log.Printf("failed to call: %v", err)
		// return nil
	}
	user := r.Userinfo
	tmp := storage.PbTypeTo(user)
	return tmp
}

func RpcUpdateUser(uName string, usr storage.User) bool {
	conn, c := RpcEstablish(addresslist)
	defer conn.Close()

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	pbUser := storage.ToPbType(usr)
	r, _ := c.UpdateUser(ctx, &pb.UpdateUserRequest{Username: uName, Usr: pbUser})
	return r.T
}

func RpcAddUser(uName string, pWord1 string, pWord2 string) bool {
	conn, c := RpcEstablish(addresslist)
	defer conn.Close()

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, _ := c.AddUser(ctx, &pb.AddUserRequest{Username: uName, Password1: pWord1, Password2: pWord2})
	// log.Printf("-------> rpcAdduser", r.T)
	return r.T

}

func RpcHasUser(uName string, pWord string) bool {
	conn, c := RpcEstablish(addresslist)
	defer conn.Close()

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, _ := c.HasUser(ctx, &pb.HasUserRequest{Username: uName, Password: pWord})
	// log.Printf("-------> rpcHasuser", r.T)
	return r.T
}

func RpcGetTwitterPage(uName string) storage.TwitterPage {
	conn, c := RpcEstablish(addresslist)
	defer conn.Close()

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, _ := c.GetTwitterPage(ctx, &pb.GetTwitterPageRequest{Username: uName})
	log.Println("--------> TwitterPage", r)
	var twit = storage.TwitterPage{UserName: r.Twit.Username, UnFollowed: r.Twit.UnFollowed, Following: r.Twit.Following, Posts: r.Twit.Posts}
	return twit
}

func RpcFollowUser(uName string, person string) bool {
	conn, c := RpcEstablish(addresslist)
	defer conn.Close()

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, _ := c.FollowUser(ctx, &pb.FollowUserRequest{Username: uName, Othername: person})
	// log.Printf("......follow user", r.T)
	return r.T
}

func RpcUnFollowUser(uName string, person string) bool {
	conn, c := RpcEstablish(addresslist)
	defer conn.Close()

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, _ := c.UnFollowUser(ctx, &pb.FollowUserRequest{Username: uName, Othername: person})
	// log.Printf("......unfollow user", r.T)
	return r.T
}

func RpcJoin(nodeID string, remoteAddr string) {
	// addresslist = append(addresslist, "local:host"+remoteAddr)
	conn, c := RpcEstablish(addresslist)
	defer conn.Close()
	log.Println("----->5")
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	var join = &pb.JoinRequest{NodeID: nodeID, RemoteAddr: remoteAddr}
	c.Join(ctx, join)
}
