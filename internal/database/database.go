package database

import (
	"encoding/json"
	"fmt"
	"log"
	"plantlife/config"

	"github.com/jmoiron/sqlx"
)

func InitDB() (*sqlx.DB, error) {
	var err error

	config, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	connStr := fmt.Sprintf("port=%s host=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DB_PORT, config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to the database")

	return db, nil

}

type Plant struct {
	Name string `json:"name"`
}

func PlantQuery(db *sqlx.DB) []byte {

	rows, err := db.Query("SELECT name from plant")
	if err != nil {
		return []byte("error connecting to the database")
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
		return []byte("error connecting to the database")
	}

	return plantsJSON

}
