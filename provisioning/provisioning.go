package provisioning

import "github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/wingedpig/loom"

// Provisioning contains the provisioning method only.
type Provisioning interface {
	Execute(c Cmd) (string, error)
	GetFile(remotefile string, localfile string) error
	Self(d Cmd) (result string, err error)
	PutFile(localfiles string, remotefile string) error
	PutString(data string, remotefile string) error
}

// Conf contains ssh and other configuration data needed for all the public functions in provisioning stage.
type Conf struct {
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

// PutFile copies one or more local files to the remote host, using scp. localfiles can contain wildcards, and remotefile can be either a directory or a file.
func (c *Conf) PutFile(localfiles string, remotefile string) error {
	return c.Put(localfiles, remotefile)
}

// Self executes a command on the FastForward API server.
func (c *Conf) Self(d Cmd) (result string, err error) {
	result, err = c.Local(d.CmdLine)
	return
}

// PutString generates a new file on the remote host containing data. The file is created with mode 0644.
func (c *Conf) PutString(data string, remotefile string) error {
	return c.Config.PutString(data, remotefile)
}
