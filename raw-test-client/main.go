package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected one argument; the peer address")
	}

	peer := os.Args[1]

	log.Printf("Connecting to %s", peer)

	c, err := net.Dial("tcp", peer)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	log.Printf("Connected; transferring data...")

	file, err := os.OpenFile("/dev/null", os.O_WRONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now().UnixNano()
	n, err := io.Copy(file, c)
	if err != nil {
		log.Printf("Error receiving data: %s", err)
	}
	end := time.Now().UnixNano()

	log.Printf("Received %d bytes in %s", n, time.Duration(end-start))
}
