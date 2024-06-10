package database

import (
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
)

// called from server routes
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
