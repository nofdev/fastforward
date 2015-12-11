#!/usr/bin/env python

# USAGE:
# ./jsonrpctest.py http://localhost:7000/v1 Provisioning.GetFile "{'User': 'ubuntu', 'Host': 'FASTFORWARD', 'RemoteFile': 'testputstring', 'Localfile': 'testputstring'}"

import urllib2
import json
import sys
import yaml


def rpc_call(url, method, args):
    data = json.dumps({
        'id': 1,
        'method': method,
        'params': [args]
    }).encode()
    req = urllib2.Request(url,
        data,
        {'Content-Type': 'application/json'})
    f = urllib2.urlopen(req)
    response = f.read()
    return json.loads(response)

url = sys.argv[1]
method = sys.argv[2]
args = yaml.load(sys.argv[3])
print rpc_call(url, method, args)
