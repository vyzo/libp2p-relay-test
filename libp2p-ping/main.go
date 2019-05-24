package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	ping "github.com/libp2p/go-libp2p/p2p/protocol/ping"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected one argument; the peer multiaddr")
	}

	a, err := ma.NewMultiaddr(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	pi, err := pstore.InfoFromP2pAddr(a)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	host, err := libp2p.New(ctx, libp2p.NoListenAddrs, libp2p.EnableRelay())
	if err != nil {
		log.Fatal(err)
	}

	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	err = host.Connect(cctx, *pi)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	pctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	ch := ping.Ping(pctx, host, pi.ID)
	defer cancel()

	for r := range ch {
		fmt.Printf("%s\n", r.RTT)
	}
}
