package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("expected one argument; the listen and destination addresses")
	}

	listenAddr := os.Args[1]
	peerAddr := os.Args[2]

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}

		go proxy(c, peerAddr)
	}
}

func proxy(c net.Conn, peerAddr string) {
	log.Printf("Incomming connection from %s", c.RemoteAddr())
	log.Printf("Connecting to %s", peerAddr)

	r, err := net.Dial("tcp", peerAddr)
	if err != nil {
		log.Printf("Error connecting to peer: %s", err)
		c.Close()
		return
	}

	log.Printf("Proxying connection...")
	go func() {
		io.Copy(c, r)
		c.Close()
	}()

	go func() {
		io.Copy(r, c)
		r.Close()
	}()
}
