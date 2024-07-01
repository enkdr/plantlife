package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Plant struct {
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	var err error

	db, err := sql.Open("sqlite3", "./db/plantlife.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name from plant")
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

	fmt.Println("OK")

}

func main() {

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
