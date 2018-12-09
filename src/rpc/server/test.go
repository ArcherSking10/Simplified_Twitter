package main

import (
	"flag"
	"fmt"
)

var storageDir string
var rpcPort string
var raftPort string
var nodeName string

func init() {
	flag.StringVar(&storageDir, "storageDir", "/tmp/dir1", "Set the storage directory")
	flag.StringVar(&rpcPort, "rpcPort", "9090", "Set Rpc bind address")
	flag.StringVar(&raftPort, "raftPort", "9091", "Set Raft bind address")
	flag.StringVar(&nodeName, "nodeName", "node0", "Set the name of server")
	flag.Usage = func() {
		fmt.Println("Usage: go run server.go [options] <data-path>")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	fmt.Println(storageDir)
	fmt.Println(rpcPort)
	fmt.Println(raftPort)
	fmt.Println(nodeName)
	// fmt.Println(flag.Arg(0))
}
