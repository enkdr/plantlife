package server

import (
	"fmt"
	"log"
	"net/http"
	"plantlife/config"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	// db   database.Service
}

func NewServer() *http.Server {

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Failed to retrieve configs:", err)
	}

	NewServer := &Server{
		port: config.PORT,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", NewServer.port),
		Handler: NewServer.RegisterRoutes(),
	}

	return server

}
