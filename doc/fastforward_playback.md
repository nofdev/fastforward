# FastForward Playback
Playback is a core components of FastForward that is an OpenStack provisioning DevOps tool and all of the OpenStack components can be deployed automation with high availability on Ubuntu based operating system.

#### Requirements
* The OpenStack bare metal hosts are in MAAS environment(recommend)
* All hosts are two NICs at least(external and internal)
* We assume that you have ceph installed, the cinder bachend default using ceph, the running instace default to using ceph as it's local storage. About ceph please visit: http://docs.ceph.com/docs/master/rbd/rbd-openstack/ or see the (Option)Ceph Guide below.
* nova user can be login to each compute node via ssh passwordless
* The FastForward Playback node is the same as ceph-deploy node where can be login to each openstack node passwordless
* The FastForward Playback node default using ~/.ssh/id_rsa ssh private key to logon remote server

#### Install FastForward Playback
Install FastForward Playback on FASTFORWARD-NODE

    pip install fastforward

#### Prepare environment
Prepare the OpenStack environment.
(NOTE) DO NOT setup eth1 in /etc/network/interfaces

    ff-openstack-env --user ubuntu --hosts CONTROLLER1,CONTROLLER2,COMPUTE1,COMPUTE2,OBJECT1,OBJECT2,BLOCK1,BLOCK2,HAPROXY1,HAPROXY2 prepare-host

#### MySQL HA
Deploy to CONTROLLER1

    ff-openstack-mysql --user ubuntu --hosts CONTROLLER1 install 
    ff-openstack-mysql --user ubuntu --hosts CONTROLLER1 config  --wsrep-cluster-address "gcomm://CONTROLLER1,CONTROLLER2" --wsrep-node-name="galera1" --wsrep-node-address="CONTROLLER1"

Deploy to CONTROLLER2

    ff-openstack-mysql --user ubuntu --hosts CONTROLLER2 install 
    ff-openstack-mysql --user ubuntu --hosts CONTROLLER2 config  --wsrep-cluster-address "gcomm://CONTROLLER1,CONTROLLER2" --wsrep-node-name="galera2" --wsrep-node-address="CONTROLLER2"

Start cluster

    ff-openstack-mysql --user ubuntu --hosts CONTROLLER1 manage --wsrep-new-cluster
    ff-openstack-mysql --user ubuntu --hosts CONTROLLER2 manage --start
    ff-openstack-mysql --user ubuntu --hosts CONTROLLER1 manage --change-root-password changeme

#### HAProxy HA
Deploy to HAPROXY1

    ff-openstack-haproxy --user ubuntu --hosts HAPROXY1 install

Deploy to HAPROXY2

    ff-openstack-haproxy --user ubuntu --hosts HAPROXY2 install

Generate the HAProxy configuration and upload to target hosts(Do not forget to edit the generated configuration)

    ff-openstack-haproxy gen-conf 
    ff-openstack-haproxy --user ubuntu --hosts HAPROXY1,HAPROXY2 config --upload-conf haproxy.cfg 

Configure Keepalived

    ff-openstack-haproxy --user ubuntu --hosts HAPROXY1 config --configure-keepalived --router_id lb1 --priority 150 --state MASTER --interface eth0 --vip CONTROLLER_VIP 
    ff-openstack-haproxy --user ubuntu --hosts HAPROXY2 config --configure-keepalived --router_id lb2 --priority 100 --state SLAVE --interface eth0 --vip CONTROLLER_VIP 

#### RabbitMQ HA
Deploy to CONTROLLER1 and CONTROLLER2

    ff-openstack-rabbitmq --user ubuntu --hosts CONTROLLER1,CONTROLLER2 install --erlang-cookie changemechangeme --rabbit-user openstack --rabbit-pass changeme
    
Create cluster

    ff-openstack-rabbitmq --user ubuntu --hosts CONTROLLER2 join-cluster --name rabbit@CONTROLLER1

