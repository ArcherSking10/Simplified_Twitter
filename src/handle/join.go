package handle

import (
	// "Simplified_Twitter/src/auth/cookie"
	"Simplified_Twitter/src/rpc/client"
	// "fmt"
	// "html/template"
	"encoding/json"
	"log"
	"net/http"
)

func Join(w http.ResponseWriter, r *http.Request) {
	log.Println("-------> Enter")
	m := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("----> 1")
	if len(m) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("----> 2")
	remoteAddr, ok := m["addr"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("----> 3")
	nodeID, ok := m["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("----> 4")
	client.RpcJoin(nodeID, remoteAddr)
}
