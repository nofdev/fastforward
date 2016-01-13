// Package main is the JSON-RPC 2.0 API server for playback.
package main

import (
	"log"
	"net/http"

	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/gorilla/rpc/v2"
	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/gorilla/rpc/v2/json"
	"github.com/nofdev/fastforward/provisioning/api/rpc/json/openstack"
)

func main() {
	s := rpc.NewServer()
	log.Printf("Playback API started")
	s.RegisterCodec(json.NewCodec(), "application/json")
	openstack := new(openstack.OpenStack)
	s.RegisterService(openstack, "")
	log.Printf("Register OpenStack service")
	r := mux.NewRouter()
	r.Handle("/v1", s)
	log.Printf("Handle API version 1")
	log.Printf("Listen on port 7001")
	http.ListenAndServe(":7001", r)
}
