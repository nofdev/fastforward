package provisioning

import "github.com/wingedpig/loom"

// Provisioning contains the provisioning method only.
type Provisioning interface {
	Execute(c Cmd) (string, error)
	GetFile(remotefile string, localfile string) error
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
func (c *Conf) Execute(d Cmd) (result string, err error) {
	if d.AptCache {
		c.Sudo("apt-get update")
	}
	if d.UseSudo {
		result, err = c.Sudo(d.CmdLine)
	} else {
		result, err = c.Run(d.CmdLine)
	}
	return
}

// GetFile copies the file from the remote host to the local FastForward server, using scp. Wildcards are not currently supported. 
func (c *Conf) GetFile(remotefile string, localfile string) error {
	return c.Get(remotefile, localfile)
}