#### Keystone HA
Create keystone database

    ff-openstack-keystone --user ubuntu --hosts CONTROLLER1 create-keystone-db --root-db-pass changeme --keystone-db-pass changeme

Install keystone on CONTROLLER1 and CONTROLLER2

    ff-openstack-keystone --user ubuntu --hosts CONTROLLER1 install --admin_token changeme --connection mysql+pymysql://keystone:changeme@CONTROLLER_VIP/keystone --memcache_servers CONTROLLER1:11211,CONTROLLER2:11211 --populate
    ff-openstack-keystone --user ubuntu --hosts CONTROLLER2 install --admin_token changeme --connection mysql+pymysql://keystone:changeme@CONTROLLER_VIP/keystone --memcache_servers CONTROLLER1:11211,CONTROLLER2:11211

Create the service entity and API endpoints

    ff-openstack-keystone --user ubuntu --hosts CONTROLLER1 create-entity-and-endpoint --os-token changeme --os-url http://CONTROLLER_VIP:35357/v3 --public-endpoint http://CONTROLLER_VIP:5000/v2.0 --internal-endpoint http://CONTROLLER_VIP:5000/v2.0 --admin-endpoint http://CONTROLLER_vip:35357/v2.0

Create projects, users, and roles

    ff-openstack-keystone --user ubuntu --hosts CONTROLLER1 create-projects-users-roles --os-token changeme --os-url http://CONTROLLER_VIP:35357/v3 --admin-pass changeme --demo-pass changeme

(OPTION) you will need to create OpenStack client environment scripts
admin-openrc.sh

    export OS_PROJECT_DOMAIN_ID=default
    export OS_USER_DOMAIN_ID=default
    export OS_PROJECT_NAME=admin
    export OS_TENANT_NAME=admin
    export OS_USERNAME=admin
    export OS_PASSWORD=changeme
    export OS_AUTH_URL=http://CONTROLLER_VIP:35357/v3
    export OS_IDENTITY_API_VERSION=3
    export OS_IMAGE_API_VERSION=2
    export OS_AUTH_VERSION=3

demo-openrc.sh

    export OS_PROJECT_DOMAIN_ID=default
    export OS_USER_DOMAIN_ID=default
    export OS_PROJECT_NAME=demo
    export OS_TENANT_NAME=demo
    export OS_USERNAME=demo
    export OS_PASSWORD=changeme
    export OS_AUTH_URL=http://CONTROLLER_VIP:5000/v3
    export OS_IDENTITY_API_VERSION=3
    export OS_IMAGE_API_VERSION=2
    export OS_AUTH_VERSION=3

#### Glance HA
Create glance database

    ff-openstack-glance --user ubuntu --hosts CONTROLLER1 create-glance-db --root-db-pass changeme --glance-db-pass changeme

Create service credentials

    ff-openstack-glance --user ubuntu --hosts CONTROLLER1 create-service-credentials --os-password changeme --os-auth-url http://CONTROLLER_VIP:35357/v3 --glance-pass changeme --endpoint http://CONTROLLER_VIP:9292

Install glance on CONTROLLER1 and CONTROLLER2

    ff-openstack-glance --user ubuntu --hosts CONTROLLER1 install --connection mysql+pymysql://glance:GLANCE_PASS@CONTROLLER_VIP/glance --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --glance-pass changeme  --swift-store-auth-address http://CONTROLLER_VIP:5000/v2.0/ --populate
    ff-openstack-glance --user ubuntu --hosts CONTROLLER2 install --connection mysql+pymysql://glance:GLANCE_PASS@CONTROLLER_VIP/glance --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --glance-pass changeme  --swift-store-auth-address http://CONTROLLER_VIP:5000/v2.0/ 


#### Nova HA
Create nova database

    ff-openstack-nova --user ubuntu --hosts CONTROLLER1 create-nova-db --root-db-pass changeme --nova-db-pass changeme 

