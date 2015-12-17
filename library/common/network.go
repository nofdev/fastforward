package common

// Interfaces takes /etc/network/interfaces config file
// Every filed are needed
type Interfaces struct {
	InternalNIC     string `json:"InternalNIC"`     // default is eth0
	InternalIP      string `json:"InternalIP"`      // the internal ip address
	InternalMask    string `json:"InternalMask"`    // the netmask for internal ip
	InternalGateway string `json:"InternalGateway"` // the gateway for internal ip
	InternalDNS1    string `json:"InternalDNS1"`    // the dns-nameservers
	InternalDNS2    string `json:"InternalDNS2"`
	ExternalNIC     string `json:"ExternalNIC"` // default is eth1
	Restart         bool   `json:"Restart"`     // restart the system
	Changed         bool   `json:"Changed"`     // changed status
}

// InitInterfaces set the Interfaces struct
func (i *Interfaces) InitInterfaces(k, v interface{}) {
	switch k {
	case "InternalNIC":
		i.InternalNIC = v.(string)
	case "InternalIP":
		i.InternalIP = v.(string)
	case "InternalMask":
		i.InternalMask = v.(string)
	case "InternalGateway":
		i.InternalGateway = v.(string)
	case "InternalDNS1":
		i.InternalDNS1 = v.(string)
	case "InternalDNS2":
		i.InternalDNS2 = v.(string)
	case "ExternalNIC":
		i.ExternalNIC = v.(string)
	case "Restart":
		i.Restart = v.(bool)
	}
}

// EtcNetworkInterface is /etc/network/interfaces file
const etcNetworkInterface = `
# This file describes the network interfaces available on your system
# and how to activate them. For more information, see interfaces(5).

# The loopback network interface
auto lo
iface lo inet loopback

# Source interfaces
# Please check /etc/network/interfaces.d before changing this file
# as interfaces may have been defined in /etc/network/interfaces.d
# NOTE: the primary ethernet device is defined in
# /etc/network/interfaces.d/eth0
# See LP: #1262951
source /etc/network/interfaces.d/*.cfg
`

// PurgeMainConf use the interfaces.d dir instead of /etc/network/interfaces
func (i *Interfaces) PurgeMainConf() {
	ParseTmpl(i, etcNetworkInterface, "interface", "/etc/network/interfaces", 0644)
}

const internalNIC = `
auto {{.InternalNIC}}
iface {{.InternalNIC}} inet static
address {{.InternalIP}}
netmask {{.InternalMask}}
gateway {{.InternalGateway}}
dns-nameservers {{.InternalDNS1}} {{.InternalDNS2}}
`

// SetInternalNIC sets the internal nic
func (i *Interfaces) SetInternalNIC() {
	ParseTmpl(i, internalNIC, "internal", "/etc/network/interfaces.d/"+i.InternalNIC+".cfg", 0644)
}

const externalNIC = `
# The public network interface
auto {{.ExternalNIC}}
iface  {{.ExternalNIC}} inet manual
up ip link set dev $IFACE up
down ip link set dev $IFACE down
`

// SetExternalNIC sets the external nic
func (i *Interfaces) SetExternalNIC() {
	ParseTmpl(i, externalNIC, "external", "/etc/network/interfaces.d/"+i.ExternalNIC+".cfg", 0644)
}
