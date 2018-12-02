package client

import (
	"context"
	"google.golang.org/grpc"
	pb "Simplified_Twitter/src/rpc/proto"
	"log"
	"Simplified_Twitter/src/storage"
	"time"
)

const (
	address = "localhost:9091"
	// defaultName = "world"
)

func RpcGetUser(uName string) storage.User {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWebClient(conn)

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
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWebClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	pbUser := storage.ToPbType(usr)
	r, err := c.UpdateUser(ctx, &pb.UpdateUserRequest{Username: uName, Usr: pbUser})
	return r.T
}

func RpcAddUser(uName string, pWord1 string, pWord2 string) bool {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWebClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, err := c.AddUser(ctx, &pb.AddUserRequest{Username: uName, Password1: pWord1, Password2: pWord2})
	// log.Printf("-------> rpcAdduser", r.T)
	return r.T

}

func RpcHasUser(uName string, pWord string) bool {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWebClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, err := c.HasUser(ctx, &pb.HasUserRequest{Username: uName, Password: pWord})
	// log.Printf("-------> rpcHasuser", r.T)
	return r.T
}

func RpcGetTwitterPage(uName string) storage.TwitterPage {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWebClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, err := c.GetTwitterPage(ctx, &pb.GetTwitterPageRequest{Username: uName})
	log.Println("--------> TwitterPage", r)
	var twit = storage.TwitterPage{UserName: r.Twit.Username, UnFollowed: r.Twit.UnFollowed, Following: r.Twit.Following, Posts: r.Twit.Posts}
	return twit
}

func RpcFollowUser(uName string, person string) bool {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWebClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, err := c.FollowUser(ctx, &pb.FollowUserRequest{Username: uName, Othername: person})
	// log.Printf("......follow user", r.T)
	return r.T
}

func RpcUnFollowUser(uName string, person string) bool {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWebClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	r, err := c.UnFollowUser(ctx, &pb.FollowUserRequest{Username: uName, Othername: person})
	// log.Printf("......unfollow user", r.T)
	return r.T
}
