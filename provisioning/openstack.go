package provisioning

import "github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/jiasir/playback/command"

// OpenStack interface takes methods for provision OpenStack.
type OpenStack interface {
	// Prepare OpenStack basic environment.
	PrepareBasicEnvirionment() error
	// Using playback-nic to setting the network for storage network.
	ConfigureStorageNetwork() error
	// Deploy HAProxy and keepalived.
	LoadBalancer() error
	// LBOptimize optimizing load balancer.
	LBOptimize() error
	// Deploy MariaDB cluster.
	MariadbCluster() error
	// Deploy RabbitMQ cluster.
	RabbtmqCluster() error
	// Deploy Keystone HA.
	Keystone() error
	// Format the disk for Swift storage, only support sdb1 and sdc1 currently.
	FormatDiskForSwift() error
	// Deploy Swift storage.
	SwiftStorage() error
	// Deploy Swift proxy HA.
	SwiftProxy() error
	// Initial Swift rings.
	InitSwiftRings() error
	// Distribute Swift ring configuration files.
	DistSwiftRingConf() error
	// Finalize Swift installation.
	FinalizeSwift() error
	// Deploy Glance HA.
	Glance() error
	// Deploy Ceph admin node.
	CephAdmin() error
	// Deploy the Ceph initial monitor.
	CephInitMon() error
	// Deploy the Ceph clients.
	CephClient() error
	// Add Ceph initial monitor(s) and gather the keys.
	GetCephKey() error
	// Add Ceph OSDs.
	AddOSD() error
	// Add Ceph monitors.
	AddCephMon() error
	// Copy the Ceph keys to nodes.
	SyncCephKey() error
	// Create the cinder ceph user and pool name.
	CephUserPool() error
	// Deploy cinder-api.
	CinderAPI() error
	// Deploy cinder-volume on controller node(Ceph backend).
	CinderVolume() error
	// Restart volume service dependency to take effect for ceph backend.
	RestartCephDeps() error
	// Deploy Nova controller.
	NovaController() error
	// Deploy Horizon.
	Dashboard() error
	// Deploy Nova computes.
	NovaComputes() error
	// Deploy Legacy networking nova-network(FlatDHCP Only).
	NovaNetwork() error
	// Deploy Orchestration components(heat).
	Heat() error
	// Enable service auto start
	AutoStart() error
	// Deploy Dns as a Service
	Designate() error
	// Convert kvm to Docker(OPTIONAL)
	KvmToDocker() error
}

// ExtraVars takes playback command line arguments.
type ExtraVars struct {
	// Ansible Playbook *.yml
	Playbook string
	// Vars: node_name
	NodeName string
	// Vars: node
	NodeSlice []string
	// Vars: node
	Node string
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
	Priority string
	// Python scripts *.py
	PythonScript string
	// Vars: my_ip
	MyIP string
	// Vars: my_storage_ip
	MyStorageIP string
	// Vars: swift_storage_storage_ip
	SwiftStorageStorageIP []string
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
// The method takes the following command of Playback:
//  playback --ansible 'openstack_haproxy.yml --extra-vars "host=lb01 router_id=lb01 state=MASTER priority=150" -vvvv'
//  playback --ansible 'openstack_haproxy.yml --extra-vars "host=lb02 router_id=lb02 state=SLAVE priority=100" -vvvv'
func (vars ExtraVars) LoadBalancer() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_haproxy.yml", "--extra-vars", "host="+vars.HostName, "router_id="+vars.RouterID, "state="+vars.State, "priority="+vars.Priority, "-vvvv")
	return nil
}

// LBOptimize optimizing load balancer.
// the method takes the floowing command of Playback:
//  python patch-limits.py
func (vars ExtraVars) LBOptimize() error {
	command.ExecuteWithOutput("python patch-limits.py")
	return nil
}

// PrepareBasicEnvirionment prepares OpenStack basic environment.
// The method takes the following command of Playback:
//  playback --ansible 'openstack_basic_environment.yml -vvvv'
func (vars ExtraVars) PrepareBasicEnvirionment() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_basic_environment.yml", "-vvvv")
	return nil
}

// MariadbCluster deploy MariaDB Cluster.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_mariadb.yml --extra-vars "host=controller01 my_ip=192.169.151.19" -vvvv'
//  playback --ansible 'openstack_mariadb.yml --extra-vars "host=controller02 my_ip=192.169.151.17" -vvvv'
//  python keepalived.py
func (vars ExtraVars) MariadbCluster() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_mariadb.yml", "--extra-vars", "host="+vars.HostName, "my_ip="+vars.MyIP, "-vvvv")
	if vars.HostName == "controller02" {
		command.ExecuteWithOutput("python keepalived.py")
	}
	return nil
}

