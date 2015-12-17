package common

import "testing"

func TestPurgeMainConf(t *testing.T) {
	i.PurgeMainConf()
}

func TestParseTmpl(t *testing.T) {
	i := &Interfaces{"eth0", "0.0.0.0", "255.255.255.255", "0.0.0.1", "0.0.0.2", "0.0.0.3", "eth1", true}
	ParseTmpl(i, EtcNetworkInterface, "interface", "test1.txt", 0644)
}