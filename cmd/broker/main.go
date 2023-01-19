package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoTalkie/go-talkie-broker/internal/server"
)

const webPort = "80"

func main() {
	app := server.Config{}

	log.Printf("Starting broker service on port %s", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
