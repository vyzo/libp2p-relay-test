#!/bin/bash
# Provisioning script for test nodes

set -eu

GOVER=1.11.5

# set up
sudo apt-get -y update

# set up go
echo ">>> Setup Go ${GOVER}"
wget https://dl.google.com/go/go${GOVER}.linux-amd64.tar.gz
sudo tar -C /usr/local -xzvf  go${GOVER}.linux-amd64.tar.gz

mkdir go
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin

# build test programs
echo ">>> Building test programs"
go get github.com/vyzo/libp2p-relay-test
cd ~/go/src/github.com/vyzo/libp2p-relay-test
go install ./...
