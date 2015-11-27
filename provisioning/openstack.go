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
	StorageIp string
	// Vars: storage_mask
	StorageMask string
	// Vars: storage_network
	StorageNetwork string
	// vars: storage_broadcast
	StorageBroadcast string
	
}
