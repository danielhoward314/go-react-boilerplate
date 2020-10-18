package controllers

import (
	"log"
	"net/http"
	"path/filepath"
)

func SpaHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// get the absolute path to prevent directory traversal
		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			// if we failed to get the absolute path respond with a 400 bad request
			// and stop
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(path)
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}