Create service credentials

    ff-openstack-nova --user ubuntu --hosts CONTROLLER1 create-service-credentials --os-password changeme --os-auth-url http://CONTROLLER_VIP:35357/v3 --nova-pass changeme --endpoint 'http://CONTROLLER_VIP:8774/v2/%\(tenant_id\)s'

Install nova on CONTROLLER1

    ff-openstack-nova --user ubuntu --hosts CONTROLLER1 install --connection mysql+pymysql://nova:NOVA_PASS@CONTROLLER_VIP/nova --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --nova-pass changeme --my-ip MANAGEMENT_IP --memcached-servers CONTROLLER1:11211,CONTROLLER2:11211 --rabbit-hosts CONTROLLER1,CONTROLLER2 --rabbit-pass changeme --glance-host CONTROLLER_VIP --neutron-endpoint http://CONTROLLER_VIP:9696 --neutron-pass changeme --metadata-proxy-shared-secret changeme --populate

Install nova on CONTROLLER2

    ff-openstack-nova --user ubuntu --hosts CONTROLLER2 install --connection mysql+pymysql://nova:NOVA_PASS@CONTROLLER_VIP/nova --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --nova-pass changeme --my-ip MANAGEMENT_IP --memcached-servers CONTROLLER1:11211,CONTROLLER2:11211 --rabbit-hosts CONTROLLER1,CONTROLLER2 --rabbit-pass changeme --glance-host CONTROLLER_VIP --neutron-endpoint http://CONTROLLER_VIP:9696 --neutron-pass changeme --metadata-proxy-shared-secret changeme

#### Nova Compute
Add nova computes

    ff-openstack-nova-compute --user ubuntu --hosts COMPUTE1 install --my-ip MANAGEMENT_IP --rabbit-hosts CONTROLLER1,CONTROLLER2 --rabbit-pass changeme --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --nova-pass changeme --novncproxy-base-url http://CONTROLLER_VIP:6080/vnc_auto.html --glance-host CONTROLLER_VIP --neutron-endpoint http://CONTROLLER_VIP:9696 --neutron-pass changeme --rbd-secret-uuid changeme-changeme-changeme-changeme
    ff-openstack-nova-compute --user ubuntu --hosts COMPUTE2 install --my-ip MANAGEMENT_IP --rabbit-hosts CONTROLLER1,CONTROLLER2 --rabbit-pass changeme --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --nova-pass changeme --novncproxy-base-url http://CONTROLLER_VIP:6080/vnc_auto.html --glance-host CONTROLLER_VIP --neutron-endpoint http://CONTROLLER_VIP:9696 --neutron-pass changeme --rbd-secret-uuid changeme-changeme-changeme-changeme

The libvirt defaults to using ceph as shared storage, the ceph pool for running instance is vms. if you do not using ceph as it's bachend, you must remove the following param:
    
    images_type = rbd
    images_rbd_pool = vms
    images_rbd_ceph_conf = /etc/ceph/ceph.conf
    rbd_user = cinder
    rbd_secret_uuid = changeme-changeme-changeme-changeme
    disk_cachemodes="network=writeback"
    live_migration_flag="VIR_MIGRATE_UNDEFINE_SOURCE,VIR_MIGRATE_PEER2PEER,VIR_MIGRATE_LIVE,VIR_MIGRATE_PERSIST_DEST,VIR_MIGRATE_TUNNELLED"


#### Neutron HA
Create nova database

    ff-openstack-neutron --user ubuntu --hosts CONTROLLER1 create-neutron-db --root-db-pass changeme --neutron-db-pass changeme 

Create service credentials

    ff-openstack-neutron --user ubuntu --hosts CONTROLLER1 create-service-credentials --os-password changeme --os-auth-url http://CONTROLLER_VIP:35357/v3 --neutron-pass changeme --endpoint http://CONTROLLER_VIP:9696

