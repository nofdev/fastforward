package common

import "testing"

func TestPurgeMainConf(t *testing.T) {
	i := &Interfaces{"eth0", "0.0.0.0", "255.255.255.255", "0.0.0.1", "eth1", true}
	i.PurgeMainConf()
}

func TestParseTmpl(t *testing.T) {
	i := &Interfaces{"eth0", "0.0.0.0", "255.255.255.255", "0.0.0.1", "eth1", true}
	ParseTmpl(i, EtcNetworkInterface, "interface", "test1.txt", 0644)
}