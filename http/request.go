package httphandler

import (
	"log"
	"net"
	"os"
	"strings"
)

func Request(conn net.Conn) {
	// Read data from the connection
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	if err != nil {
		log.Fatalf("Error reading: %s", err.Error())
	}

	// Write data to the connection
	path := Path(buf)

	if path == "/" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.Contains(path, "echo/") {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n" + strings.SplitN(path, "echo/", 2)[1]))
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

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