Install Neutron for self-service

    ff-openstack-neutron --user ubuntu --hosts CONTROLLER1 install --connection mysql+pymysql://neutron:NEUTRON_PASS@CONTROLLER_VIP/neutron --rabbit-hosts CONTROLLER1,CONTROLLER2 --rabbit-pass changeme --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --neutron-pass changeme --nova-url http://CONTROLLER_VIP:8774/v2 --nova-pass changeme --public-interface eth1 --local-ip MANAGEMENT_INTERFACE_IP --nova-metadata-ip CONTROLLER_VIP --metadata-proxy-shared-secret changeme-changeme-changeme-changeme --populate
    ff-openstack-neutron --user ubuntu --hosts CONTROLLER2 install --connection mysql+pymysql://neutron:NEUTRON_PASS@CONTROLLER_VIP/neutron --rabbit-hosts CONTROLLER1,CONTROLLER2 --rabbit-pass changeme --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --neutron-pass changeme --nova-url http://CONTROLLER_VIP:8774/v2 --nova-pass changeme --public-interface eth1 --local-ip MANAGEMENT_INTERFACE_IP --nova-metadata-ip CONTROLLER_VIP --metadata-proxy-shared-secret changeme-changeme-changeme-changeme 


#### Neutron Agent
Install neutron agent on compute nodes
    
    ff-openstack-neutron-agent --user ubuntu --hosts COMPUTE1 install --rabbit-hosts CONTROLLER1,CONTROLLER2 --rabbit-pass changeme --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --neutron-pass changeme --public-interface eth1 --local-ip MANAGEMENT_INTERFACE_IP 
    ff-openstack-neutron-agent --user ubuntu --hosts COMPUTE2 install --rabbit-hosts CONTROLLER1,CONTROLLER2 --rabbit-pass changeme --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --neutron-pass changeme --public-interface eth1 --local-ip MANAGEMENT_INTERFACE_IP 


#### Horizon HA
Install horizon on controller nodes

    ff-openstack-horizon --user ubuntu --hosts CONTROLLER1,CONTROLLER2 install --openstack-host CONTROLLER_VIP  --memcache CONTROLLER1:11211 --time-zone Asia/Shanghai 


#### Cinder HA
Create cinder database

    ff-openstack-cinder --user ubuntu --hosts CONTROLLER1 create-cinder-db --root-db-pass changeme --cinder-db-pass changeme 

Create cinder service creadentials

    ff-openstack-cinder --user ubuntu --hosts CONTROLLER1 create-service-credentials --os-password changeme --os-auth-url http://CONTROLLER_VIP:35357/v3 --cinder-pass changeme --endpoint-v1 'http://CONTROLLER_VIP:8776/v1/%\(tenant_id\)s' --endpoint-v2 'http://CONTROLLER_VIP:8776/v2/%\(tenant_id\)s'

Install cinder-api and cinder-volume on controller nodes, the volume backend defaults to ceph (you must have ceph installed)
    
    ff-openstack-cinder --user ubuntu --hosts CONTROLLER1 install --connection mysql+pymysql://cinder:CINDER_PASS@CONTROLLER_VIP/cinder --rabbit-pass changeme --rabbit-hosts CONTROLLER1,CONTROLLER2 --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --cinder-pass changeme --my-ip MANAGEMENT_INTERFACE_IP --glance-host CONTROLLER_VIP --rbd-secret-uuid changeme-changeme-changeme-changeme --populate
    ff-openstack-cinder --user ubuntu --hosts CONTROLLER2 install --connection mysql+pymysql://cinder:CINDER_PASS@CONTROLLER_VIP/cinder --rabbit-pass changeme --rabbit-hosts CONTROLLER1,CONTROLLER2 --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --cinder-pass changeme --my-ip MANAGEMENT_INTERFACE_IP --glance-host CONTROLLER_VIP --rbd-secret-uuid changeme-changeme-changeme-changeme

