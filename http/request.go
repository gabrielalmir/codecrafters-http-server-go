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
	method := Method(buf)

	// generic router implementation
	router := Router{}

	router.AddRoute(Route{Path: "^/$", Handler: handleRoot, Method: "GET"})
	router.AddRoute(Route{Path: "^/echo/.*$", Handler: handleEcho, Method: "GET"})
	router.AddRoute(Route{Path: "^/user-agent$", Handler: handleUserAgent, Method: "GET"})
	router.AddRoute(Route{Path: "^/files/.*$", Handler: handleFile, Method: "GET"})

	route, ok := router.Route(path, method)

	if !ok {
		route = Route{Path: "^/404$", Handler: handleNotFound}
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

func handleRoot(r []byte) string {
	return "HTTP/1.1 200 OK\r\n\r\n"
}

func handleEcho(r []byte) string {
	path := Path(r)
	message := strings.Split(path, "/echo/")[1]
	return fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %s\r\n\r\n%s", strconv.Itoa(len(message)), message)
}

func handleUserAgent(r []byte) string {
	req := strings.Split(string(r), "\r\n")
	var userAgent string
	for _, v := range req {
		if strings.Contains(v, "User-Agent") {
			userAgent = strings.Split(v, "User-Agent: ")[1]
		}
	}
	return fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %s\r\n\r\n%s", strconv.Itoa(len(userAgent)), userAgent)
}

func handleNotFound(r []byte) string {
	return "HTTP/1.1 404 Not Found\r\n\r\n"
}

func handleFile(r []byte) string {

	path := Path(r)
	method := Method(r)

	filename := strings.Split(path, "/files/")[1]
	directory := os.Getenv("DIRECTORY")

	if method == "GET" {
		file := File{directory: directory, matcher: filename}
		content, err := file.Handle()

		if err != nil {
			return handleNotFound(r)
		}

		return fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %s\r\n\r\n%s", strconv.Itoa(len(content)), content)
	}

	return "HTTP/1.1 405 Method Not Allowed\r\n\r\n"
}
