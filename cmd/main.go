package main

import (
	"log"

	"plantlife/internal/server"

	_ "github.com/lib/pq"
)


func main() {

	server := server.NewServer()

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
