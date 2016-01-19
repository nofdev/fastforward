#!/usr/bin/env bash

# MySQL HA
# Configration:
# mysql:
#   vip: '192.168.77.3'
#   root-password: changeme
#   sst-password: changeme
# mysql-hacluster:
#   corosync_transport: unicast
juju add-unit mysql --to lxc:7
juju deploy hacluster mysql-hacluster
juju add-relation mysql mysql-hacluster


# RabbitMQ HA
# Configration
# rabbitmq-server:
#   vip: '192.168.77.4'
#   vip_cidr: 24
juju add-unit rabbitmq-server --to lxc:7

# Keystone
# Configration:
# keystone:
#   admin-user: 'admin'
#   admin-password: 'openstack'
#   admin-token: 'changeme'
#   vip: '192.168.77.5'
# keystone-hacluster:
#   corosync_transport: unicast
juju add-unit keystone --to lxc:5
juju deploy hacluster keystone-hacluster
juju add-relation keystone keystone-hacluster

# Cloud Controller
# Configration
# nova-cloud-controller:
#   vip: '192.168.77.6'
#   network-manager: 'Neutron'
#   quantum-security-groups: yes
# ncc-hacluster:
#   corosync_transport: unicast
juju add-unit nova-cloud-controller --to lxc:5
juju deploy hacluster ncc-hacluster
juju add-relation nova-cloud-controller ncc-hacluster

# Glance
# Configration
# glance:
#   vip: '192.168.77.7'
# glance-hacluster:
#   corosync_transport: unicast
juju add-unit glance --to lxc:5
juju deploy hacluster glance-hacluster
juju add-relation glance glance-hacluster


# Cinder
# Configration
# cinder:
#   block-device: 'None'
#   vip: '192.168.77.8'
# cinder-hacluster:
#   corosync_transport: unicast
juju add-unit cinder --to lxc:5
juju deploy hacluster cinder-hacluster
juju add-relation cinder cinder-hacluster
juju add-relation cinder ceph


# Neutron
# Configration
# neutron-gateway:
#   ext-port: 'eth1'
juju add-unit neutron-gateway --to lxc:5
juju add-relation neutron-gateway mysql
juju add-relation neutron-gateway:amqp-nova rabbitmq-server:amqp


# Swift
# swift-proxy:
#   zone-assignment: 'manual'
#   replicas: 3
#   swift-hash: 'fdfef9d4-8b06-11e2-8ac0-531c923c8fae'
#   vip: '192.168.77.9'
# swift-hacluster:
#   corosync_transport: unicast
# swift-storage-z1:
#   zone: 1
#   block-device: 'vdb'
# swift-storage-z2:
#   zone: 2
#   block-device: 'vdb'
# swift-storage-z3:
#   zone: 3
#   block-device: 'vdb'
juju deploy --config ha_config.yaml swift-proxy --to lxc:0
juju add-unit swift-proxy --to lxc:1
juju add-unit swift-proxy --to lxc:2
juju deploy --config ha_config.yaml hacluster swift-hacluster
juju deploy --config ha_config.yaml swift-storage swift-storage-z1 --to lxc:0
juju deploy --config ha_config.yaml swift-storage swift-storage-z2 --to lxc:1
juju deploy --config ha_config.yaml swift-storage swift-storage-z3 --to lxc:2
juju add-relation swift-proxy swift-hacluster
juju add-relation swift-proxy keystone
juju add-relation swift-proxy swift-storage-z1
juju add-relation swift-proxy swift-storage-z2
juju add-relation swift-proxy swift-storage-z3

# Horizon
# Configration:
# openstack-dashboard:
#   vip: '192.168.77.10'
juju add-unit openstack-dashboard --to lxc:5
juju deploy hacluster dashboard-hacluster
juju add-relation openstack-dashboard dashboard-hacluster
