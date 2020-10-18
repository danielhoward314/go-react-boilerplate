package server

import (
	"log"
	"net/http"
	"os"
	"time"

	controllers "github.com/danielhoward314/go-react-boilerplate/server/controllers"
	middleware "github.com/danielhoward314/go-react-boilerplate/server/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Run(html, webpack string) {

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/health", controllers.GetHealth).Methods("GET")

	apiRouter.HandleFunc("/health", controllers.GetHealth).Methods(http.MethodGet, http.MethodOptions)

	env := goDotEnvVariable("ENVIRONMENT")
	if env == "production" {
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
