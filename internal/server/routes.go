package server

import (
	"fmt"
	"net/http"
	"plantlife/internal/database"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.IndexHandler)
	mux.HandleFunc("/home", s.HomeHandler)
	mux.HandleFunc("/plant", s.PlantHandler)
	return mux
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index handler")
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home handler")
}

func (s *Server) PlantHandler(w http.ResponseWriter, r *http.Request) {

	plantsJSON := database.PlantQuery(s.db)

	w.Header().Set("Content-Type", "application/json")
	w.Write(plantsJSON)
}
