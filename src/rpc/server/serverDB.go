package serverDB

import (
	// pb "Simplified_Twitter/src/rpc/proto"
	"Simplified_Twitter/src/storage"
	// "bytes"
	// "context"
	// "encoding/json"
	// "flag"
	// "fmt"
	"github.com/hashicorp/raft"
	// "github.com/hashicorp/raft-boltdb"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
	// "io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

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
