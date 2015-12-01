package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/nofdev/fastforward/provisioning"
)

// OpenStack API.
type OpenStack struct {}

// Args takes extra-vars of Playback
type Args struct {
	provisioning.ExtraVars
}

// Result contains the API call results.
type Result interface{}

// ConfigureStorageNetwork takes playback-nic to set up the storage network.
// Args: {"PlaybackNic.Purge": bool, "PlaybackNic.Public": bool, "PlaybackNic.Private": bool, "PlaybackNic.Host": string, "PlaybackNic.User": string, "PlaybackNic.Address": string, "PlaybackNic.NIC": string, "PlaybackNic.Netmask": string, "PlaybackNic.Gateway": string}
func (o *OpenStack) ConfigureStorageNetwork(r *http.Request, args *Args, result *Result) error {
	*result = args.ExtraVars.ConfigureStorageNetwork()
	return nil
}

func main() {
	s := rpc.NewServer()
	log.Printf("OpenStack API started")
	s.RegisterCodec(json.NewCodec(), "application/json")
	openstack := new(OpenStack)
	s.RegisterService(openstack, "")
	log.Printf("Register OpenStack service")
	r := mux.NewRouter()
	r.Handle("/v1", s)
	log.Printf("Handle API version 1")
	log.Printf("Listen on port 7001")
	http.ListenAndServe(":7001", r)
}