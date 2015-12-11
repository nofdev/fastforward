#!/usr/bin/env bash

go get -u -v github.com/nofdev/fastforward
mkdir /etc/fastforward
cp $GOPATH/src/github.com/nofdev/fastforward/config/fastforward.conf /etc/fastforward
