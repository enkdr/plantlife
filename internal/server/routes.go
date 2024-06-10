package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.IndexHandler)
	mux.HandleFunc("/home", s.HomeHandler)
	mux.HandleFunc("/plants", s.PlantHandler)
	return mux
}
