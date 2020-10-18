package main

import (
	"flag"

	"github.com/danielhoward314/go-react-boilerplate/server"
)

func main() {
	var entry string
	var static string
	var port string

	flag.StringVar(&entry, "entry", "./index.html", "the entrypoint to serve.")
	flag.StringVar(&static, "static", ".", "the directory to serve static files from.")
	flag.StringVar(&port, "port", "8000", "the `port` to listen on.")
	flag.Parse()

	server.Run(entry, static, port)
}
