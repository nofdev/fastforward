/*
Package main is the ff command line interface.
*/
package main

import "os"
import "log"
import "github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/mitchellh/cli"
import "github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/jiasir/playback/command"
import "github.com/nofdev/fastforward/config"

func main() {
	c := cli.NewCLI("ff", "0.0.3")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"provision-api": provisionCommandFactory,
		"playback-api":  playbackAPICommandFactory,
		"playback":      playbackCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}

type provision struct{}
type playbackAPI struct{}
type playback struct{}

// c is the FastForward configuration instance.
var conf *config.Conf
var c = conf.LoadConf()

// Run takes provision-api command.
func (p provision) Run(args []string) int {
	// TODO: Refactor arg parse.
	// TODO: Refactor exit status.
	for _, arg := range args {
		if arg == "start" {
			command.ExecuteWithOutput("provision-api")
		}
	}
	return 0
}

// Help takes the help for provision.
func (p provision) Help() string {
	return "<start> Start the provision-api on 0.0.0.0:7000"
}

// Synopsis takes the synopsis of playback-api.
func (p provision) Synopsis() string {
	return "Start the provision-api on 0.0.0.0:7000"
}

func provisionCommandFactory() (cli.Command, error) {
	return &provision{}, nil
}

// Run takes playback-api command.
func (p playbackAPI) Run(args []string) int {
	for _, arg := range args {
		if arg == "start" {
			if c.DEFAULT["provisioning_driver"] == "playback" {
				command.ExecuteWithOutput("playback-api")
			} else {
				log.Fatalf("The driver is not playback, current is: %s", c.DEFAULT["provisioning_driver"])
			}
		}
	}
	return 0
}

// Help takes the help for playback-api.
func (p playbackAPI) Help() string {
	return "<start> Start the playback-api on 0.0.0.0:7001"
}

// Synopsis takes the synopsis of playback-api.
func (p playbackAPI) Synopsis() string {
	return "Start the playback-api on 0.0.0.0:7001"
}

func playbackAPICommandFactory() (cli.Command, error) {
	return &playbackAPI{}, nil
}

func playbackCommandFactory() (cli.Command, error) {
	return &playback{}, nil
}

// Run takes ansible-playbook. Ansible has no API client for golang, We need to use the cmd-line.
func (p playback) Run(args []string) int {
	command.ExecuteWithOutput("ansible-playbook", args...)
	// TODO: Refactor exit status.
	return 0
}

// Help takes the command line help.
func (p playback) Help() string {
	return "<playback> <yaml> [--extra-vars] Run the playback yaml"
}

func (p playback) Synopsis() string {
	return "Run the playback yaml"
}
