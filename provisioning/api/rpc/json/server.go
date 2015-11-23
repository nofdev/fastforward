package main

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"net"
	
	"github.com/nofdev/fastforward/provisioning"
)

// Args struct will be the json args
type Args struct {
	User string
	Host string
	Output bool
	Abort bool
	provisioning.Cmd
}

type Api int
var i provisioning.Provisioning

// Execute command from api
func (a *Api) Exec(args *Args, reply *string) error {
	c, err := provisioning.MakeConfig(args.User, args.Host, args.Output, args.Abort)
	checkError(err)

	i = c
	i.Execute(args.Cmd)
	return nil
}

func main() {
	provision := new(Api)
	rpc.Register(provision)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
