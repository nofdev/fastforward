package provisioning

import "github.com/jiasir/playback/command"

// OpenStack interface takes methods for provision OpenStack.
type OpenStack interface {
	// Prepare OpenStack basic environment.
	PrepareBasicEnvirionment()
	// Using playback-nic to setting the network for storage network.
	ConfigureStorageNetwork()
	// Deploy HAProxy and keepalived.
	LoadBalancer()
	// Deploy MariaDB cluster.
	MariadbCluster()
	// Deploy RabbitMQ cluster.
	RabbtmqCluster()
	// Deploy Keystone HA.
	Keystone()
	// Format the disk for Swift storage, only support sdb1 and sdc1 currently.
	FormatDiskForSwift()
	// Deploy Swift storage.
	SwiftStorage()
	// Deploy Swift proxy HA.
	SwiftProxy()
	// Initial Swift rings.
	InitSwiftRings()
	// Distribute Swift ring configuration files.
	DistSwiftRingConf()
	// Finalize Swift installation.
	FinalizeSwift()
	// Deploy Glance HA.
	Glance()
	// Deploy Ceph admin node.
	CephAdmin()
	// Deploy the Ceph initial monitor.
	CephInitMon()
	// Deploy the Ceph clients.
	CephClient()
	// Add Ceph initial monitor(s) and gather the keys.
	GetCephKey()
	// Add Ceph OSDs.
	AddOSD()
	// Add Ceph monitors.
	AddCephMon()
	// Copy the Ceph keys to nodes.
	SyncCephKey()
	// Create the cinder ceph user and pool name.
	CephUserPool()
	// Deploy cinder-api.
	CinderAPI()
	// Deploy cinder-volume on controller node(Ceph backend).
	CinderVolume()
	// Restart volume service dependency to take effect for ceph backend.
	RestartCephDeps()
	// Deploy Nova controller.
	NovaController()
	// Deploy Horizon.
	Dashboard()
	// Deploy Nova computes.
	NovaComputes()
	// Deploy Legacy networking nova-network(FlatDHCP Only).
	NovaNetwork()
	// Deploy Orchestration components(heat).
	Heat()
	// Enable service auto start
	AutoStart()
	// Deploy Dns as a Service
	Designate()
	// Convert kvm to Docker(OPTIONAL)
	KvmToDocker()
}

// ExtraVars takes playback command line arguments.
type ExtraVars struct {
	// Ansible Playbook *.yml
	Playbook string
	// Vars: node_name
	NodeName string
	// Vars: host
	HostIP string
	// Vars: storage_ip
	StorageIP string
	// Vars: storage_mask
	StorageMask string
	// Vars: storage_network
	StorageNetwork string
	// Vars: storage_broadcast
	StorageBroadcast string
	// Command line playback-nic
	PlaybackNic PlaybackNic
	// Vars: host
	HostName string
	// Vars: router_id
	RouterID string
	// Vars: state
	State string
	// Vars: priority
	Priority int
	// Python scripts *.py
	PythonScript string
	// Vars: my_ip
	MyIP string
	// Vars: my_storage_ip
	MyStorageIP string
	// Vars: swift_storage_storage_ip
	SwiftStorageStorageIP string
	// Vars: device_name
	DeviceName string
	// Vars: device_weight
	DeviceWeight int
	// Vars: hosts
	Hosts string
	// Vars: client
	ClientName string
	// Vars: disk
	Disk string
	// Vars: partition
	Partition string
}

// PlaybackNic using playback-nic command instaed of openstack_interface.yml
type PlaybackNic struct {
	// Args: purge
	Purge bool
	// Args: public
	Public bool
	// Args: private
	Private bool
	// Args: host
	Host string
	// Args: user
	User string
	// Args: address
	Address string
	// Args: nic
	NIC string
	// Args: netmask
	Netmask string
	// Args: gateway
	Gateway string
	// Args: dns-nameservers
	DNS string
}

// ConfigureStorageNetwork takes playback-nic to set up the storage network.
// Playback examples:
// Purge the configuration and set address to 192.169.151.19 for eth1 of host 192.169.150.19 as public interface:
//	playback-nic --purge --public --host 192.169.150.19 --user ubuntu --address 192.169.151.19 --nic eth1 --netmask 255.255.255.0 --gateway 192.169.151.1 --dns-nameservers "192.169.11.11 192.169.11.12"
//Setting address to 192.168.1.12 for eth2 of host 192.169.150.19 as private interface:
//	playback-nic --private --host 192.169.150.19 --user ubuntu --address 192.168.1.12 --nic eth2 --netmask 255.255.255.0
func (vars ExtraVars) ConfigureStorageNetwork() error {
	if vars.PlaybackNic.Purge {
		if vars.PlaybackNic.Public {
			command.ExecuteWithOutput("playback-nic", "--purge", "--public", "--host", vars.PlaybackNic.Host, "--user", vars.PlaybackNic.User, "--address", vars.PlaybackNic.Address, "--nic", vars.PlaybackNic.NIC, "--netmask", vars.PlaybackNic.Netmask, "--gateway", vars.PlaybackNic.Gateway, "--dns-nameservers", vars.PlaybackNic.DNS)
		}
	}
	if vars.PlaybackNic.Private {
		command.ExecuteWithOutput("playback-nic", "--private", "--host", vars.PlaybackNic.Host, "--user", vars.PlaybackNic.Host, "--address", vars.PlaybackNic.Address, "--nic", vars.PlaybackNic.NIC, "--netmask", vars.PlaybackNic.Netmask)
	}
	return nil
}

