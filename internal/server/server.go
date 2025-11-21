package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("route %v", r.URL)
		json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
	})

	return mux
}