// RabbtmqCluster deploy RabbitMQ Cluster.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_rabbitmq.yml --extra-vars "host=controller01" -vvvv'
//  playback --ansible 'openstack_rabbitmq.yml --extra-vars "host=controller02" -vvvv'
func (vars ExtraVars) RabbtmqCluster() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_rabbitmq.yml", "--extra-vars", "host="+vars.HostName, "-vvvv")
	return nil
}

// Keystone method deploy the Keystone components.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_keystone.yml --extra-vars "host=controller01" -vvvv'
//  playback --ansible 'openstack_keystone.yml --extra-vars "host=controller02" -vvvv'
func (vars ExtraVars) Keystone() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_keystone.yml", "--extra-vars", "host="+vars.HostName, "-vvvv")
	return nil
}

// FormatDiskForSwift formats devices for Swift Storage (sdb1 and sdc1).
// Each of the swift nodes, /dev/sdb1 and /dev/sdc1, must contain a suitable partition table with one partition occupying the entire device.
// Although the Object Storage service supports any file system with extended attributes (xattr), testing and benchmarking indicate the best performance and reliability on XFS.
// The method takes the folowing commands of Playback:
//  playback --ansible 'openstack_storage_partitions.yml --extra-vars "host=compute05" -vvvv'
//  playback --ansible 'openstack_storage_partitions.yml --extra-vars "host=compute06" -vvvv'
func (vars ExtraVars) FormatDiskForSwift() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_storage_partitions.yml", "--extra-vars", "host="+vars.HostName, "-vvvv")
	return nil
}

// SwiftStorage deploy Swift storage.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_swift_storage.yml --extra-vars "host=compute05 my_storage_ip=192.168.1.16" -vvvv'
//  playback --ansible 'openstack_swift_storage.yml --extra-vars "host=compute06 my_storage_ip=192.168.1.15" -vvvv'
func (vars ExtraVars) SwiftStorage() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_swift_storage.yml", "--extra-vars", "host="+vars.HostName, "my_storage_ip="+vars.MyStorageIP, "-vvvv")
	return nil
}

// SwiftProxy deploy Swift proxy HA.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_swift_proxy.yml --extra-vars "host=controller01" -vvvv'
//  playback --ansible 'openstack_swift_proxy.yml --extra-vars "host=controller02" -vvvv'
func (vars ExtraVars) SwiftProxy() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_swift_proxy.yml", "--extra-vars", "host="+vars.HostName, "-vvvv")
	return nil
}

// InitSwiftRings initial Swift rings.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_swift_builder_file.yml -vvvv'
//  playback --ansible 'openstack_swift_add_node_to_the_ring.yml --extra-vars "swift_storage_storage_ip=192.168.1.16 device_name=sdb1 device_weight=100" -vvvv'
//  playback --ansible 'openstack_swift_add_node_to_the_ring.yml --extra-vars "swift_storage_storage_ip=192.168.1.16 device_name=sdc1 device_weight=100" -vvvv'
//  playback --ansible 'openstack_swift_add_node_to_the_ring.yml --extra-vars "swift_storage_storage_ip=192.168.1.15 device_name=sdb1 device_weight=100" -vvvv'
//  playback --ansible 'openstack_swift_add_node_to_the_ring.yml --extra-vars "swift_storage_storage_ip=192.168.1.15 device_name=sdc1 device_weight=100" -vvvv'
//  playback --ansible 'openstack_swift_rebalance_ring.yml -vvvv'
func (vars ExtraVars) InitSwiftRings() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_swift_builder_file.yml", "-vvvv")
	command.ExecuteWithOutput("playback", "--ansible", "openstack_swift_add_node_to_the_ring.yml", "--extra-vars", "swift_storage_storage_ip="+vars.SwiftStorageStorageIP[0], "device_name=sdb1", "device_weight=100", "-vvvv")
	command.ExecuteWithOutput("playback", "--ansible", "openstack_swift_add_node_to_the_ring.yml", "--extra-vars", "swift_storage_storage_ip="+vars.SwiftStorageStorageIP[0], "device_name=sdc1", "device_weight=100", "-vvvv")
	command.ExecuteWithOutput("playback", "--ansible", "openstack_swift_add_node_to_the_ring.yml", "--extra-vars", "swift_storage_storage_ip="+vars.SwiftStorageStorageIP[1], "device_name=sdb1", "device_weight=100", "-vvvv")
	command.ExecuteWithOutput("playback", "--ansible", "openstack_swift_add_node_to_the_ring.yml", "--extra-vars", "swift_storage_storage_ip="+vars.SwiftStorageStorageIP[1], "device_name=sdc1", "device_weight=100", "-vvvv")
	command.ExecuteWithOutput("playback", "--ansible", "openstack_swift_rebalance_ring.yml", "-vvvv")
	return nil
}