// LoadBalancer deploy a HAProxy and Keepalived for OpenStack HA.
func (vars ExtraVars) LoadBalancer() error {
	return nil
}

// PrepareBasicEnvirionment prepares OpenStack basic environment.
func (vars ExtraVars) PrepareBasicEnvirionment() error {
	return nil
}

// MariadbCluster deploy MariaDB Cluster.
func (vars ExtraVars) MariadbCluster() error {
	return nil
}

// RabbtmqCluster deploy RabbitMQ Cluster.
func (vars ExtraVars) RabbtmqCluster() error {
	return nil
}

// Keystone method deploy the Keystone components.
func (vars ExtraVars) Keystone() error {
	return nil
}

// FormatDiskForSwift formats devices for Swift Storage (sdb1 and sdc1).
func (vars ExtraVars) FormatDiskForSwift() error {
	return nil
}

// SwiftStorage deploy Swift storage.
func (vars ExtraVars) SwiftStorage() error {
	return nil
}

// SwiftProxy deploy Swift proxy HA.
func (vars ExtraVars) SwiftProxy() error {
	return nil
}

// InitSwiftRings initial Swift rings.
func (vars ExtraVars) InitSwiftRings() error {
	return nil
}

// DistSwiftRingConf destribute Swift ring configuration files.
func (vars ExtraVars) DistSwiftRingConf() error {
	return nil
}

// FinalizeSwift finalize Swift installation.
func (vars ExtraVars) FinalizeSwift() error {
	return nil
}

// Glance deploy Glance HA.
func (vars ExtraVars) Glance() error {
	return nil
}

// CephAdmin deploy the Ceph admin node.
func (vars ExtraVars) CephAdmin() error {
	return nil
}

// CephInitMon deploy the Ceph initial monitor.
func (vars ExtraVars) CephInitMon() error {
	return nil
}

// CephClient deploy the Ceph client.
func (vars ExtraVars) CephClient() error {
	return nil
}

// GetCephKey add Ceph initial monitors and gather the keys.
func (vars ExtraVars) GetCephKey() error {
	return nil
}

// AddOSD add the Ceph OSDs.
func (vars ExtraVars) AddOSD() error {
	return nil
}

// AddCephMon add the Ceph monitors.
func (vars ExtraVars) AddCephMon() error {
	return nil
}

// SyncCephKey copy the Ceph keys to nodes.
func (vars ExtraVars) SyncCephKey() error {
	return nil
}

// CephUserPool creates the cinder ceph user and pool name.
func (vars ExtraVars) CephUserPool() error {
	return nil
}

// CinderAPI deploy cinder-api.
func (vars ExtraVars) CinderAPI() error {
	return nil
}

// CinderVolume deploy cinder-volume on controller node(ceph backend).
func (vars ExtraVars) CinderVolume() error {
	return nil
}

// RestartCephDeps restart volume service dependency to take effect for ceph backend.
func (vars ExtraVars) RestartCephDeps() error {
	return nil
}

// NovaController deploy Nova controller.
func (vars ExtraVars) NovaController() error {
	return nil
}

// Dashboard deploy Horizon.
func (vars ExtraVars) Dashboard() error {
	return nil
}

// NovaComputes deploy Nova computes.
func (vars ExtraVars) NovaComputes() error {
	return nil
}

// NovaNetwork deploy legacy networking nova-network(FLATdhcp Only).
func (vars ExtraVars) NovaNetwork() error {
	return nil
}

// Heat deploy orchestration components(heat).
func (vars ExtraVars) Heat() error {
	return nil
}

// AutoStart fix the service can not auto start when sys booting.
func (vars ExtraVars) AutoStart() error {
	return nil
}

// Designate deploy DNS as a Service.
func (vars ExtraVars) Designate() error {
	return nil
}

// KvmToDocker converts kvm to docker(OPTIONAL).
func (vars ExtraVars) KvmToDocker() error {
	return nil
}
