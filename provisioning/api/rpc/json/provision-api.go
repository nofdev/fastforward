// Package main is a provisioning api based on JSON-RPC 2.0
package main

import (
	"log"
	"net/http"

	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/gorilla/rpc/v2"
	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/gorilla/rpc/v2/json"
	"github.com/nofdev/fastforward/provisioning"
)

// Provisioning for API.
type Provisioning struct{}

// File contains filename for GetFile and PutFile.
type File struct {
	RemoteFile string
	LocalFile  string
	Data       string
}

// Args contains the method arguments for ssh login.
type Args struct {
	provisioning.Conf
	provisioning.Cmd
	File
}

// Result contains the API call results.
type Result interface{}

// checkFile ensure the args about PutFile and GetFile return the error to API caller.
func checkFile(r *http.Request, args *Args, result *Result) {
	if args.RemoteFile == "" || args.LocalFile == "" {
		*result = "RemoteFile or LocalFile are needed."
		log.Printf("Request: %s, Error: %s", *r, *result)
	}
}

// Exec takes a command to be executed from API on the remote server.
//  Args: {'User': 'USERNAME', 'Host': 'REMOTESERVER', 'CmdLine': 'CMD'}
//  Optional Args: {'AptCache': true, 'UseSudo': true, 'Password': 'PASS', 'KeyFiles': 'IDRSA', 'DisplayOutput': true, 'AbortOnError': true}
func (p *Provisioning) Exec(r *http.Request, args *Args, result *Result) error {
	cmd := provisioning.Cmd{AptCache: args.AptCache, UseSudo: args.UseSudo, CmdLine: args.CmdLine}
	i := provisioning.Provisioning(args)
	*result, _ = i.Execute(cmd)
	log.Printf("Request: %s, Method: Exec, Args: %s, Result: %s", *r, *args, *result)
	return nil
}

// GetFile copies the file from the remote host to the local FastForward server, using scp. Wildcards are not currently supported.
//  Args: {'User': 'USERNAME', 'Host': 'REMOTESERVER', 'RemoteFile': 'FILENAME', 'LocalFile': 'FILENAME'}
//  Optional Args: {'Password': 'PASS', 'KeyFiles': 'IDRSA', 'DisplayOutput': true, 'AbortOnError': true}
func (p *Provisioning) GetFile(r *http.Request, args *Args, result *Result) error {
	checkFile(r, args, result)
	i := provisioning.Provisioning(args)
	*result = i.GetFile(args.RemoteFile, args.LocalFile)
	log.Printf("Request: %s, Method: GetFile, Args: %s, Result: %s", *r, *args, *result)
	return nil
}

// PutFile copies one or more local files to the remote host, using scp. localfiles can contain wildcards, and remotefile can be either a directory or a file.
//  Args: {'User': 'USERNAME', 'Host': 'REMOTESERVER', 'LocalFile': 'FILENAME', 'RemoteFile': 'FILENAME'}
//  Optional Args: {'Password': 'PASS', 'KeyFiles': 'IDRSA', 'DisplayOutput': true, 'AbortOnError': true}
func (p *Provisioning) PutFile(r *http.Request, args *Args, result *Result) error {
	checkFile(r, args, result)
	i := provisioning.Provisioning(args)
	*result = i.PutFile(args.LocalFile, args.RemoteFile)
	log.Printf("Request: %s, Method: PutFile, Args: %s, Result: %s", *r, *args, *result)
	return nil
}

// PutString generates a new file on the remote host containing data. The file is created with mode 0644.
//  Args: {'User': 'USERNAME', 'Host': 'REMOTESERVER', 'Data': 'STRING', 'RemoteFile': 'FILENAME'}
//  Optional Args: {'Password': 'PASS', 'KeyFiles': 'IDRSA', 'DisplayOutput': true, 'AbortOnError': true}
func (p *Provisioning) PutString(r *http.Request, args *Args, result *Result) error {
	i := provisioning.Provisioning(args)
	*result = i.PutString(args.Data, args.RemoteFile)
	log.Printf("Request: %s, Method: PutString, Args: %s, Result: %s", *r, *args, *result)
	return nil
}

// Self executes a command on the FastForward API server.
//  Args: {'CmdLine': 'CMD'}
func (p *Provisioning) Self(r *http.Request, args *Args, result *Result) error {
	i := provisioning.Provisioning(args)
	*result, _ = i.Self(args.Cmd)
	log.Printf("Request: %s, Method: Self, Args: %s, Result: %s", *r, *args, *result)
	return nil
}

func main() {
	s := rpc.NewServer()
	log.Printf("API Server started")
	s.RegisterCodec(json.NewCodec(), "application/json")
	provisioning := new(Provisioning)
	s.RegisterService(provisioning, "")
	log.Printf("Register Provisioning service")
	r := mux.NewRouter()
	r.Handle("/v1", s)
	log.Printf("Handle API version 1")
	log.Printf("Listen on port 7000")
	http.ListenAndServe(":7000", r)
}
