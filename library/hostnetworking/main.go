package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/jiasir/playback/command"
	"github.com/nofdev/fastforward/library/common"
)

func main() {
	var interfaces = new(common.Interfaces)
	// ansible puts module args to a file wich is os.Args[1] on remote server
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	args := strings.Split(string(data), " ") // [k=v k=v k=v]

	// arg is k=v
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			k := strings.Split(arg, "=")[0] // the key
			v := strings.Split(arg, "=")[1] // the value
			// set the k v to struct
			interfaces.InitInterfaces(k, v)
		}
	}

	interfaces.Changed = true
	// use interfaces.d instead of /etc/network/interfaces
	interfaces.PurgeMainConf()
	// setup internal nic
	interfaces.SetInternalNIC()

	if interfaces.ExternalNIC != "" {
		// setup external nic
		interfaces.SetExternalNIC()
	}

	if interfaces.Restart {
		// restart the system for take effect
		command.ExecuteWithOutput("sudo", "shuwdown", "-r", "+1", "FastForward takes reboot")
	}

	output, err := json.Marshal(*interfaces) //produce JSON from interfaces struct
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", output)
	}

}

// for more details see http://docs.ansible.com/ansible/developing_modules.html
