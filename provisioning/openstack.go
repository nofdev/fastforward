package provisioning

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
	CinderApi()
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
