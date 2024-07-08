package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

var db *sql.DB

type Plant struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var err error

	db, err := sql.Open("sqlite3", "./db/plants.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, description FROM plant")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var plants []Plant
	for rows.Next() {
		var plant Plant
		if err := rows.Scan(&plant.Name, &plant.Description); err != nil {
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

	fmt.Println("OK")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow requests from React app
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handlerWithCors := c.Handler(mux)

	fmt.Println("Serving on localhost:3000")
	err := http.ListenAndServe(":3000", handlerWithCors)
	if err != nil {
		log.Fatal(err)
	}
}
