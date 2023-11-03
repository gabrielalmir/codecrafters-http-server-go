package httphandler

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
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
		resp := "HTTP/1.1 200 OK\r\n\r\n"
		_, err = conn.Write([]byte(resp))
	} else if strings.Contains(path, "/echo") {
		content := strings.SplitN(path, "echo/", 2)[1]
		contentLength := len(content)
		resp := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %s\r\n\r\n%s", strconv.Itoa(contentLength), content)
		_, err = conn.Write([]byte(resp))
	} else if strings.Contains(path, "/user-agent") {
		req := strings.Split(string(buf), "\r\n")
		var userAgent string
		for _, v := range req {
			if strings.Contains(v, "User-Agent") {
				userAgent = strings.Split(v, "User-Agent: ")[1]
			}
		}
		resp := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %s\r\n\r\n%s", strconv.Itoa(len(userAgent)), userAgent)

		_, err = conn.Write([]byte(resp))
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
