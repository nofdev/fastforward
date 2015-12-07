package main

import "os"
import "log"
import "github.com/mitchellh/cli"
import "github.com/jiasir/playback/command"
import "github.com/nofdev/fastforward/config"

func main() {
	c := cli.NewCLI("FastForward", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"provision-api": provisionCommandFactory,
		"playback-api":  playbackAPICommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}

type provision struct{}
type playbackAPI struct{}

// c is the FastForward configuration instance.
var conf *config.Conf
var c = conf.LoadConf()

// Run takes provision-api command.
func (p provision) Run(args []string) int {
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
