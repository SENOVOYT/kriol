// Filename: cmd/api/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// The Application version number
const version = "1.0.0"
// The Configuration setting

type config struct {
	port int
	env string // Development , staging, Production, etc.
}
//Dependency Injection
type application struct {
	config config
	logger *log.Logger
}
func main() {

	var cfg config
	//read in the flags that are needed to populate our config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env","development", "Environment (development | stagging | production )")
	flag.Parse()
	//create a logger
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime )
	//create an instance of our application struct
	app := &application{
		config: cfg,
		logger: logger,
	}
	//Create our new servermux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	//create our Http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),// using the routes function from routes,go
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	//start our server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}