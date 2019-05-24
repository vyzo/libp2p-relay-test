package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	inet "github.com/libp2p/go-libp2p-net"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	proto "github.com/libp2p/go-libp2p-protocol"
	identify "github.com/libp2p/go-libp2p/p2p/protocol/identify"
	ma "github.com/multiformats/go-multiaddr"
)

const TestProtocol = proto.ID("/libp2p/test/data")

var relayAddrs []ma.Multiaddr
var testFilePath string

func init() {
	for _, s := range []string{
		"/ip4/147.75.84.57/tcp/4001/p2p/QmZP8NCi1L2LS8K2DoG175tH4mSe8Z4ygcVXkwFxnyeMLL",
		"/ip4/147.75.80.35/tcp/4001/p2p/QmU5jkMcfaZ4N1B4MzMXdCZY2pJ3re5YaPB7UjiyqShwT9",
		"/ip4/147.75.80.33/tcp/4001/p2p/QmSn4heyhN27h8QX6hEgWeLw2kMo75ZEwj1ECk1c4khUPN",
		"/ip4/147.75.106.163/tcp/4001/p2p/QmRdjvsyoNjA2ZfAQBtQ7A2m5NmtSXLgxE55Brn1AUjZ1v",
		"/ip4/139.178.64.247/tcp/4001/p2p/QmVbReJM8RpHxZcMCmdSubBfBX7VdYiCn4piAmFxaGUDmJ",
		"/ip4/147.75.105.219/tcp/4001/p2p/QmTtFWmQ3qrp166m96ibL2jW2Doz4tJjo2CwQfYNaFb3XZ",
		"/ip4/139.178.69.3/tcp/4001/p2p/QmdGQoGuK3pao6bRDqGSDvux5SFHa4kC2XNFfHFcvcbydY",
		"/ip4/147.75.109.65/tcp/4001/p2p/QmcYZo7xDLm8sNakKe8UK9AXjoGXGvngpD6apqmTqu7HzU",
		"/ip4/139.178.69.15/tcp/4001/p2p/QmTg7DdGo519B7KdzFppTaK5i8WGWu9erGhxsLafxRd6gp",
	} {
		a, err := ma.NewMultiaddr(s)
		if err != nil {
			log.Fatal(err)
		}
		relayAddrs = append(relayAddrs, a)
	}
}

func main() {
	identify.ClientVersion = "test-server/0.1"

	listenAddrStr := flag.String("listen", "/ip4/0.0.0.0/tcp/4001", "server listen address")
	relayAddrStr := flag.String("relays", "", "comma separated list of relays")
	testFile := flag.String("file", "data", "data file to serve")

	flag.Parse()

	if *relayAddrStr != "" {
		relayAddrs = relayAddrs[:0]
		for _, s := range strings.Split(*relayAddrStr, ",") {
			a, err := ma.NewMultiaddr(s)
			if err != nil {
				log.Fatal(err)
			}
			relayAddrs = append(relayAddrs, a)
		}
	}

	if _, err := os.Stat(*testFile); err != nil {
		log.Fatal(err)
	}
	testFilePath = *testFile

	ctx := context.Background()

	host, err := libp2p.New(ctx, libp2p.ListenAddrStrings(*listenAddrStr))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("I am %s\n", host.ID().Pretty())

	host.SetStreamHandler(TestProtocol, handleStream)

	log.Printf("Connecting to relays...")
	for _, a := range relayAddrs {
		pi, err := pstore.InfoFromP2pAddr(a)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Connecting to %s", pi.ID.Pretty())

		cctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		err = host.Connect(cctx, *pi)
		cancel()

		if err != nil {
			log.Printf("ERROR connecting to %s: %s", pi.ID.Pretty(), err)
		}
	}
	log.Printf("Ready")

	select {}
}

func handleStream(s inet.Stream) {
	defer s.Close()

	log.Printf("Incoming connection from %s", s.Conn().RemoteMultiaddr())

	file, err := os.Open(testFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now().UnixNano()
	n, err := io.Copy(s, file)
	if err != nil {
		log.Printf("Error transmiting file: %s", err)
	}
	end := time.Now().UnixNano()
	log.Printf("Transmitted %d bytes in %s", n, time.Duration(end-start))
}
