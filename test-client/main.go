package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	proto "github.com/libp2p/go-libp2p-protocol"
	ma "github.com/multiformats/go-multiaddr"
)

const TestProtocol = proto.ID("/libp2p/test/data")

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

	host, err := libp2p.New(ctx, libp2p.NoListenAddrs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connecting to %s", pi.ID.Pretty())

	cctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err = host.Connect(cctx, *pi)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected; requesting data...")

	s, err := host.NewStream(cctx, pi.ID, TestProtocol)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	file, err := os.OpenFile("/dev/null", os.O_WRONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Printf("Transfering data...")

	start := time.Now().UnixNano()
	n, err := io.Copy(file, s)
	if err != nil {
		log.Printf("Error receiving data: %s", err)
	}
	end := time.Now().UnixNano()

	log.Printf("Received %d bytes in %s", n, time.Duration(end-start))
}
