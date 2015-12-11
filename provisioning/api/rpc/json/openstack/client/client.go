package client

import (
	"net/http"
	"log"
	"bytes"
	
	"github.com/gorilla/rpc/v2/json"
	"github.com/nofdev/fastforward/provisioning/api/rpc/json/openstack"
	"github.com/nofdev/fastforward/provisioning"
)

var url string
var args *openstack.Args
var method string
var result openstack.Result

func checkErr(err error) error{
	if err !=nil {
		log.Printf("%s", err)
		return err
	}
	return nil
}

// Do holds a JSON-RPC request.
func Do(url, method string, args *openstack.Args) error {
	message, err := json.EncodeClientRequest(method, args)
	checkErr(err)
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	checkErr(err)
	req.Header.Set("Content-Type", "application/json")
	
	client := new(http.Client)
	resp, err := client.Do(req); if err != nil {
		log.Printf("Error in sending request to %s. %s", url, err)
		return err
	}
	defer resp.Body.Close()
	
	err = json.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Printf("Couldn't decode response. %s", err)
		return err
	}
	log.Printf("url: %s, method: %s, args: %s", url, method, args)
	
	return nil
}
