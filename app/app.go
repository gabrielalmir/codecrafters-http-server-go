package main

import (
	"http-server-starter-go/handler"
	"os"
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
	return handler.SendResponse(r, 200, map[string]string{}, "Hello World")
}

func handleEcho(r []byte) string {
	path := handler.Path(r)
	message := strings.Split(path, "/echo/")[1]
	return handler.SendResponse(r, 200, map[string]string{}, message)
}

func handleUserAgent(r []byte) string {
	req := strings.Split(string(r), "\r\n")
	var userAgent string
	for _, v := range req {
		if strings.Contains(v, "User-Agent") {
			userAgent = strings.Split(v, "User-Agent: ")[1]
		}
	}
	return handler.SendResponse(r, 200, map[string]string{}, userAgent)
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

		return handler.SendResponse(r, 200, map[string]string{"Content-Type": "application/octet-stream"}, content)
	}

	return handler.SendResponse(r, 405, map[string]string{}, "Method Not Allowed")
}
