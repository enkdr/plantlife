package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Plant struct {
	Name string `json:"name"`
}

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

	rows, err := s.db.Query("SELECT name from plant")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var plants []Plant
	for rows.Next() {
		var plant Plant
		if err := rows.Scan(&plant.Name); err != nil {
			log.Println(err)
			continue
		}
		plants = append(plants, plant)
	}

	plantsJSON, err := json.Marshal(plants)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(plantsJSON)
}
