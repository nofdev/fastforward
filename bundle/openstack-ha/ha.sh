#!/usr/bin/env bash
# juju version 1.25.2
# maas version 1.9.0

# ceph mon and osd
juju deploy --config ha_config.yaml ceph --to 1
juju add-unit ceph --to 2
juju add-unit ceph --to 3
juju deploy --config ha_config.yaml ceph-osd --to 1
juju add-unit ceph-osd --to 2
juju add-unit ceph-osd --to 3
juju add-relation ceph ceph-osd

# mysql
juju deploy --config ha_config.yaml percona-cluster mysql --to lxc:1
juju add-unit mysql --to lxc:2
juju add-unit mysql --to lxc:3
juju deploy --config ha_config.yaml hacluster mysql-hacluster
juju add-relation mysql mysql-hacluster

# rabbitmq
juju deploy rabbitmq-server --to lxc:1
juju add-unit rabbitmq-server --to lxc:2
juju add-unit rabbitmq-server --to lxc:3

# keystone
juju deploy --config ha_config.yaml keystone --to lxc:1
juju add-unit keystone --to lxc:2
juju add-unit keystone --to lxc:3
juju deploy --config ha_config.yaml hacluster keystone-hacluster
juju add-relation keystone keystone-hacluster
juju add-relation keystone mysql

# cloud controller
juju deploy --config ha_config.yaml nova-cloud-controller --to lxc:1
juju add-unit nova-cloud-controller --to lxc:2
juju add-unit nova-cloud-controller --to lxc:3
juju deploy --config ha_config.yaml hacluster ncc-hacluster
juju add-relation nova-cloud-controller ncc-hacluster
juju add-relation nova-cloud-controller mysql
juju add-relation nova-cloud-controller keystone
juju add-relation nova-cloud-controller rabbitmq-server

# glance
juju deploy --config ha_config.yaml glance --to lxc:1
juju add-unit glance --to lxc:2
juju add-unit glance --to lxc:3
juju deploy --config ha_config.yaml hacluster glance-hacluster
juju add-relation glance glance-hacluster
juju add-relation glance mysql
juju add-relation glance nova-cloud-controller
juju add-relation glance ceph
juju add-relation glance keystone

# cinder
juju deploy --config ha_config.yaml cinder --to lxc:1
juju add-unit cinder --to lxc:2
juju add-unit cinder --to lxc:3
juju deploy --config ha_config.yaml hacluster cinder-hacluster
juju add-relation cinder cinder-hacluster
juju add-relation cinder mysql
juju add-relation cinder keystone
juju add-relation cinder nova-cloud-controller
juju add-relation cinder rabbitmq-server
juju add-relation cinder ceph
juju add-relation cinder glance

# neutron
juju deploy --config ha_config.yaml neutron-gateway --to 1
juju add-unit neutron-gateway --to 2
juju add-relation neutron-gateway mysql
juju add-relation neutron-gateway:amqp rabbitmq-server:amqp
juju add-relation neutron-gateway:amqp-nova rabbitmq-server:amqp
juju add-relation neutron-gateway nova-cloud-controller

# nova
juju deploy --config ha_config.yaml nova-compute --to 4
juju add-unit nova-compute --to 5
juju add-unit nova-compute --to 6
juju add-unit nova-compute --to 7
juju add-unit nova-compute --to 8
juju add-unit nova-compute --to 9
juju add-unit nova-compute --to 10
juju add-relation nova-compute nova-cloud-controller
juju add-relation nova-compute rabbitmq-server
juju add-relation nova-compute glance
juju add-relation nova-compute ceph

# swift
juju deploy --config ha_config.yaml swift-proxy --to 4
juju add-unit swift-proxy --to 5
juju add-unit swift-proxy --to 6
juju deploy --config ha_config.yaml hacluster swift-hacluster
juju deploy --config ha_config.yaml swift-storage swift-storage-z1 --to 4
juju deploy --config ha_config.yaml swift-storage swift-storage-z2 --to 5
juju deploy --config ha_config.yaml swift-storage swift-storage-z3 --to 6
juju add-relation swift-proxy swift-hacluster
juju add-relation swift-proxy keystone
juju add-relation swift-proxy swift-storage-z1
juju add-relation swift-proxy swift-storage-z2
juju add-relation swift-proxy swift-storage-z3

# horizon
juju deploy --config ha_config.yaml openstack-dashboard --to lxc:1
juju add-unit openstack-dashboard --to lxc:2
juju add-unit openstack-dashboard --to lxc:3
juju deploy hacluster dashboard-hacluster
juju add-relation openstack-dashboard dashboard-hacluster
juju add-relation openstack-dashboard keystone