// DistSwiftRingConf destribute Swift ring configuration files.
// Copy the account.ring.gz, container.ring.gz, and object.ring.gz files to the /etc/swift directory on each storage node and any additional nodes running the proxy service.
func (vars ExtraVars) DistSwiftRingConf() error {
	// Playback have not implement the automation currently.
	// TODO: Distribute ring configuration files.
	return nil
}

// FinalizeSwift finalize Swift installation.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_swift_finalize_installation.yml --extra-vars "hosts=swift_proxy" -vvvv'
//  playback --ansible 'openstack_swift_finalize_installation.yml --extra-vars "hosts=swift_storage" -vvvv'
func (vars ExtraVars) FinalizeSwift() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_swift_finalize_installation.yml", "--extra-vars", "hosts="+vars.Hosts, "-vvvv")
	return nil
}

// Glance deploy Glance HA.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_glance.yml --extra-vars "host=controller01" -vvvv'
//  playback --ansible 'openstack_glance.yml --extra-vars "host=controller02" -vvvv'
func (vars ExtraVars) Glance() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_glance.yml", "--extra-vars", "host="+vars.HostName, "-vvvv")
	return nil
}

// CephAdmin deploy the Ceph admin node.
// Ensure the admin node must be have password-less SSH access to Ceph nodes. When ceph-deploy logs in to a Ceph node as a user, that particular user must have passwordless sudo privileges.
// Copy SSH public key to each Ceph node from Ceph admin node:
//  ssh-keygen
//  ssh-copy-id ubuntu@ceph_node
// The method takes the following command of Playback:
//  playback --ansible 'openstack_ceph_admin.yml -vvvv'
func (vars ExtraVars) CephAdmin() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_admin.yml", "-vvvv")
	return nil
}

// CephInitMon deploy the Ceph initial monitor.
// The method takes the following command of Playback:
//  playback --ansible 'openstack_ceph_initial_mon.yml -vvvv'
func (vars ExtraVars) CephInitMon() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_initial_mon.yml", "-vvvv")
	return nil
}

// CephClient deploy the Ceph client.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_ceph_client.yml --extra-vars "client=controller01" -vvvv'
//  playback --ansible 'openstack_ceph_client.yml --extra-vars "client=controller02" -vvvv'
//  playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute01" -vvvv'
//  playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute02" -vvvv'
//  playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute03" -vvvv'
//  playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute04" -vvvv'
//  playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute05" -vvvv'
//  playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute06" -vvvv'
func (vars ExtraVars) CephClient() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_client.yml", "--extra-vars", "client="+vars.ClientName, "-vvvv")
	return nil
}

// GetCephKey add Ceph initial monitors and gather the keys.
// The method takes the following command of Playback:
//  playback --ansible 'openstack_ceph_gather_keys.yml -vvvv'
func (vars ExtraVars) GetCephKey() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_gather_keys.yml", "-vvvv")
	return nil
}

// AddOSD add the Ceph OSDs.
// Only sopport sdb and sdc.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute01 disk=sdb partition=sdb1" -vvvv'
//  playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute01 disk=sdc partition=sdc1" -vvvv'
//  playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute02 disk=sdb partition=sdb1" -vvvv'
//  playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute02 disk=sdc partition=sdc1" -vvvv'
//  playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute03 disk=sdb partition=sdb1" -vvvv'
//  playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute03 disk=sdc partition=sdc1" -vvvv'
//  playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute04 disk=sdb partition=sdb1" -vvvv'
//  playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute04 disk=sdc partition=sdc1" -vvvv'
// Only support two nodes for sdb and sdc currently.
func (vars ExtraVars) AddOSD() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_osd.yml", "--extra-vars", "node="+vars.NodeSlice[0], "disk=sdb", "partition=sdb1", "-vvvv")
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_osd.yml", "--extra-vars", "node="+vars.NodeSlice[0], "disk=sdc", "partition=sdc1", "-vvvv")
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_osd.yml", "--extra-vars", "node="+vars.NodeSlice[1], "disk=sdb", "partition=sdb1", "-vvvv")
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_osd.yml", "--extra-vars", "node="+vars.NodeSlice[1], "disk=sdc", "partition=sdc1", "-vvvv")
	return nil
}

