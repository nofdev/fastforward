// Package openstack provides an API of JSON-RPC 2.0 for Playback.
package openstack

import (
	"net/http"

	"github.com/nofdev/fastforward/provisioning"
)

// OpenStack API.
type OpenStack struct{}

// Args takes extra-vars of Playback
type Args struct {
	provisioning.ExtraVars
}

// Result contains the API call results.
type Result interface{}

// ConfigureStorageNetwork takes playback-nic to set up the storage network.
//  Args: {"PlaybackNic.Purge": bool, "PlaybackNic.Public": bool, "PlaybackNic.Private": bool, "PlaybackNic.Host": string, "PlaybackNic.User": string, "PlaybackNic.Address": string, "PlaybackNic.NIC": string, "PlaybackNic.Netmask": string, "PlaybackNic.Gateway": string}
func (o *OpenStack) ConfigureStorageNetwork(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.ConfigureStorageNetwork()
	return nil
}

// TODO: Args comment.
// TODO: Return value of errors.

// LoadBalancer deploy a HAProxy and Keepalived for OpenStack HA.
//  Args: {"HostName": string, "RouterID": string, "State": string, "Priority": int}
func (o *OpenStack) LoadBalancer(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.LoadBalancer()
	return nil
}

// LBOptimize optimizing load balancer.
func (o *OpenStack) LBOptimize(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.LBOptimize()
	return nil
}

// PrepareBasicEnvirionment prepares OpenStack basic environment.
func (o *OpenStack) PrepareBasicEnvirionment(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.PrepareBasicEnvirionment()
	return nil
}

// MariadbCluster deploy MariaDB Cluster.
//  Args: {"HostName": string, "MyIP": string}
func (o *OpenStack) MariadbCluster(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.MariadbCluster()
	return nil
}

// RabbtmqCluster deploy RabbitMQ Cluster.
//  Args: {"HostName": string}
func (o *OpenStack) RabbtmqCluster(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.RabbtmqCluster()
	return nil
}

// Keystone method deploy the Keystone components.
//  Args: {"HostName": string, }
func (o *OpenStack) Keystone(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.Keystone()
	return nil
}

// FormatDiskForSwift formats devices for Swift Storage (sdb1 and sdc1).
//  Args: {"HostName": string}
func (o *OpenStack) FormatDiskForSwift(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.FormatDiskForSwift()
	return nil
}

// SwiftStorage deploy Swift storage.
//  Args: {"HostName": string}
func (o *OpenStack) SwiftStorage(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.SwiftStorage()
	return nil
}

// SwiftProxy deploy Swift proxy HA.
//  Args: {"HostName": string}
func (o *OpenStack) SwiftProxy(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.SwiftProxy()
	return nil
}

// InitSwiftRings initial Swift rings.
//  Args: {"SwiftStorageStorageIP[0]": string, "SwiftStorageStorageIP[1]": string}
func (o *OpenStack) InitSwiftRings(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.InitSwiftRings()
	return nil
}

// DistSwiftRingConf destribute Swift ring configuration files.
func (o *OpenStack) DistSwiftRingConf(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.DistSwiftRingConf()
	return nil
}

// FinalizeSwift finalize Swift installation.
//  Args: {"Hosts": string}
func (o *OpenStack) FinalizeSwift(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.FinalizeSwift()
	return nil
}

// Glance deploy Glance HA.
// Args: {"HostName": string}
func (o *OpenStack) Glance(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.Glance()
	return nil
}

// CephAdmin deploy the Ceph admin node.
func (o *OpenStack) CephAdmin(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.CephAdmin()
	return nil
}

// CephInitMon deploy the Ceph initial monitor.
func (o *OpenStack) CephInitMon(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.CephInitMon()
	return nil
}

// CephClient deploy the Ceph client.
//  Args: {"ClientName": string}
func (o *OpenStack) CephClient(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.CephClient()
	return nil
}

// GetCephKey add Ceph initial monitors and gather the keys.
func (o *OpenStack) GetCephKey(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.GetCephKey()
	return nil
}

// AddOSD add the Ceph OSDs.
//  Args: {"NodeSlice[0]": string, "NodeSlice[1]": string}
func (o *OpenStack) AddOSD(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.AddOSD()
	return nil
}

// AddCephMon add the Ceph monitors.
//  Args: {"Node": string}
func (o *OpenStack) AddCephMon(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.AddCephMon()
	return nil
}

// SyncCephKey copy the Ceph keys to nodes.
//  Args: {"Node": string}
func (o *OpenStack) SyncCephKey(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.SyncCephKey()
	return nil
}

// CephUserPool creates the cinder ceph user and pool name.
func (o *OpenStack) CephUserPool(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.CephUserPool()
	return nil
}

// CinderAPI deploy cinder-api.
//  Args: {"HostName": string}
func (o *OpenStack) CinderAPI(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.CinderAPI()
	return nil
}

// CinderVolume deploy cinder-volume on controller node(ceph backend).
//  Args: {"HostName": string}
func (o *OpenStack) CinderVolume(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.CinderVolume()
	return nil
}

// RestartCephDeps restart volume service dependency to take effect for ceph backend.
func (o *OpenStack) RestartCephDeps(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.RestartCephDeps()
	return nil
}

// NovaController deploy Nova controller.
//  Args: {"HostName": string}
func (o *OpenStack) NovaController(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.NovaController()
	return nil
}

// Dashboard deploy Horizon.
func (o *OpenStack) Dashboard(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.Dashboard()
	return nil
}

// NovaComputes deploy Nova computes.
//  Args: {"HostName": string, "MyIP": string}
func (o *OpenStack) NovaComputes(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.NovaComputes()
	return nil
}

// NovaNetwork deploy legacy networking nova-network(FLATdhcp Only).
//  Args: {"HostName": string, "MyIP": string}
func (o *OpenStack) NovaNetwork(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.NovaNetwork()
	return nil
}

// Heat deploy orchestration components(heat).
//  Args: {"HostName": string}
func (o *OpenStack) Heat(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.Heat()
	return nil
}

// AutoStart fix the service can not auto start when sys booting.
func (o *OpenStack) AutoStart(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.AutoStart()
	return nil
}

// Designate deploy DNS as a Service.
func (o *OpenStack) Designate(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.Designate()
	return nil
}

// KvmToDocker converts kvm to docker(OPTIONAL).
func (o *OpenStack) KvmToDocker(r *http.Request, args *Args, result *Result) error {
	i := provisioning.OpenStack(args)
	*result = i.KvmToDocker()
	return nil
}
