package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Time struct holds the JSON data
type Time struct {
	Time    string `json:"time"` // this will be changed to {"time": string}
	Changed bool   `json:"changed"`
}

func main() {
	// the format is a little bit wired. see http://golang.org/pkg/time/#pkg-constants
	f := "2015-05-05 20:15:00"

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	arg := strings.Split(string(data), " ")[0]

	if strings.Contains(arg, "=") {
		f = strings.Split(arg, "=")[1]
	}

	t := time.Now()
	m := Time{t.Format(f), true} // create Time struct with string
	s, err := json.Marshal(m)    // produce JSON from Time struct
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", s)
	}
}

// for more details see http://docs.ansible.com/ansible/developing_modules.html