#### Swift proxy HA
Create the Identity service credentials

    ff-openstack-swift --user ubuntu --hosts CONTROLLER1 create-service-credentials --os-password changeme --os-auth-url http://CONTROLLER_VIP:35357/v3 --swift-pass changeme --public-internal-endpoint 'http://CONTROLLER_VIP:8080/v1/AUTH_%\(tenant_id\)s' --admin-endpoint http://CONTROLLER_VIP:8080/v1 

Install swift proxy

    ff-openstack-swift --user ubuntu --hosts CONTROLLER1,CONTROLLER2 install --auth-uri http://CONTROLLER_VIP:5000 --auth-url http://CONTROLLER_VIP:35357 --swift-pass changeme --memcache-servers CONTROLLER1:11211,CONTROLLER2:11211 


#### Swift storage
Prepare disks on storage node

    ff-openstack-swift-storage --user ubuntu --hosts OBJECT1,OBJECT2 prepare-disks --name sdb,sdc,sdd,sde

Install swift storage on storage node

    ff-openstack-swift-storage --user ubuntu --hosts OBJECT1 install --address MANAGEMENT_INTERFACE_IP --bind-ip MANAGEMENT_INTERFACE_IP 
    ff-openstack-swift-storage --user ubuntu --hosts OBJECT2 install --address MANAGEMENT_INTERFACE_IP --bind-ip MANAGEMENT_INTERFACE_IP 

Create account ring on controller node

    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 create-account-builder-file --partitions 10 --replicas 3 --moving 1 
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 account-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sdb --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 account-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sdc --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 account-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sdd --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 account-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sde --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 account-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sdb --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 account-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sdc --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 account-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sdd --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 account-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sde --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 account-builder-rebalance
Create container ring on controller node
    
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 create-container-builder-file --partitions 10 --replicas 3 --moving 1 
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 container-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sdb --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 container-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sdc --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 container-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sdd --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 container-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sde --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 container-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sdb --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 container-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sdc --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 container-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sdd --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 container-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sde --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 container-builder-rebalance

Create object ring on controller node
    
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 create-object-builder-file --partitions 10 --replicas 3 --moving 1 
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 object-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sdb --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 object-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sdc --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 object-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sdd --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 object-builder-add --region 1 --zone 1 --ip OBJECT1_MANAGEMENT_IP --device sde --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 object-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sdb --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 object-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sdc --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 object-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sdd --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 object-builder-add --region 1 --zone 1 --ip OBJECT2_MANAGEMENT_IP --device sde --weight 100
    ff-openstack-swift-storage --user ubuntu --hosts CONTROLLER1 object-builder-rebalance

 Sync the builder file from controller node to each storage node and other any proxy node

    ff-openstack-swift-storage --user ubuntu --host CONTROLLER1 sync-builder-file --to CONTROLLER2,OBJECT1,OBJECT2

Finalize installation on all nodes

    ff-openstack-swift --user ubuntu --hosts CONTROLLER1,CONTROLLER2,OBJECT1,OBJECT2 finalize-install --swift-hash-path-suffix changeme --swift-hash-path-prefix changeme



TODO:
    nova ssh keys
    esxi backend
    
####(Option) Ceph Guide
Create ceph cluster directory

    mkdir ceph-cluster
    cd ceph-cluster

Create cluster and add initial monitor(s) to the ceph.conf

    ff-openstack-ceph-deploy new  CONTROLLER1 CONTROLLER2 COMPUTE1 COMPUTE2 BLOCK1 BLOCK2
    echo "osd pool default size = 2" | tee -a ceph.conf

Install ceph client

    ff-openstack-ceph-deploy install ff-openstack-NODE CONTROLLER1 CONTROLLER2 COMPUTE1 COMPUTE2 BLOCK1 BLOCK2

Add the initial monitor(s) and gather the keys

    ff-openstack-ceph-deploy mon create-initial

If you want to add additional monitors, do that
    
    ff-openstack-ceph-deploy mon add {additional-monitor}

