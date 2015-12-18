# Ansible module
FastForward provides a large number ansible modules for OpenStack provisioning.

## hostnetworking module
Module Args
* InternalNIC=[nic_name]: which network card to be used for internal network. e.g. eth0
* InternalIP=[ip]: the internal ip address. e.g. 10.0.0.10
* InternalMask=[mask]: the internal ip netmask. e.g. 255.255.255.0
* InternalGateway=[ip]: the internal ip gateway. e.g. 10.0.0.1
* InternalDNS1=[ip]: the internal dns1. e.g. 10.0.0.2
* InternalDNS2=[ip]: the internal dns2. e.g. 10.0.0.3
* ExternalNIC=[nic_name]: which network card to be used for external network. e.g. eth1
* Restart=[true] | [false]: (OPTION) restart the system to take effect

To set the networking for controller01

	---
	- name: set networking
	  hosts: controller01
	  sudo: true
	  tasks:
        - name: set networking
		  hostnetworking: InternalNIC=eth0 InternalIP=x.x.x.x InternalMask=x.x.x.x InternalGateway=x.x.x.x InternalDNS1=x.x.x.x InternalDNS2=x.x.x.x ExternalNIC=eth1 Restart=true

## ntpserver module
Module Args
* Server=[ip]: the upstream ntp for your ntp server

Install ntp server on controller01

	---
	- name: install ntp server
	  hosts: controller01
	  sudo: true
	  tasks:
	    - name: install ntp server
	      ntpserver: Server=s1a.time.edu.cn

Install ntp client on the other nodes

	---
	- name: install ntp client
	  hosts: computes
	  sudo: true
	  tasks:
	    - name: install ntp client
		  ntpserver: Server=[CONTROLLER01_IP]

## ospackages module
Module Args
* Version=[liberty]: the cloud-archive version

Install openstack packages and python clients on all nodes

	---
	- name: install packages and python clients
	  hosts: all
	  sudo: true
	  tasks:
	    - name: install packages and python clients
		  ospackages: Version=liberty
 