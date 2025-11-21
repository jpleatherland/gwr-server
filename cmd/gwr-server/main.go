package main

import (
	"github.com/jpleatherland/gwr-server/internal/server"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: server.NewHandler(),
	}

	log.Printf("Listening on %v", server.Addr)
	log.Fatal(server.ListenAndServe())
}
