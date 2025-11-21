package main

import (
	"github.com/jpleatherland/gwr-server/server/server"
	"net/http"
)

func main() {

	http.HandleFunc("/", server.handler)
	return
}
