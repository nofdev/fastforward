# FastForward
FastForward is a DevOps automate platform.

[![Build Status](https://travis-ci.org/nofdev/fastforward.svg?branch=master)](https://travis-ci.org/nofdev/fastforward)

## API Reference

### Basic Provisioning API
[![GoDoc](https://godoc.org/github.com/nofdev/fastforward/provisioning?status.svg)](https://godoc.org/github.com/nofdev/fastforward/provisioning)

Endpoint:

	http://0.0.0.0:7000/v1
	
Example Request:

	./jsonrpctest.py http://YOUR_FF_SERVER:7000/v1 \
	Provisioning.Exec \
	"{'User': 'ubuntu', \
	'Host': 'YOUR_REMOTE_SERVER', \
	'DisplayOutput': true, \
	'AbortOnError': true, \
	'AptCache': false, \
	'UseSudo': true, \
	'CmdLine': 'echo FastForward'}"

Example Response:

	{u'id': 1, u'result': u'FastForward\n', u'error': None}
	
Query Parameters:
* User - The username for remote server.
* Host - The remote server ip or FQDN.
* DisplayOutput - true/false, Show the execution output.
* AbortOnError - true/false, Ignore errors if set to true.
* AptCache - true/false, Using apt-get update before installation if set to true.
* UseSudo - true/false, Using sudo privilege for execution if set to true.
* CmdLine - The command line to be executed.

Status Codes:
* 200 - No error.
* 400 - Bad parameter.
* 500 - Server error.

### OpenStack Provisioning API
[![GoDoc](https://godoc.org/github.com/nofdev/fastforward/provisioning/api/rpc/json/openstack?status.svg)](https://godoc.org/github.com/nofdev/fastforward/provisioning/api/rpc/json/openstack)

Endpoint:

	http://0.0.0.0:7001/v1

Example Request:

	./jsonrpctest.py http://YOUR_FF_SERVER:7001/v1 \
	OpenStack.NovaController \
	"{'HostName': 'controller01'}"

Example Resopnse:

	{u'id': 1, u'result': None, u'error': None}

