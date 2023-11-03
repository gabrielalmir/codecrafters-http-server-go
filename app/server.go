package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	// Bind to TCP port 4221 on all interfaces
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	// Read data from the connection
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		os.Exit(1)
	}

	// Write data to the connection
	_, err = conn.Write([]byte("HTTP/1.1 200 OK \r\n\r\n"))

	if err != nil {
		fmt.Println("Error writing: ", err.Error())
		os.Exit(1)
	}

	err = conn.Close()

	if err != nil {
		fmt.Println("Error closing: ", err.Error())
		os.Exit(1)
	}
}
