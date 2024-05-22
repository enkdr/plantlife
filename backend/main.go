package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

var db *sql.DB

type Plant struct {
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
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

}

func main() {

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	var err error
	for retries := 0; retries < 10; retries++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			break
		}
		log.Printf("Error connecting to database: %v. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after multiple retries: %v", err)
	} else {
		fmt.Println("postgres connected")
	}

	http.HandleFunc("/", handler)

	fmt.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
