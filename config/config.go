package config

import (
	"log"

	"github.com/alyu/configparser"
)

// Configure takes the FastForward Configuration.
type Configure interface {
	LoadConf() (conf *configparser.Configuration)
}

// Conf is the configuration data structure.
type Conf struct {
	DEFAULT  map[string]string
	PLAYBACK map[string]string
}

// LoadConf loads the FastForward configuration and return the Conf pointer.
func (c *Conf) LoadConf() *Conf {
	path := "fastforward.conf"
	conf, err := configparser.Read(path)
	if err != nil {
		log.Fatal(err)
	}
	DefaultSection, err := conf.Section("DEFAULT")
	checkErr(err)

	PlaybackSection, err := conf.Section("PLAYBACK")
	checkErr(err)

	FFconf := &Conf{
		DEFAULT: map[string]string{"provisioning_driver": DefaultSection.Options()["provisioning_driver"],
			"orchestration_driver": DefaultSection.Options()["orchestration_driver"],
			"monitoring_driver":    DefaultSection.Options()["monitoring_driver"]},
		PLAYBACK: map[string]string{"use_ansible": PlaybackSection.Options()["use_ansible"],
			"ansible_cfg": PlaybackSection.Options()["ansible_cfg"],
			"private_key": PlaybackSection.Options()["private_key"]},
	}
	return FFconf
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
