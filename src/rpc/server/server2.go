package main

import (
	pb "Simplified_Twitter/src/rpc/proto"
	"Simplified_Twitter/src/storage"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	port = ":9093"
)

// Store is a simple key-value store, where all changes are made via Raft consensus.
type DB struct {
	RaftDir  string
	RaftBind string
	inmem    bool

	mu        sync.Mutex
	UsersInfo map[string]storage.User // The key-value store for the system.
	mp        map[string]string

	raft *raft.Raft // The consensus mechanism

	logger *log.Logger
}

type command struct {
	Op   string
	Name string
	Info storage.User
}

func (db *DB) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReply, error) {
	// db.mu.Lock()
	// defer db.mu.Unlock()
	var uName string = in.Uname
	// var tmp storage.User = db.UsersInfo[uName]
	tmp, _ := db.Get(uName)
	user := storage.ToPbType(tmp)
	// log.Printf("------> server user", user)
	return &pb.GetUserReply{Userinfo: user}, nil
}

func (db *DB) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.BoolReply, error) {
	// db.mu.Lock()
	// defer db.mu.Unlock()
	uName := in.Username
	pWord1 := in.Password1
	pWord2 := in.Password2
	if pWord1 != pWord2 {
		return &pb.BoolReply{T: false}, nil
	}
	if uName == "" || pWord1 == "" {
		return &pb.BoolReply{T: false}, nil
	}
	curUser := storage.User{uName, pWord1, storage.Twitlist{}, []string{uName}}
	// if _, ok := db.UsersInfo[uName]; ok {
	if _, ok := db.Get(uName); ok {
		return &pb.BoolReply{T: false}, nil
	}
	// Use uName as key put curUser inside
	db.Set(uName, curUser)
	// db.UsersInfo[uName] = curUser

	return &pb.BoolReply{T: true}, nil
}

func (db *DB) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.BoolReply, error) {
	// db.mu.Lock()
	// defer db.mu.Unlock()
	uName := in.Username
	usr := storage.PbTypeTo(in.Usr)
	if uName != usr.UserName {
		return &pb.BoolReply{T: false}, nil
	}
	// if _, ok := db.UsersInfo[uName]; ok != true {
	if _, ok := db.Get(uName); ok != true {
		return &pb.BoolReply{T: false}, nil
	}
	// db.UsersInfo[uName] = usr
	db.Set(uName, usr)
	return &pb.BoolReply{T: true}, nil
}

func (db *DB) HasUser(ctx context.Context, in *pb.HasUserRequest) (*pb.BoolReply, error) {
	// db.mu.Lock()
	// defer db.mu.Unlock()
	uName := in.Username
	pWord := in.Password
	if uName == "" || pWord == "" {
		return &pb.BoolReply{T: false}, nil
	}
	// Check Whether User in usersInfo
	user, exist := db.Get(uName)
	if exist && user.PassWord == pWord {
		return &pb.BoolReply{T: true}, nil
	}
	return &pb.BoolReply{T: false}, nil
}

func (db *DB) FollowUser(ctx context.Context, in *pb.FollowUserRequest) (*pb.BoolReply, error) {
	// db.mu.Lock()
	// defer db.mu.Unlock()
	uName := in.Username
	otherName := in.Othername
	if user, ok := db.Get(uName); ok {
		if storage.Contains(user.Following, otherName) {
			return &pb.BoolReply{T: false}, nil
		}
		user.Following = append(user.Following, otherName)
		// db.UsersInfo[uName] = user
		db.Set(uName, user)
		return &pb.BoolReply{T: true}, nil
	}
	return &pb.BoolReply{T: false}, nil
}

func (db *DB) UnFollowUser(ctx context.Context, in *pb.FollowUserRequest) (*pb.BoolReply, error) {
	// db.mu.Lock()
	// defer db.mu.Unlock()
	uName := in.Username
	otherName := in.Othername
	if user, ok := db.Get(uName); ok {
		if !storage.Contains(user.Following, otherName) {
			return &pb.BoolReply{T: false}, nil
		}
		user.Following = storage.Deletes(user.Following, otherName)
		db.Set(uName, user)
		return &pb.BoolReply{T: true}, nil
	}
	return &pb.BoolReply{T: false}, nil
}

