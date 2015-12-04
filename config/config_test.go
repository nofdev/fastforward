package config

import (
	"log"
	"testing"

	"github.com/alyu/configparser"
)

func TestReadConfig(t *testing.T) {
	config, err := configparser.Read("fastforward.conf")
	if err != nil {
		t.Error(err)
	}
	log.Println(config)
}
