# Ansible module
FastForward provides a large number ansible modules for OpenStack provisioning.

## hostnetworking module
Module Args
* InternalNIC: which network card to be used for internal network. e.g. eth0
* InternalIP: the internal ip address. e.g. 10.0.0.10
* InternalMask: the internal ip netmask. e.g. 255.255.255.0
* InternalGateway: the internal ip gateway. e.g. 10.0.0.1
* InternalDNS1: the internal dns1. e.g. 10.0.0.2
* InternalDNS2: the internal dns2. e.g. 10.0.0.3
* ExternalNIC: which network card to be used for external network. e.g. eth1

To set the networking for controller01

	---
	- name: set networking
	hosts: controller01
	sudo: true
	tasks:
		- name: set networking
		  hostnetworking: InternalNIC=eth0 InternalIP=x.x.x.x InternalMask=x.x.x.x InternalGateway=x.x.x.x InternalDNS1=x.x.x.x InternalDNS2=x.x.x.x ExternalNIC=eth1

