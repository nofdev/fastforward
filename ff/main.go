package main

import "os"
import "log"
import "github.com/mitchellh/cli"
import "github.com/jiasir/playback/command"

func main() {
	c := cli.NewCLI("FastForward", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"provision-api": provisionCommandFactory,
		"playback-api":  playbackCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}

type provision struct{}
type playback struct{}

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
func (p playback) Run(args []string) int {
	for _, arg :=range args {
		if arg == "start" {
			command.ExecuteWithOutput("playback-api")
		}
	}	
	return 0
}

// Help takes the help for playback-api.
func (p playback) Help() string {
	return "<start> Start the playback-api on 0.0.0.0:7001"
}

// Synopsis takes the synopsis of playback-api.
func (p playback) Synopsis() string {
	return "Start the playback-api on 0.0.0.0:7001"
}

func playbackCommandFactory() (cli.Command, error) {
	return &playback{}, nil
}
