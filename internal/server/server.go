package server

import (
	"encoding/json"
	_"golang.org/x/oauth2"
	_"golang.org/x/oauth2/google"
	"log"
	"net/http"

	mw "github.com/jpleatherland/gwr-server/internal/middleware"
)

type Server struct {
	Addr string
	Config *mw.Resources // or whatever resources you need
	mux    *http.ServeMux
}

func NewServer(resources *mw.Resources) *Server {
	s := &Server{
		Addr: ":8080",
		Config: resources,
		mux:    http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("route %v", r.URL)
		json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
	})

	s.mux.HandleFunc("GET /auth/google", func(w http.ResponseWriter, r *http.Request) {

		log.Printf("route %v", r.URL)
		log.Printf("resources %v", s.Config.ClientId)
		// url:= s.Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	s.mux.ServeHTTP(w, r)
}
