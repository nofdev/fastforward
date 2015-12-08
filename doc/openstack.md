# OpenStack Provisioning
Provision an OpenStack environment.

#### Define a inventory file
The inventory file at `inventory`, the default setting is the Vagrant testing node. You can according to your environment to change parameters.
    
#### To define your variables in vars/openstack
The `vars/openstack/openstack.yml` is all the parameters.
* openstack.yml


### Deploy Playback
    pip install playback
    
### initial the configurations
    playback --init
    cd .playback

### Configure storage network(DEPRECATED! Using playback-nic instead of it)
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=lb01 storage_ip=192.168.1.10 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=lb02 storage_ip=192.168.1.11 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=controller01 storage_ip=192.168.1.12 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=controller02 storage_ip=192.168.1.13 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=compute01 storage_ip=192.168.1.14 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=compute02 storage_ip=192.168.1.20 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=compute03 storage_ip=192.168.1.18 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=compute04 storage_ip=192.168.1.17 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=compute05 storage_ip=192.168.1.16 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
    ff playback --ansible 'openstack_interfaces.yml --extra-vars "node_name=compute06 storage_ip=192.168.1.15 storage_mask=255.255.255.0 storage_network=192.168.1.0 storage_broadcast=192.168.1.255" -vvvv'
 
### Using playback-nic
Purge the configuration and set address to 192.169.151.19 for eth1 of host 192.169.150.19 as public interface:
    
    playback-nic --purge --public --host 192.169.150.19 --user ubuntu --address 192.169.151.19 --nic eth1 --netmask 255.255.255.0 --gateway 192.169.151.1 --dns-nameservers "192.169.11.11 192.169.11.12"
Setting address to 192.168.1.12 for eth2 of host 192.169.150.19 as private interface:

    playback-nic --private --host 192.169.150.19 --user ubuntu --address 192.168.1.12 --nic eth2 --netmask 255.255.255.0

### HAProxy and Keepalived
    ff playback --ansible 'openstack_haproxy.yml --extra-vars "host=lb01 router_id=lb01 state=MASTER priority=150" -vvvv'
    ff playback --ansible 'openstack_haproxy.yml --extra-vars "host=lb02 router_id=lb02 state=SLAVE priority=100" -vvvv'
    python patch-limits.py
    
### Prepare OpenStack basic environment
    ff playback --ansible 'openstack_basic_environment.yml -vvvv'

### MariaDB Cluster
    ff playback --ansible 'openstack_mariadb.yml --extra-vars "host=controller01 my_ip=192.169.151.19" -vvvv'
    ff playback --ansible 'openstack_mariadb.yml --extra-vars "host=controller02 my_ip=192.169.151.17" -vvvv'
    python keepalived.py

### RabbitMQ Cluster
    ff playback --ansible 'openstack_rabbitmq.yml --extra-vars "host=controller01" -vvvv'
    ff playback --ansible 'openstack_rabbitmq.yml --extra-vars "host=controller02" -vvvv'

### Keystone
    ff playback --ansible 'openstack_keystone.yml --extra-vars "host=controller01" -vvvv'
    ff playback --ansible 'openstack_keystone.yml --extra-vars "host=controller02" -vvvv'

### Format devices for Swift Storage (sdb1 and sdc1)
Each of the swift nodes, /dev/sdb1 and /dev/sdc1, must contain a suitable partition table with one partition occupying the entire device. Although the Object Storage service supports any file system with extended attributes (xattr), testing and benchmarking indicate the best performance and reliability on XFS.

    ff playback --ansible 'openstack_storage_partitions.yml --extra-vars "host=compute05" -vvvv'
    ff playback --ansible 'openstack_storage_partitions.yml --extra-vars "host=compute06" -vvvv'

### Swift Storage
    ff playback --ansible 'openstack_swift_storage.yml --extra-vars "host=compute05 my_storage_ip=192.168.1.16" -vvvv'
    ff playback --ansible 'openstack_swift_storage.yml --extra-vars "host=compute06 my_storage_ip=192.168.1.15" -vvvv'

    
### Swift Proxy
    ff playback --ansible 'openstack_swift_proxy.yml --extra-vars "host=controller01" -vvvv'
    ff playback --ansible 'openstack_swift_proxy.yml --extra-vars "host=controller02" -vvvv'                    
    
