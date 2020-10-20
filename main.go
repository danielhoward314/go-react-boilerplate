package main

import (
	"github.com/danielhoward314/go-react-boilerplate/server"
	"github.com/namsral/flag"
)

func main() {
	var env string
	var port int

	var html string
	var webpack string

	flag.StringVar(&env, "env", env, "dev | prod")
	flag.IntVar(&port, "port", port, "Port number")
	flag.StringVar(&html, "html", html, "path of index.html to serve")
	flag.StringVar(&webpack, "webpack", webpack, "path to webpack bundle's dir")
	flag.Parse()
	server.Run(port, env, html, webpack)
}
