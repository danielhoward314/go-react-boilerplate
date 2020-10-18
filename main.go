package main

import (
	"flag"

	"github.com/danielhoward314/go-react-boilerplate/server"
)

func main() {
	var html string
	var webpack string

	flag.StringVar(&html, "html", "./index.html", "the path of the html to serve.")
	flag.StringVar(&webpack, "webpack", ".", "the directory to serve the webpack bundle from.")
	flag.Parse()

	server.Run(html, webpack)
}