Add ceph osd(s)

    ff-openstack-ceph-deploy osd create --zap-disk --fs-type ext4 BLOCK1:/dev/sdb
    ff-openstack-ceph-deploy osd create --zap-disk --fs-type ext4 BLOCK1:/dev/sdc
    ff-openstack-ceph-deploy osd create --zap-disk --fs-type ext4 BLOCK2:/dev/sdb
    ff-openstack-ceph-deploy osd create --zap-disk --fs-type ext4 BLOCK2:/dev/sdc

Sync admin key

    ff-openstack-ceph-deploy admin ff-openstack-NODE CONTROLLER1 CONTROLLER2 COMPUTE1 COMPUTE2 BLOCK1 BLOCK2
    ssh {ceph-client-node} sudo chmod +r /etc/ceph/ceph.client.admin.keyring

Create osd pool for cinder and running instance

    ceph osd pool create volumes 512
    ceph osd pool create vms 512

Setup ceph client authentication

    ceph auth get-or-create client.cinder mon 'allow r' osd 'allow class-read object_prefix rbd_children, allow rwx pool=volumes, allow rwx pool=vms'

Add the keyrings for `client.cinder` to appropriate nodes and change their ownership

    ceph auth get-or-create client.cinder | ssh {CINDER-VOLUME-NODE} sudo tee /etc/ceph/ceph.client.cinder.keyring
    ssh {CINDER-VOLUME-NODE} sudo chown cinder:cinder /etc/ceph/ceph.client.cinder.keyring

Nodes running `nova-compute` need the keyring file for the `nova-compute` process

    ceph auth get-or-create client.cinder | ssh {COMPUTE-NODE} sudo tee /etc/ceph/ceph.client.cinder.keyring

They also need to store the secret key of the `client.cinder user` in `libvirt`. The libvirt process needs it to access the cluster while attaching a block device from Cinder.
Create a temporary copy of the secret key on the nodes running `nova-compute`

    ceph auth get-key client.cinder | ssh {COMPUTE-NODE} tee client.cinder.key

Then, on the `compute nodes`, add the secret key to `libvirt` and remove the temporary copy of the key(the uuid is the same as your --rbd-secret-uuid option, you have to save the uuid for later)
    
    uuidgen
    457eb676-33da-42ec-9a8c-9293d545c337

    cat > secret.xml <<EOF
    <secret ephemeral='no' private='no'>
      <uuid>457eb676-33da-42ec-9a8c-9293d545c337</uuid>
      <usage type='ceph'>
        <name>client.cinder secret</name>
      </usage>
    </secret>
    EOF
    sudo virsh secret-define --file secret.xml
    Secret 457eb676-33da-42ec-9a8c-9293d545c337 created
    sudo virsh secret-set-value --secret 457eb676-33da-42ec-9a8c-9293d545c337 --base64 $(cat client.cinder.key) && rm client.cinder.key secret.xml

Now on every compute nodes edit your Ceph configuration file, add the client section

    [client]
    rbd cache = true
    rbd cache writethrough until flush = true
    rbd concurrent management ops = 20


If you want to remove osd 

    ssh {OSD-NODE} sudo stop ceph-mon-all && sudo stop ceph-osd-all
    ceph osd out {OSD-NUM}
    ceph osd crush remove osd.{OSD-NUM} 
    ceph auth del osd.{OSD-NUM} 
    ceph osd rm {OSD-NUM} 
    ceph osd crush remove {HOST} 

If you want to remove monitor

    ceph mon remove {MON-ID}


#### Library Use
```
from playback.api import *
admin_token = 'changeme'
connection = 'mysql+pymysql://keystone:changeme@CONTROLLER_VIP/keystone'
memcache_servers = 'CONTROLLER1:11211,CONTROLLER2:11211'

keystone = Keystone(user='ubuntu', hosts='controller1,controller2')
execute(keystone._install_keystone, admin_token, connection, memcache_servers)
```
