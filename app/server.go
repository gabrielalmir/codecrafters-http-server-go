package main

import (
	"log"
	"net"
	"os"

	httphandler "github.com/codecrafters-io/http-server-starter-go/http"
)

func main() {
	// Bind to TCP port 4221 on all interfaces
	listener, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %s", err.Error())
			os.Exit(1)
		}

		go httphandler.Request(conn)
	}
}
