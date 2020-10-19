package server

import (
	"log"
	"net/http"
	"time"

	"github.com/danielhoward314/go-react-boilerplate/server/controllers"
	"github.com/danielhoward314/go-react-boilerplate/server/middleware"
	"github.com/gorilla/mux"
)

func Run(env, html, webpack string) {

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/health", controllers.GetHealth).Methods(http.MethodGet, http.MethodOptions)

	if env == "prod" {
		// Serve webpack assets directly.
		r.PathPrefix("/dist/js/").Handler(http.StripPrefix("/dist/js/", http.FileServer(http.Dir(webpack))))

		// Catch-all: Serve our JavaScript application's entry-point (index.html).
		r.PathPrefix("/").HandlerFunc(controllers.SpaHandler(html))
	}

	r.Use(mux.CORSMethodMiddleware(r))

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
