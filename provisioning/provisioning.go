package provisioning

import "github.com/wingedpig/loom"

// Provisioning contains the provisioning method only.
type Provisioning interface {
	Execute(c Cmd) (string, error)
}

// Conf contains ssh and other configuration data needed for all the public functions in provisioning stage.
type Conf struct  {
	loom.Config
}

// MakeConfig generate the configuration for ssh login.
func MakeConfig(user, host string, output, abort bool) (*Conf, error) {
	return &Conf{loom.Config{User: user, Host: host,
		DisplayOutput: output, AbortOnError: abort}}, nil
}

// Cmd contains command line configurations.
type Cmd struct {
	// AptCache is using the apt-get update before execute the command line.
	AptCache bool

	// CmdLine contains the command line.
	CmdLine string

	// UseSudo is the sudo privilege.
	UseSudo bool

}

// Execute takes a command and runs it on the remote host over ssh.
func (c *Conf) Execute(d Cmd) (string, error) {
	if d.AptCache {
		return c.Sudo("apt-get update")
	}
	if d.UseSudo {
		return c.Sudo(d.CmdLine)
	}
	return c.Run(d.CmdLine)
}

