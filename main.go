package main

import (
	"flag"

	"github.com/danielhoward314/go-react-boilerplate/server"
)

func main() {
	var env string
	var html string
	var webpack string
	flag.StringVar(&env, "env", "dev", "the environment to run in: dev | prod ")
	flag.StringVar(&html, "html", "./index.html", "the path of index.html to serve.")
	flag.StringVar(&webpack, "webpack", ".", "the path to webpack bundle's dir.")
	flag.Parse()
	server.Run(env, html, webpack)
}
