package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/jpleatherland/gwr-server/internal/middleware"
	"github.com/jpleatherland/gwr-server/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	oauthClientId := os.Getenv("OAUTH_CLIENT_ID")
	oauthClientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	oauthRedirectUrl := os.Getenv("OAUTH_REDIRECT_URI")

	resources := middleware.Resources{
		ClientId:          oauthClientId,
		ClientSecret:      oauthClientSecret,
		ClientRedirectURI: oauthRedirectUrl,
	}

	srv := server.NewServer(&resources)

	log.Printf("Listening on %v", srv.Addr)
	log.Fatal(http.ListenAndServe(":8080", srv))
}