### Initial swift rings
    ff playback --ansible 'openstack_swift_builder_file.yml -vvvv'
    ff playback --ansible 'openstack_swift_add_node_to_the_ring.yml --extra-vars "swift_storage_storage_ip=192.168.1.16 device_name=sdb1 device_weight=100" -vvvv'
    ff playback --ansible 'openstack_swift_add_node_to_the_ring.yml --extra-vars "swift_storage_storage_ip=192.168.1.16 device_name=sdc1 device_weight=100" -vvvv'
    ff playback --ansible 'openstack_swift_add_node_to_the_ring.yml --extra-vars "swift_storage_storage_ip=192.168.1.15 device_name=sdb1 device_weight=100" -vvvv'
    ff playback --ansible 'openstack_swift_add_node_to_the_ring.yml --extra-vars "swift_storage_storage_ip=192.168.1.15 device_name=sdc1 device_weight=100" -vvvv'
    ff playback --ansible 'openstack_swift_rebalance_ring.yml -vvvv'
    
### Distribute ring configuration files
Copy the `account.ring.gz`, `container.ring.gz`, and `object.ring.gz` files to the `/etc/swift` directory on each storage node and any additional nodes running the proxy service.

### Finalize swift installation
    ff playback --ansible 'openstack_swift_finalize_installation.yml --extra-vars "hosts=swift_proxy" -vvvv'
    ff playback --ansible 'openstack_swift_finalize_installation.yml --extra-vars "hosts=swift_storage" -vvvv'

### Glance (Swift backend)
    ff playback --ansible 'openstack_glance.yml --extra-vars "host=controller01" -vvvv'
    ff playback --ansible 'openstack_glance.yml --extra-vars "host=controller02" -vvvv'
    
### To deploy the Ceph admin node
Ensure the admin node must be have password-less SSH access to Ceph nodes. When ceph-deploy logs in to a Ceph node as a user, that particular user must have passwordless sudo privileges.

Copy SSH public key to each Ceph node from Ceph admin node
    
    ssh-keygen
    ssh-copy-id ubuntu@ceph_node

Deploy the Ceph admin node

    ff playback --ansible 'openstack_ceph_admin.yml -vvvv'

### To deploy the Ceph initial monitor
    ff playback --ansible 'openstack_ceph_initial_mon.yml -vvvv'
    
### To deploy the Ceph clients
    ff playback --ansible 'openstack_ceph_client.yml --extra-vars "client=controller01" -vvvv'
    ff playback --ansible 'openstack_ceph_client.yml --extra-vars "client=controller02" -vvvv'
    ff playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute01" -vvvv'
    ff playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute02" -vvvv'
    ff playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute03" -vvvv'
    ff playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute04" -vvvv'
    ff playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute05" -vvvv'
    ff playback --ansible 'openstack_ceph_client.yml --extra-vars "client=compute06" -vvvv'


### To add Ceph initial monitor(s) and gather the keys
    ff playback --ansible 'openstack_ceph_gather_keys.yml -vvvv'

### To add Ceph OSDs
    ff playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute01 disk=sdb partition=sdb1" -vvvv'
    ff playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute01 disk=sdc partition=sdc1" -vvvv'
    ff playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute02 disk=sdb partition=sdb1" -vvvv'
    ff playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute02 disk=sdc partition=sdc1" -vvvv'
    ff playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute03 disk=sdb partition=sdb1" -vvvv'
    ff playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute03 disk=sdc partition=sdc1" -vvvv'
    ff playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute04 disk=sdb partition=sdb1" -vvvv'
    ff playback --ansible 'openstack_ceph_osd.yml --extra-vars "node=compute04 disk=sdc partition=sdc1" -vvvv'


### To add Ceph monitors
    ff playback --ansible 'openstack_ceph_mon.yml --extra-vars "node=compute01" -vvvv'
    ff playback --ansible 'openstack_ceph_mon.yml --extra-vars "node=compute02" -vvvv'
    ff playback --ansible 'openstack_ceph_mon.yml --extra-vars "node=compute03" -vvvv'
    ff playback --ansible 'openstack_ceph_mon.yml --extra-vars "node=compute04" -vvvv'


### To copy the Ceph keys to nodes
Copy the configuration file and admin key to your admin node and your Ceph Nodes so that you can use the ceph CLI without having to specify the monitor address and ceph.client.admin.keyring each time you execute a command.
    
    ff playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=controller01" -vvvv'
    ff playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=controller02" -vvvv'
    ff playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute01" -vvvv'
    ff playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute02" -vvvv'
    ff playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute03" -vvvv'
    ff playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute04" -vvvv'
    ff playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute05" -vvvv'
    ff playback --ansible 'openstack_ceph_copy_keys.yml --extra-vars "node=compute06" -vvvv'


### Create the cinder ceph user and pool name
    ff playback --ansible 'openstack_ceph_cinder_pool_user.yml -vvvv'

