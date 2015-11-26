/*
Package provisioning provides an API of JSON-RPC 2.0

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
	User - The username for remote server.
	Host - The remote server ip or FQDN.
	DisplayOutput - true/false, Show the execution output.
	AbortOnError - true/false, Ignore errors if set to true.
	AptCache - true/false, Using apt-get update before installation if set to true.
	UseSudo - true/false, Using sudo privilege for execution if set to true.
	CmdLine - The command line to be executed.

Status Codes:
	200 - No error.
	400 - Bad parameter.
	500 - Server error.
*/
package provisioning
