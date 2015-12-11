/*
Package client provides a playback api client for ff command line.
Example:
	import (
	"log"
	"github.com/nofdev/fastforward/provisioning"
	"github.com/nofdev/fastforward/provisioning/api/rpc/json/openstack"
	"github.com/nofdev/fastforward/provisioning/api/rpc/json/openstack/client"
	)
	
	// The playback API address.
	url = "http://localhost:7001/v1"
	// Define arguments for load balancer.
	args = &openstack.Args{provisioning.ExtraVars{HostName: "localhost", RouterID: "51", State: "Master", Priority: "50"}}
	// Deploy load balancer for OpenStack HA.
	err := client.LoadBalancer(args)
	if err != nil {
		log.Fatal(err)
	}
*/
package client