### cinder-api
    ff playback --ansible 'openstack_cinder_api.yml --extra-vars "host=controller01" -vvvv'
    ff playback --ansible 'openstack_cinder_api.yml --extra-vars "host=controller02" -vvvv'
    
### Install cinder-volume on controller node(Ceph backend)
    ff playback --ansible 'openstack_cinder_volume_ceph.yml --extra-vars "host=controller01" -vvvv'
    ff playback --ansible 'openstack_cinder_volume_ceph.yml --extra-vars "host=controller02" -vvvv'
    
Copy the ceph.client.cinder.keyring from ceph-admin node to /etc/ceph/ceph.client.cinder.keyring of cinder volume nodes and nova-compute nodes to using the ceph client.

    ceph auth get-or-create client.cinder | ssh ubuntu@controller01 sudo tee /etc/ceph/ceph.client.cinder.keyring
    ceph auth get-or-create client.cinder | ssh ubuntu@controller02 sudo tee /etc/ceph/ceph.client.cinder.keyring
    ceph auth get-or-create client.cinder | ssh ubuntu@compute01 sudo tee /etc/ceph/ceph.client.cinder.keyring
    ceph auth get-or-create client.cinder | ssh ubuntu@compute02 sudo tee /etc/ceph/ceph.client.cinder.keyring
    ceph auth get-or-create client.cinder | ssh ubuntu@compute03 sudo tee /etc/ceph/ceph.client.cinder.keyring
    ceph auth get-or-create client.cinder | ssh ubuntu@compute04 sudo tee /etc/ceph/ceph.client.cinder.keyring
    ceph auth get-or-create client.cinder | ssh ubuntu@compute05 sudo tee /etc/ceph/ceph.client.cinder.keyring
    ceph auth get-or-create client.cinder | ssh ubuntu@compute06 sudo tee /etc/ceph/ceph.client.cinder.keyring


### Restart volume service dependency to take effect for ceph backend
    python restart_cindervol_deps.py ubuntu@controller01 ubuntu@controller02

#### Nova Controller
    ff playback --ansible 'openstack_compute_controller.yml --extra-vars "host=controller01" -vvvv'
    ff playback --ansible 'openstack_compute_controller.yml --extra-vars "host=controller02" -vvvv'

### Add Dashboard
    ff playback --ansible 'openstack_horizon.yml -vvvv'

#### Nova Computes
    ff playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute01 my_ip=192.169.151.16" -vvvv'
    ff playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute02 my_ip=192.169.151.22" -vvvv'
    ff playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute03 my_ip=192.169.151.18" -vvvv'
    ff playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute04 my_ip=192.169.151.25" -vvvv'
    ff playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute05 my_ip=192.169.151.12" -vvvv'
    ff playback --ansible 'openstack_compute_node.yml --extra-vars "host=compute06 my_ip=192.169.151.14" -vvvv'

### Install Legacy networking nova-network(FlatDHCP Only)
    ff playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute01 my_ip=192.169.151.16" -vvvv'
    ff playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute02 my_ip=192.169.151.22" -vvvv'
    ff playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute03 my_ip=192.169.151.18" -vvvv'
    ff playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute04 my_ip=192.169.151.25" -vvvv'
    ff playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute05 my_ip=192.169.151.12" -vvvv'
    ff playback --ansible 'openstack_nova_network_compute.yml --extra-vars "host=compute06 my_ip=192.169.151.14" -vvvv'

Create initial network. For example, using an exclusive slice of 172.16.0.0/16 with IP address range 172.16.0.1 to 172.16.255.254:
    
    nova network-create ext-net --bridge br100 --multi-host T --fixed-range-v4 172.16.0.0/16
    nova floating-ip-bulk-create --pool ext-net 192.169.151.65/26
    nova floating-ip-bulk-list

Extend the demo-net pool:
    
    nova floating-ip-bulk-create --pool ext-net 192.169.151.128/26
    nova floating-ip-bulk-create --pool ext-net 192.169.151.192/26
    nova floating-ip-bulk-list

### Orchestration components(heat)
     ff playback --ansible 'openstack_heat_controller.yml --extra-vars "host=controller01" -vvvv'
     ff playback --ansible 'openstack_heat_controller.yml --extra-vars "host=controller02" -vvvv'
     
### Enable service auto start
    python patch-autostart.py

### Dns as a Service
    ff playback --ansible openstack_dns.yml

execute this on controller01:
    
    bash /mnt/designate-keystone-setup
    nohup designate-central > /dev/null 2>&1 &
    nohup designate-api > /dev/null 2>&1 &
    
### Convert kvm to docker(OPTIONAL)
    ff playback --novadocker --user ubuntu --hosts compute06
    