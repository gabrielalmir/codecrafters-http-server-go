package main

import (
	"fmt"
	"http-server-starter-go/handler"
	"os"
	"strconv"
	"strings"
)

type App struct {
	Router handler.Router
}

func NewApp() *App {
	app := &App{
		Router: handler.Router{},
	}

	app.Router.AddRoute(handler.Route{Path: "^/$", Handler: handleRoot, Method: "GET"})
	app.Router.AddRoute(handler.Route{Path: "^/echo/.*$", Handler: handleEcho, Method: "GET"})
	app.Router.AddRoute(handler.Route{Path: "^/user-agent$", Handler: handleUserAgent, Method: "GET"})
	app.Router.AddRoute(handler.Route{Path: "^/files/.*$", Handler: handleFile, Method: "GET"})

	return app
}

func handleRoot(r []byte) string {
	return "HTTP/1.1 200 OK\r\n\r\n"
}

func handleEcho(r []byte) string {
	path := handler.Path(r)
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

func handleFile(r []byte) string {

	path := handler.Path(r)
	method := handler.Method(r)

	filename := strings.Split(path, "/files/")[1]
	directory := os.Getenv("DIRECTORY")

	if method == "GET" {
		file := handler.File{Filename: filename, Directory: directory}
		content, err := file.Handle()

		if err != nil {
			return handler.NotFound(r)
		}

		return fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %s\r\n\r\n%s", strconv.Itoa(len(content)), content)
	}

	return "HTTP/1.1 405 Method Not Allowed\r\n\r\n"
}
