#!/usr/bin/env bash

go get -u -v github.com/nofdev/fastforward
sudo mkdir /etc/fastforward
sudo cp $GOPATH/src/github.com/nofdev/fastforward/config/fastforward.conf /etc/fastforward
sudo chmod +x $GOPATH/src/github.com/nofdev/fastforward/bin/linux_amd64/*
sudo cp -r $GOPATH/src/github.com/nofdev/fastforward/bin/linux_amd64/* /usr/local/bin
