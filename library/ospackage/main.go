package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nofdev/fastforward/library/common"
)

func main() {
	var packages = new(common.Packages)
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
			packages.InitPackages(k, v)
		}
	}

	packages.Changed = true
	// install packages repo and openstack client
	packages.InstallRepo()
	packages.InstallClient()

	output, err := json.Marshal(*packages) //produce JSON from interfaces struct
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", output)
	}

}

// for more details see http://docs.ansible.com/ansible/developing_modules.html
