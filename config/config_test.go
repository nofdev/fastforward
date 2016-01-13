package config

import (
	"log"
	"testing"

	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/alyu/configparser"
	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	config, err := configparser.Read("fastforward.conf")
	if err != nil {
		t.Error(err)
	}
	log.Printf("full configuration:\n%s", config)
}

func TestSection(t *testing.T) {
	config, err := configparser.Read("fastforward.conf")
	if err != nil {
		t.Error(err)
	}
	section, err := config.Section("DEFAULT")
	if err != nil {
		t.Error(err)
	}
	log.Printf("the default section:\n%s", section)

	section, err = config.Section("PLAYBACK")
	if err != nil {
		t.Error(err)
	}
	log.Printf("the playback section:\n%s", section)

}

func TestOption(t *testing.T) {
	config, err := configparser.Read("fastforward.conf")
	if err != nil {
		t.Error(err)
	}
	section, err := config.Section("DEFAULT")
	if err != nil {
		t.Error(err)
	}
	options := section.Options()
	log.Printf("option names:\n%s", options["provisioning_driver"])
	assert.Equal(t, "playback", options["provisioning_driver"])
}

func TestLoadConf(t *testing.T) {
	var v Conf
	conf := v.LoadConf()
	log.Printf("LoadConf:\n%s", conf)
	assert.Equal(t, "playback", conf.DEFAULT["provisioning_driver"])
}
