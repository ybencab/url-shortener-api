package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/ybencab/url-shortener/storage"
)

type Server struct {
	listenAddr string
	store      storage.Storer
}

func NewServer(listenAddr string, store storage.Storer) *Server {
	return &Server{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/url", s.handleGetURLs)
	router.HandleFunc("/url/create", s.handleCreateURLs)
	router.HandleFunc("/url/{slug}", s.handleGetURLBySlug)

  corsOpts := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"},
    AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
    AllowedHeaders:   []string{"Content-Type"},
  })
  handler := corsOpts.Handler(router)

	return http.ListenAndServe(s.listenAddr, handler)
}
