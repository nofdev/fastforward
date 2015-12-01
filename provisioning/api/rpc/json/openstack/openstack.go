package openstack

import (
	"net/http"

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
