package server

import (
	"fmt"
	"net/http"
	"plantlife/internal/database"
)

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index handler")
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home handler")

	err := s.templates.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (s *Server) PlantHandler(w http.ResponseWriter, r *http.Request) {

	plantsJSON := database.PlantQuery(s.db)

	w.Header().Set("Content-Type", "application/json")
	w.Write(plantsJSON)
}
