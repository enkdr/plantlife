package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.IndexHandler)
	mux.HandleFunc("/home", s.HomeHandler)
	mux.HandleFunc("/plants", s.PlantHandler)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fs))

	return mux
}
