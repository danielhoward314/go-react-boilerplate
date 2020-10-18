package server

import (
	"log"
	"net/http"
	"time"

	controllers "github.com/danielhoward314/go-react-boilerplate/server/controllers"
	"github.com/gorilla/mux"
)

func Run(entry, static, port string) {

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/health", controllers.GetHealth).Methods("GET")

	// Serve static assets directly.
	r.PathPrefix("/dist").Handler(http.FileServer(http.Dir(static)))

	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").HandlerFunc(controllers.SpaHandler(entry))

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
