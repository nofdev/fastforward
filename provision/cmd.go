package provision

import "github.com/wingedpig/loom"

// The provisioning interface
type Provisioning interface {
	Execute(c Cmd) (string, error)
}

// Config contains ssh and other configuration data needed for all the public functions in playback
type Conf struct  {
	loom.Config
}

// To make the config for ssh login for instance
func MakeConfig(user, host string, output, abort bool) (*Conf, error) {
	return &Conf{loom.Config{User: user, Host: host,
		DisplayOutput: output, AbortOnError: abort}}, nil
}

// Execute command line
type Cmd struct {
	// Using apt-get update if set to true
	AptCache bool

	// Command line to execute
	CmdLine string

	// Using sudo to execute command line
	UseSudo bool

}

// Execute the command
func (c *Conf) Execute(d Cmd) (string, error) {
	if d.AptCache {
		c.Sudo("apt-get update")
	}
	if d.UseSudo {
		c.Sudo(d.CmdLine)
	} else {
		c.Run(d.CmdLine)
	}
	return "", nil
}