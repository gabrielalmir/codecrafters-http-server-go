package handler

import (
	"log"
	"net"
	"os"
	"strings"
)

func Request(conn net.Conn, router Router) {
	// Read data from the connection
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	if err != nil {
		log.Fatalf("Error reading: %s", err.Error())
	}

	// Write data to the connection
	path := Path(buf)
	method := Method(buf)

	route, ok := router.Route(path, method)

	if !ok {
		conn.Write([]byte(NotFound(buf)))
		conn.Close()
	}

	_, err = conn.Write([]byte(route.Handler(buf)))

	if err != nil {
		log.Fatalf("Error writing: %s", err.Error())
		os.Exit(1)
	}

	err = conn.Close()

	if err != nil {
		log.Fatalf("Error closing: %s", err.Error())
		os.Exit(1)
	}
}

func Path(r []byte) string {
	return strings.SplitN(string(r), " ", 3)[1]
}

func Method(r []byte) string {
	return strings.SplitN(string(r), " ", 3)[0]
}
