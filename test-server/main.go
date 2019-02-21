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
		"/ip4/35.175.201.95/tcp/4001/ipfs/QmPZ3ZfLNSiiQXj75Ft5Qs2qnPCt1Csp1nZBUppfY997Yb",
		"/ip4/18.232.56.199/tcp/4001/ipfs/QmWKiiMyB1qtBs6atLC7UXTk6kw8au2p6LGrET6AkmTCRr",
		"/ip4/35.172.224.10/tcp/4001/ipfs/QmYwiV8rxauHpPCbLukjwHAsy3AkLU4U7VF7NrMTxWsJS8",
		"/ip4/18.188.53.89/tcp/4001/ipfs/Qmbyw5e8pemFWb27Z6JJtR8gQ3hMRRkzynSYn66DmQ6yP4",
		"/ip4/18.224.32.43/tcp/4001/ipfs/QmaXbbcs7LRFuEoQcxfXqziZATzS68WT5DgFjYFgn3YYLX",
		"/ip4/18.191.192.216/tcp/4001/ipfs/QmeaqKEz7NipyX5YJhtFb7NdAGjgp4wzGfHEpXESrtSL6Y",
		"/ip4/18.236.86.245/tcp/4001/ipfs/Qmbn1WTN8WRPPhmayAe2eiMdcgKqdtm9MbWKga59bhFEoW",
		"/ip4/34.209.142.65/tcp/4001/ipfs/QmaZK14yReY9Bio4Wxj8NLH7vmEBh91J4rS7HVktRhHiV8",
		"/ip4/54.191.81.56/tcp/4001/ipfs/QmbYSi1jdhnz5D4PiY5AB1CVRYK37A4K4jNamp655wTvWC",
		"/ip4/34.251.29.51/tcp/4001/ipfs/QmTxKmg6bHwH3kvAwiuFpir9Ee3ogStqjkNeg8Lt7V4iWW",
		"/ip4/54.171.172.82/tcp/4001/ipfs/QmSqKqGZ5yr8adWK5txTn1BqG3Em6RhjL3LM2Q4pjUfYMP",
		"/ip4/54.194.107.96/tcp/4001/ipfs/QmTt4EmqFDsYJiypRD5U4dUDzramwLVCgDQJS9GzkHWJGR",
		"/ip4/3.121.215.168/tcp/4001/ipfs/QmbLjvQtV1kkiBh9cWGeXaBoaoWkdnm6QDUFrf2LAkecTH",
		"/ip4/18.184.154.188/tcp/4001/ipfs/QmZ5UovPa7H7Mqe31crwhmoEgAz8XZ6s1stzzWSNKFzuqU",
		"/ip4/18.184.60.108/tcp/4001/ipfs/QmQUPL5gvWp3VDBrouhpZx2L6un2ZBzMbLmcnjK45TMuVv",
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

		cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		err = host.Connect(cctx, *pi)
		cancel()

		if err != nil {
			log.Printf("ERROR connecting to %s: %s", pi.ID.Pretty(), err)
		}
	}
	log.Printf("Done")
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
