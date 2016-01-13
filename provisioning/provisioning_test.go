package provisioning

import "testing"
import "github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/wingedpig/loom"

type Config struct {
	loom.Config
}

// TODO: Refactor testing. the host: "TESTSERVER"" is unreachable. Could not used for Travis-CI.
func TestRun(t *testing.T) {
	c := &Config{loom.Config{User: "ubuntu", Host: "TESTSERVER",
		DisplayOutput: true, AbortOnError: true}}
	_, err := c.Run("ls -la")
	if err != nil {
		t.Errorf("Run cmd error, %s", err)
	}
}

func TestSudo(t *testing.T) {
	c := &Config{loom.Config{User: "ubuntu", Host: "TESTSERVER",
		DisplayOutput: true, AbortOnError: true}}
	_, err := c.Sudo("ls -la")
	if err != nil {
		t.Errorf("Run sudo error, %s", err)
	}
}

func TestPutString(t *testing.T) {
	c := &Config{loom.Config{User: "ubuntu", Host: "TESTSERVER",
		DisplayOutput: true, AbortOnError: true}}
	err := c.PutString("TestPutString\nTestPutString", "~/testputstring")
	if err != nil {
		t.Errorf("Run putstring error, %s", err)
	} else {
		c.Run("cat ~/testputstring")
	}
}

func TestPut(t *testing.T) {
	c := &Config{loom.Config{User: "ubuntu", Host: "TESTSERVER",
		DisplayOutput: true, AbortOnError: true}}
	err := c.Put("./remote.iml", "~/remote.iml")
	if err != nil {
		t.Errorf("Run put error, %s", err)
	} else {
		c.Run("cat ~/remote.iml")
	}
}

// Local support linux only
func TestLocal(t *testing.T) {
	c := &Config{loom.Config{User: "ubuntu", Host: "TESTSERVER",
		DisplayOutput: true, AbortOnError: true}}
	_, err := c.Local("echo testlocal")
	if err != nil {
		t.Errorf("Run local error, %s", err)
	}
}

func TestGet(t *testing.T) {
	c := &Config{loom.Config{User: "ubuntu", Host: "TESTSERVER",
		DisplayOutput: true, AbortOnError: true}}
	err := c.Get("~/remote.iml", "./remote.iml1")
	if err != nil {
		t.Errorf("Run get error, %s", err)
	}
}

func TestDeploy(t *testing.T) {
	c, err := MakeConfig("ubuntu", "TESTSERVER", true, true)
	if err != nil {
		t.Errorf("Make config error, %s", err)
	}
	cmd := Cmd{AptCache: true, UseSudo: true, CmdLine: "ls -la"}

	var i Provisioning
	i = c
	i.Execute(cmd)
}
