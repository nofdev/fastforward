// The config package is that OpenStack configuration.
// This configuration is the yaml file "../vars/openstack/openstack.yml".
package config

import (
	"github.com/nofdev/fastforward/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"text/template"
)

// The the location for configuration file of playback.
// This file must be a yaml file.
const CONFIGFILE string = "../vars/openstack/openstack.yml"

// The configuration of OpenStack
// Each var is the key(CONFIGFILE) for template
type Config struct {
	Openstack_admin_user string
	Openstack_admin_pass string
}

// Parsing yaml form CONFIGFILE.
// The CONFIGFILE is a const type for yaml location.
// Return the Config struct.
func (c *Config) Parse() Config {
	source, err := ioutil.ReadFile(CONFIGFILE)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &c)
	if err != nil {
		panic(err)
	}
	return *c
}

// Generate configuration file from a template.
// temp string is a const of configuration.
func (c *Config) GenConf(temp string, newConf string) (err error) {
	t, _ := template.New("GenConf").Parse(temp)

	clean, err := os.OpenFile(newConf, os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer clean.Close()

	target, _ := os.OpenFile(newConf, os.O_WRONLY, 0644)
	defer target.Close()

	err = t.Execute(target, c.Parse())
	if err != nil {
		panic(err)
	}
	return
}
