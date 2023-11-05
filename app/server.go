package main

import (
	"flag"
	"http-server-starter-go/handler"
	"log"
	"net"
	"os"
	"strconv"
)

func init() {
	host := flag.String("host", "0.0.0.0", "Host to serve")
	port := flag.Int("port", 4221, "Port to serve")
	directory := flag.String("directory", "./", "Directory to serve")

	flag.Parse()

	os.Setenv("DIRECTORY", *directory)
	os.Setenv("HOST", *host)
	os.Setenv("PORT", strconv.Itoa(*port))
}

func main() {
	listener, err := net.Listen("tcp", os.Getenv("HOST")+":"+os.Getenv("PORT"))

	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
		os.Exit(1)
	}

	app := NewApp()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %s", err.Error())
			os.Exit(1)
		}

		go handler.Request(conn, app.Router)
	}
}