// // Get Rid of the arrtribute of time
// // Just leave username + contents

func GetContents(arr storage.Twitlist) []string {
	var ret []string
	for _, twit := range arr {
		tmp := twit.User + ": " + twit.Contents
		ret = append(ret, tmp)
	}
	return ret
}

func (db *DB) IsLeader(ctx context.Context, in *pb.IsLeaderRequest) (*pb.BoolReply, error) {
	log.Println("------>", db.raft.State())
	var tmp bool = db.raft.State() == raft.Leader
	log.Println("-----> lisite", tmp)
	return &pb.BoolReply{T: tmp}, nil
}

func (db *DB) GetTwitterPage(ctx context.Context, in *pb.GetTwitterPageRequest) (*pb.GetTwitterPageReply, error) {
	// db.mu.Lock()
	// defer db.mu.Unlock()
	uName := in.Username
	user, _ := db.Get(uName)
	// log.Printf("-------> TwitterPage Userinfo ", user)
	UserName := user.UserName
	Following := user.Following
	// log.Printf("..............", Following)
	var UnFollowed []string
	var Posts storage.Twitlist
	// Get all Posts information
	for name, userInfo := range db.UsersInfo {
		if storage.Contains(Following, name) {
			for _, post := range userInfo.Posts {
				Posts = append(Posts, post)
			}
		} else {
			UnFollowed = append(UnFollowed, name)
		}
	}
	Posts = storage.Sort(Posts)
	newPosts := storage.GetContents(Posts)
	// Remove the user itself from following list (just not shown in screen but in memory)
	Following = storage.Deletes(Following, uName)
	var twit = &pb.TwitterPage{Username: UserName, UnFollowed: UnFollowed, Following: Following, Posts: newPosts}
	return &pb.GetTwitterPageReply{Twit: twit}, nil

}

type fsm DB

// Join joins a node, identified by nodeID and located at addr, to this store.
// The node must be ready to respond to Raft communications at that address.
func (s *DB) Join(ctx context.Context, in *pb.JoinRequest) (*pb.BoolReply, error) {
	nodeID := in.NodeID
	addr := in.RemoteAddr

	log.Println("------> Join.....")
	s.logger.Printf("received join request for remote node %s at %s", nodeID, addr)

	configFuture := s.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		s.logger.Printf("failed to get raft configuration: %v", err)
		return &pb.BoolReply{T: false}, err
	}
	for _, srv := range configFuture.Configuration().Servers {
		if srv.ID == raft.ServerID(nodeID) || srv.Address == raft.ServerAddress(addr) {
			if srv.Address == raft.ServerAddress(addr) && srv.ID == raft.ServerID(nodeID) {
				s.logger.Printf("node %s at %s already member of cluster, ignoring join request", nodeID, addr)
				return &pb.BoolReply{T: false}, nil
			}
			future := s.raft.RemoveServer(srv.ID, 0, 0)
			if err := future.Error(); err != nil {
				return &pb.BoolReply{T: false}, fmt.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
			}
		}
	}
	f := s.raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if f.Error() != nil {
		fmt.Println(f.Error())
		return &pb.BoolReply{T: false}, f.Error()
	}
	s.logger.Printf("node %s at %s joined successfully", nodeID, addr)
	return &pb.BoolReply{T: true}, nil
}

