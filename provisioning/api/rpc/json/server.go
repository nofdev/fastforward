package main

import (
    "github.com/gorilla/mux"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"
    "log"
    "net/http"
    "github.com/nofdev/fastforward/provisioning"
)

type Config struct {
    provisioning.Conf
}

type Args struct {
    provisioning.Conf
    provisioning.Cmd
}

type Result interface {}

func (this *Config) Exec(r *http.Request, args *Args, result *Result) error {

	c, err := provisioning.MakeConfig(args.User, args.Host, args.DisplayOutput, args.AbortOnError); if err != nil {
		log.Printf("Make config error, %s", err)
	}
	cmd := provisioning.Cmd{AptCache: args.AptCache, UseSudo: args.UseSudo, CmdLine: args.CmdLine}

	var i provisioning.Provisioning
	i = c
	*result, _ = i.Execute(cmd)
    return nil
}

func main() {
    s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(), "application/json")
    s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
    config := new(Config)
    s.RegisterService(config, "")
    r := mux.NewRouter()
    r.Handle("/v1", s)
    http.ListenAndServe(":7000", r)
}