package common

import "os/exec"
import "log"

// Packages struct is the module args
type Packages struct {
	Version string `json:"Version"`
	Changed bool   `json:"Changed"`
}

// InitPackages takes the assignment for Packages struct
func (p *Packages) InitPackages(k, v interface{}) {
	switch k {
	case "Version":
		p.Version = v.(string)
	}
}

// InstallRepo enable the OpenStack repository
func (p *Packages) InstallRepo() {
	// install software-properties-common
	cmd := exec.Command("sudo", "apt-get", "install", "software-properties-common", "-y")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("install software-properties-common error: %s", err)
	}

	// add add-apt-repository
	cmd = exec.Command("sudo", "add-apt-repository", "cloud-archive:"+p.Version)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("add-apt-repository error: %s", err)
	}

	// do the apt update for finalize the installation
	cmd = exec.Command("sudo", "apt-get", "update")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("apt update error: %s", err)
	}
	
	// do the apt dist upgrade for the finalize the installation
	cmd = exec.Command("sudo", "apt-get", "dist-upgrade", "-y")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("apt dist upgrade error: %s", err)
	}
}

// InstallClient install the openstack python client
func (p *Packages) InstallClient() {
	cmd := exec.Command("sudo", "apt-get", "install", "python-openstackclient", "-y")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("install python-openstackclient error: %s", err)
	}
}
