package provisioning

import "testing"
import "github.com/jiasir/playback/command"

func TestCommand(t *testing.T) {
	command.ExecuteWithOutput("echo", "OK")
}
