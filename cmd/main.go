package main

import (
	"log"
	"net/http"

	"plantlife/internal/server"

	_ "github.com/lib/pq"
)

func main() {

	server := server.NewServer()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
