package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	// TODO: Add db connection
	app := Config{}
	// create the server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	log.Printf("Starting authentication service on %s port\n", webPort)

	// start the server
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
