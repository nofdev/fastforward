package client

import (
	"testing"

	"github.com/nofdev/fastforward/provisioning/api/rpc/json/openstack"
	"github.com/nofdev/fastforward/provisioning"
)


func TestDo(t *testing.T) {
	url = "http://localhost:7001/v1"
	args = &openstack.Args{provisioning.ExtraVars{HostName: "localhost", RouterID: "51", State: "Master", Priority: "50"}}
	method = "OpenStack.LoadBalancer"
	err := Do(url, method, args)
	if err != nil {
		t.Error(err)
	}
}
