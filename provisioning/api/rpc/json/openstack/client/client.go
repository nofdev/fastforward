package client

import (
	"bytes"
	"log"
	"net/http"

	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/gorilla/rpc/v2/json"
	"github.com/nofdev/fastforward/provisioning/api/rpc/json/openstack"
)

var url string
var method string
var result openstack.Result

func checkErr(err error) error {
	if err != nil {
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
	resp, err := client.Do(req)
	if err != nil {
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

// ConfigureStorageNetwork takes playback-nic to set up the storage network.
//  Args: {"PlaybackNic.Purge": bool, "PlaybackNic.Public": bool, "PlaybackNic.Private": bool, "PlaybackNic.Host": string, "PlaybackNic.User": string, "PlaybackNic.Address": string, "PlaybackNic.NIC": string, "PlaybackNic.Netmask": string, "PlaybackNic.Gateway": string}
func ConfigureStorageNetwork(args *openstack.Args) error {
	method = "OpenStack.ConfigureStorageNetwork"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// LoadBalancer deploy a HAProxy and Keepalived for OpenStack HA.
//  Args: {"HostName": string, "RouterID": string, "State": string, "Priority": int}
func LoadBalancer(args *openstack.Args) error {
	method = "OpenStack.LoadBalancer"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// LBOptimize optimizing load balancer.
func LBOptimize(args *openstack.Args) error {
	method = "OpenStack.LBOptimize"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// PrepareBasicEnvirionment prepares OpenStack basic environment.
func PrepareBasicEnvirionment(args *openstack.Args) error {
	method = "OpenStack.PrepareBasicEnvirionment"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// MariadbCluster deploy MariaDB Cluster.
//  Args: {"HostName": string, "MyIP": string}
func MariadbCluster(args *openstack.Args) error {
	method = "OpenStack.MariadbCluster"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// RabbtmqCluster deploy RabbitMQ Cluster.
//  Args: {"HostName": string}
func RabbtmqCluster(args *openstack.Args) error {
	method = "OpenStack.RabbtmqCluster"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Keystone method deploy the Keystone components.
//  Args: {"HostName": string, }
func Keystone(args *openstack.Args) error {
	method = "OpenStack.Keystone"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// FormatDiskForSwift formats devices for Swift Storage (sdb1 and sdc1).
//  Args: {"HostName": string}
func FormatDiskForSwift(args *openstack.Args) error {
	method = "OpenStack.FormatDiskForSwift"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// SwiftStorage deploy Swift storage.
//  Args: {"HostName": string}
func SwiftStorage(args *openstack.Args) error {
	method = "OpenStack.SwiftStorage"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// SwiftProxy deploy Swift proxy HA.
//  Args: {"HostName": string}
func SwiftProxy(args *openstack.Args) error {
	method = "OpenStack.SwiftProxy"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// InitSwiftRings initial Swift rings.
//  Args: {"SwiftStorageStorageIP[0]": string, "SwiftStorageStorageIP[1]": string}
func InitSwiftRings(args *openstack.Args) error {
	method = "OpenStack.InitSwiftRings"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DistSwiftRingConf destribute Swift ring configuration files.
func DistSwiftRingConf(args *openstack.Args) error {
	method = "OpenStack.DistSwiftRingConf"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// FinalizeSwift finalize Swift installation.
//  Args: {"Hosts": string}
func FinalizeSwift(args *openstack.Args) error {
	method = "OpenStack.FinalizeSwift"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Glance deploy Glance HA.
// Args: {"HostName": string}
func Glance(args *openstack.Args) error {
	method = "OpenStack.Glance"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CephAdmin deploy the Ceph admin node.
func CephAdmin(args *openstack.Args) error {
	method = "OpenStack.CephAdmin"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CephInitMon deploy the Ceph initial monitor.
func CephInitMon(args *openstack.Args) error {
	method = "OpenStack.CephInitMon"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CephClient deploy the Ceph client.
//  Args: {"ClientName": string}
func CephClient(args *openstack.Args) error {
	method = "OpenStack.CephClient"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetCephKey add Ceph initial monitors and gather the keys.
func GetCephKey(args *openstack.Args) error {
	method = "OpenStack.GetCephKey"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// AddOSD add the Ceph OSDs.
//  Args: {"NodeSlice[0]": string, "NodeSlice[1]": string}
func AddOSD(args *openstack.Args) error {
	method = "OpenStack.AddOSD"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// AddCephMon add the Ceph monitors.
//  Args: {"Node": string}
func AddCephMon(args *openstack.Args) error {
	method = "OpenStack.AddCephMon"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// SyncCephKey copy the Ceph keys to nodes.
//  Args: {"Node": string}
func SyncCephKey(args *openstack.Args) error {
	method = "OpenStack.SyncCephKey"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CephUserPool creates the cinder ceph user and pool name.
func CephUserPool(args *openstack.Args) error {
	method = "OpenStack.CephUserPool"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CinderAPI deploy cinder-api.
//  Args: {"HostName": string}
func CinderAPI(args *openstack.Args) error {
	method = "OpenStack.CinderAPI"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CinderVolume deploy cinder-volume on controller node(ceph backend).
//  Args: {"HostName": string}
func CinderVolume(args *openstack.Args) error {
	method = "OpenStack.CinderVolume"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// RestartCephDeps restart volume service dependency to take effect for ceph backend.
func RestartCephDeps(args *openstack.Args) error {
	method = "OpenStack.RestartCephDeps"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// NovaController deploy Nova controller.
//  Args: {"HostName": string}
func NovaController(args *openstack.Args) error {
	method = "OpenStack.NovaController"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Dashboard deploy Horizon.
func Dashboard(args *openstack.Args) error {
	method = "OpenStack.Dashboard"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// NovaComputes deploy Nova computes.
//  Args: {"HostName": string, "MyIP": string}
func NovaComputes(args *openstack.Args) error {
	method = "OpenStack.NovaComputes"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// NovaNetwork deploy legacy networking nova-network(FLATdhcp Only).
//  Args: {"HostName": string, "MyIP": string}
func NovaNetwork(args *openstack.Args) error {
	method = "OpenStack.NovaNetwork"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Heat deploy orchestration components(heat).
//  Args: {"HostName": string}
func Heat(args *openstack.Args) error {
	method = "OpenStack.Heat"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// AutoStart fix the service can not auto start when sys booting.
func AutoStart(args *openstack.Args) error {
	method = "OpenStack.AutoStart"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Designate deploy DNS as a Service.
func Designate(args *openstack.Args) error {
	method = "OpenStack.Designate"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// KvmToDocker converts kvm to docker(OPTIONAL).
func KvmToDocker(args *openstack.Args) error {
	method = "OpenStack.KvmToDocker"
	err := Do(url, method, args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