func (s *DB) Open(enableSingle bool, localID string) error {
	// Setup Raft configuration.
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(localID)

	// Setup Raft communication.
	addr, err := net.ResolveTCPAddr("tcp", s.RaftBind)
	if err != nil {
		return err
	}
	transport, err := raft.NewTCPTransport(s.RaftBind, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}
	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshots, err := raft.NewFileSnapshotStore(s.RaftDir, 2, os.Stderr)
	if err != nil {
		return fmt.Errorf("file snapshot store: %s", err)
	}
	// Create the log store and stable store.
	var logStore raft.LogStore
	var stableStore raft.StableStore
	if s.inmem {
		logStore = raft.NewInmemStore()
		stableStore = raft.NewInmemStore()
	} else {
		boltDB, err := raftboltdb.NewBoltStore(filepath.Join(s.RaftDir, "raft.db"))
		if err != nil {
			return fmt.Errorf("new bolt store: %s", err)
		}
		logStore = boltDB
		stableStore = boltDB
	}
	// Instantiate the Raft systems.
	ra, err := raft.NewRaft(config, (*fsm)(s), logStore, stableStore, snapshots, transport)
	if err != nil {
		return fmt.Errorf("new raft: %s", err)
	}
	s.raft = ra
	if enableSingle {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		ra.BootstrapCluster(configuration)
	}
	return nil
}

type fsmSnapshot struct {
	UsersInfo map[string]storage.User
	mp        map[string]string
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	// Clone the map.
	o1 := make(map[string]storage.User)
	o2 := make(map[string]string)
	for k, v := range f.UsersInfo {
		o1[k] = v
	}
	for k, v := range f.mp {
		o2[k] = v
	}
	return &fsmSnapshot{UsersInfo: o1, mp: o2}, nil
}

// Restore stores the key-value store to a previous state.
func (f *fsm) Restore(rc io.ReadCloser) error {
	return nil
}

// Get returns the value for the given key.
func (s *DB) Get(name string) (storage.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	info, ok := s.UsersInfo[name]
	return info, ok
}

// Set sets the value for the given key.
func (s *DB) Set(name string, info storage.User) error {
	if s.raft.State() != raft.Leader {
		return fmt.Errorf("not leader")
	}
	c := &command{
		Op:   "set",
		Name: name,
		Info: info,
	}
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	f := s.raft.Apply(b, 2*time.Second)
	return f.Error()
}

// Apply applies a Raft log entry to the key-value store.
func (f *fsm) Apply(l *raft.Log) interface{} {
	var c command
	if err := json.Unmarshal(l.Data, &c); err != nil {
		panic(fmt.Sprintf("failed to unmarshal command: %s", err.Error()))
	}
	switch c.Op {
	case "set":
		return f.applySet(c.Name, c.Info)
	default:
		panic(fmt.Sprintf("unrecognized command op: %s", c.Op))
	}
}

func (f *fsm) applySet(name string, info storage.User) interface{} {
	f.mu.Lock()
	defer f.mu.Unlock()
	fmt.Println(name)
	fmt.Println(info)
	f.UsersInfo[name] = info
	return nil
}

func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	return nil
}

func join(nodeID, raftAddr string) error {
	b, err := json.Marshal(map[string]string{"addr": raftAddr, "id": nodeID})
	if err != nil {
		return err
	}
	log.Println("-------> Enter join")
	// http.Post(request_url, "application/x-www-form-urlencoded", body)
	resp, err := http.Post(fmt.Sprintf("http://:9090/join"), "application-type/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	log.Println("======")
	defer resp.Body.Close()

	return nil
}

func (f *fsmSnapshot) Release() {}

func main() {
	// go func() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Println("starting server 3")
	WebDB3 := &DB{}
	WebDB3.inmem = false
	WebDB3.RaftDir = "./tmp3"
	WebDB3.RaftBind = ":14000"
	WebDB3.logger = log.New(os.Stderr, "[store3] ", log.LstdFlags)
	WebDB3.UsersInfo = make(map[string]storage.User)
	WebDB3.mp = make(map[string]string)
	if err := WebDB3.Open(":13000" == "", "node3"); err != nil {
		log.Fatalf("failed to open store: %s", err.Error())
	}
	time.Sleep(5 * time.Second)
	// WebDB.Join("node2", ":13000")
	join("node3", ":14000")
	//}()
	pb.RegisterWebServer(s, WebDB3)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// time.Sleep(5 * time.Second)
}
