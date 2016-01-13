package provisioning

import "testing"
import "github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/jiasir/playback/command"

func TestCommand(t *testing.T) {
	command.ExecuteWithOutput("echo", "OK")
}
