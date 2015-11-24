package main

import (
    "github.com/gorilla/mux"
    "github.com/gorilla/rpc/v2"
    "github.com/gorilla/rpc/v2/json"
    "log"
    "net/http"
    "github.com/nofdev/fastforward/provisioning"
)

// Provisioning contains the configuration of ssh.
type Provisioning struct {
    provisioning.Conf
}

// File contains filename for GetFile and PutFile.
type File struct {
    RemoteFile string
    LocalFile string
}
// Args contains the method arguments for ssh login.
type Args struct {
    provisioning.Conf
    provisioning.Cmd
    File
}

// Result contains the API call results.
type Result interface {}

// Exec takes a command to be executed from API on the remote server.
func (p *Provisioning) Exec(r *http.Request, args *Args, result *Result) error {
	c, err := provisioning.MakeConfig(args.User, args.Host, args.DisplayOutput, args.AbortOnError); if err != nil {
		log.Printf("Make config error, %s", err)
	}
    
	cmd := provisioning.Cmd{AptCache: args.AptCache, UseSudo: args.UseSudo, CmdLine: args.CmdLine}

	var i provisioning.Provisioning
	i = c
	*result, _ = i.Execute(cmd)
    return nil
}

// GetFile copies the file from the remote host to the local FastForward server, using scp. Wildcards are not currently supported. 
func (p *Provisioning) GetFile(r *http.Request, args *Args, result *Result) error {
    c, err := provisioning.MakeConfig(args.User, args.Host, args.DisplayOutput, args.AbortOnError); if err != nil {
        log.Printf("Make config error, %s", err)
    }
    
    if args.RemoteFile== "" || args.LocalFile == "" {
        *result = "RemoteFile or LocalFile are needed."
    }
    *result = c.GetFile(args.RemoteFile, args.LocalFile)
    return nil
}

func main() {
    s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(), "application/json")
    provisioning := new(Provisioning)
    s.RegisterService(provisioning, "")
    r := mux.NewRouter()
    r.Handle("/v1", s)
    http.ListenAndServe(":7000", r)
}