// AddCephMon add the Ceph monitors.
// The method takes the following command of Playback:
//  playback --ansible 'openstack_ceph_mon.yml --extra-vars "node=compute01" -vvvv'
//  playback --ansible 'openstack_ceph_mon.yml --extra-vars "node=compute02" -vvvv'
//  playback --ansible 'openstack_ceph_mon.yml --extra-vars "node=compute03" -vvvv'
//  playback --ansible 'openstack_ceph_mon.yml --extra-vars "node=compute04" -vvvv'
func (vars ExtraVars) AddCephMon() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_mon.yml", "--extra-vars", "node="+vars.Node, "-vvvv")
	return nil
}

// SyncCephKey copy the Ceph keys to nodes.
// Copy the configuration file and admin key to your admin node and your Ceph Nodes so that you can use the ceph CLI without having to specify the monitor address and ceph.client.admin.keyring each time you execute a command.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=controller01" -vvvv'
//  playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=controller02" -vvvv'
//  playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute01" -vvvv'
//  playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute02" -vvvv'
//  playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute03" -vvvv'
//  playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute04" -vvvv'
//  playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute05" -vvvv'
//  playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute06" -vvvv'
func (vars ExtraVars) SyncCephKey() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_copy_keys.yml", "--extra-vars", "node="+vars.Node, "-vvvv")
	return nil
}

// CephUserPool creates the cinder ceph user and pool name.
// The method takes the following command of Playback:
//  playback --ansible 'openstack_ceph_cinder_pool_user.yml -vvvv'
func (vars ExtraVars) CephUserPool() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_ceph_cinder_pool_user.yml", "-vvvv")
	return nil
}

// CinderAPI deploy cinder-api.
// The method takes the following command of Playback:
//  playback --ansible 'openstack_cinder_api.yml --extra-vars "host=controller01" -vvvv'
//  playback --ansible 'openstack_cinder_api.yml --extra-vars "host=controller02" -vvvv'
func (vars ExtraVars) CinderAPI() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_cinder_api.yml", "--extra-vars", "host="+vars.HostName, "-vvvv")
	return nil
}

// CinderVolume deploy cinder-volume on controller node(ceph backend).
// The method takes the following command of Playback:
//  playback --ansible 'openstack_cinder_volume_ceph.yml --extra-vars "host=controller01" -vvvv'
//  playback --ansible 'openstack_cinder_volume_ceph.yml --extra-vars "host=controller02" -vvvv'
// Copy the ceph.client.cinder.keyring from ceph-admin node to /etc/ceph/ceph.client.cinder.keyring of cinder volume nodes and nova-compute nodes to using the ceph client:
//  ceph auth get-or-create client.cinder | ssh ubuntu@controller01 sudo tee /etc/ceph/ceph.client.cinder.keyring
//  ceph auth get-or-create client.cinder | ssh ubuntu@controller02 sudo tee /etc/ceph/ceph.client.cinder.keyring
//  ceph auth get-or-create client.cinder | ssh ubuntu@compute01 sudo tee /etc/ceph/ceph.client.cinder.keyring
//  ceph auth get-or-create client.cinder | ssh ubuntu@compute02 sudo tee /etc/ceph/ceph.client.cinder.keyring
//  ceph auth get-or-create client.cinder | ssh ubuntu@compute03 sudo tee /etc/ceph/ceph.client.cinder.keyring
//  ceph auth get-or-create client.cinder | ssh ubuntu@compute04 sudo tee /etc/ceph/ceph.client.cinder.keyring
//  ceph auth get-or-create client.cinder | ssh ubuntu@compute05 sudo tee /etc/ceph/ceph.client.cinder.keyring
//  ceph auth get-or-create client.cinder | ssh ubuntu@compute06 sudo tee /etc/ceph/ceph.client.cinder.keyring
func (vars ExtraVars) CinderVolume() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_cinder_volume_ceph.yml", "--extra-vars", "host="+vars.HostName, "-vvvv")
	return nil
}

