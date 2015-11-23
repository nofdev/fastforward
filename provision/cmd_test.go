package provision

import "testing"
import "github.com/wingedpig/loom"

type Ssh struct {
	loom.Config
}

func TestRun(t *testing.T) {
	c := &Ssh{loom.Config{User:"ubuntu",Host:"10.32.151.68",
		DisplayOutput:true, AbortOnError:true}}
	_, err := c.Run("ls -la")
	if err != nil {
		t.Errorf("Run cmd error, %s", err)
	}
}

func TestSudo(t *testing.T) {
	c := &Ssh{loom.Config{User:"ubuntu",Host:"10.32.151.68",
		DisplayOutput:true, AbortOnError:true}}
	_, err := c.Sudo("ls -la"); if err != nil {
		t.Errorf("Run sudo error, %s", err)
	}
}

func TestPutString(t *testing.T) {
	c := &Ssh{loom.Config{User:"ubuntu",Host:"10.32.151.68",
		DisplayOutput:true, AbortOnError:true}}
	err := c.PutString("TestPutString\nTestPutString", "~/testputstring"); if err != nil {
		t.Errorf("Run putstring error, %s", err)
	} else {
		c.Run("cat ~/testputstring")
	}
}

func TestPut(t *testing.T) {
	c := &Ssh{loom.Config{User:"ubuntu",Host:"10.32.151.68",
		DisplayOutput:true, AbortOnError:true}}
	err := c.Put("./remote.iml", "~/remote.iml"); if err != nil {
		t.Errorf("Run put error, %s", err)
	} else {
		c.Run("cat ~/remote.iml")
	}
}

// Local support linux only
func TestLocal(t *testing.T) {
	c := &Ssh{loom.Config{User:"ubuntu",Host:"10.32.151.68",
		DisplayOutput:true, AbortOnError:true}}
	_, err := c.Local("echo testlocal"); if err != nil {
		t.Errorf("Run local error, %s", err)
	}
}

func TestGet(t *testing.T) {
	c := &Ssh{loom.Config{User:"ubuntu",Host:"10.32.151.68",
		DisplayOutput:true, AbortOnError:true}}
	err := c.Get("~/remote.iml", "./remote.iml1"); if err != nil {
		t.Errorf("Run get error, %s", err)
	}
}

func TestDeploy(t *testing.T) {
	c, err := MakeConfig("ubuntu", "10.32.151.68", true, true); if err != nil {
		t.Errorf("Make config error, %s", err)
	}
	cmd := Cmd{AptCache: true, UseSudo:true, CmdLine: "ls -la"}

	var i Provisioning
	i = c
	i.Execute(cmd)
}