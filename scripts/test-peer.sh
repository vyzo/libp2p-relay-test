#!/bin/bash
# script to test a peer through direct connection and all the relays

set -eu

if [ $# -ne 2 ]; then
   echo "Usage $0 peer-id ip-address"
   exit 42
fi

SCRIPT_PATH=$(dirname $(readlink -f $0))
ID=$1
IP=$2

echo ">>> Testing direct connection"
test-client /ip4/$IP/tcp/4001/ipfs/$ID

echo ">>> Testing relayed connections"
for relay in $(cat $SCRIPT_PATH/relays.txt); do
    echo ">>> Test $relay"
    test-client $relay/p2p-circuit/ipfs/$ID
done