// RestartCephDeps restart volume service dependency to take effect for ceph backend.
// The method takes the following command of Playback:
//  python restart_cindervol_deps.py ubuntu@controller01 ubuntu@controller02
// Only support controller01 and controller02 currently.
func (vars ExtraVars) RestartCephDeps() error {
	command.ExecuteWithOutput("python", "restart_cindervol_deps.py", "ubuntu@controller01", "ubuntu@controller02")
	return nil
}

// NovaController deploy Nova controller.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_compute_controller.yml --extra-vars "host=controller01" -vvvv'
//  playback --ansible 'openstack_compute_controller.yml --extra-vars "host=controller02" -vvvv'
func (vars ExtraVars) NovaController() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_compute_controller.yml", "--extra-vars", "host="+vars.HostName, "-vvvv")
	return nil
}

// Dashboard deploy Horizon.
// The method takes the following command of Playback:
//  playback --ansible 'openstack_horizon.yml -vvvv'
func (vars ExtraVars) Dashboard() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_horizon.yml", "-vvvv")
	return nil
}

// NovaComputes deploy Nova computes.
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute01 my_ip=192.169.151.16" -vvvv'
//  playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute02 my_ip=192.169.151.22" -vvvv'
//  playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute03 my_ip=192.169.151.18" -vvvv'
//  playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute04 my_ip=192.169.151.25" -vvvv'
//  playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute05 my_ip=192.169.151.12" -vvvv'
//  playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute06 my_ip=192.169.151.14" -vvvv'
func (vars ExtraVars) NovaComputes() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_compute_node.yml", "--extra-vars", "host="+vars.HostName, "my_ip="+vars.MyIP, "-vvvv")
	return nil
}

// NovaNetwork deploy legacy networking nova-network(FLATdhcp Only).
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute01 my_ip=192.169.151.16" -vvvv'
//  playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute02 my_ip=192.169.151.22" -vvvv'
//  playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute03 my_ip=192.169.151.18" -vvvv'
//  playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute04 my_ip=192.169.151.25" -vvvv'
//  playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute05 my_ip=192.169.151.12" -vvvv'
//  playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute06 my_ip=192.169.151.14" -vvvv'
// Create initial network. For example, using an exclusive slice of 172.16.0.0/16 with IP address range 172.16.0.1 to 172.16.255.254:
//  nova network-create ext-net --bridge br100 --multi-host T --fixed-range-v4 172.16.0.0/16
//  nova floating-ip-bulk-create --pool ext-net 192.169.151.65/26
//  nova floating-ip-bulk-list
// Extend the demo-net pool:
//  nova floating-ip-bulk-create --pool ext-net 192.169.151.128/26
//  nova floating-ip-bulk-create --pool ext-net 192.169.151.192/26
//  nova floating-ip-bulk-list
func (vars ExtraVars) NovaNetwork() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_nova_network_compute.yml", "--extra-vars", "host="+vars.HostName, "my_ip="+vars.MyIP, "-vvvv")
	return nil
}

// Heat deploy orchestration components(heat).
// The method takes the following commands of Playback:
//  playback --ansible 'openstack_heat_controller.yml --extra-vars "host=controller01" -vvvv'
//  playback --ansible 'openstack_heat_controller.yml --extra-vars "host=controller02" -vvvv'
func (vars ExtraVars) Heat() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_heat_controller.yml", "--extra-vars", "host="+vars.HostName, "-vvvv")
	return nil
}

// AutoStart fix the service can not auto start when sys booting.
// The method takes the following command of Playback:
//  python patch-autostart.py
func (vars ExtraVars) AutoStart() error {
	command.ExecuteWithOutput("python", "patch-autostart.py")
	return nil
}

// Designate deploy DNS as a Service.
// The method takes the following command of Playback:
//  playback --ansible openstack_dns.yml -vvvv
// Execute on controller01:
//  bash /mnt/designate-keystone-setup
//  nohup designate-central > /dev/null 2>&1 &
//  nohup designate-api > /dev/null 2>&1 &
func (vars ExtraVars) Designate() error {
	command.ExecuteWithOutput("playback", "--ansible", "openstack_dns.yml", "-vvvv")
	return nil
}

// KvmToDocker converts kvm to docker(OPTIONAL).
// The method takes the following command of Playback:
//  playback --novadocker --user ubuntu --hosts compute06
func (vars ExtraVars) KvmToDocker() error {
	command.ExecuteWithOutput("playback", "--novadocker", "--user", "ubuntu", "--hosts", "compute06")
	return nil
}
