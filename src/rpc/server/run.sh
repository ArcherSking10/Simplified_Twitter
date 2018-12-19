rm -rf /tmp/node1
rm -rf /tmp/node2
rm -rf /tmp/node3
go run server.go -storageDir /tmp/node1 -nodeName node1 -rpcPort :9091 -raftPort :12000 -isLeader=true
go run server.go -storageDir /tmp/node2 -nodeName node2 -rpcPort :9092 -raftPort :13000 -isLeader=false
go run server.go -storageDir /tmp/node3 -nodeName node3 -rpcPort :9093 -raftPort :14000 -isLeader=false
