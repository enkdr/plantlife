package server

import (
	"fmt"
	"log"
	"net/http"
	"plantlife/config"
	"plantlife/internal/database"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   *sqlx.DB
}

func NewServer() *http.Server {

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Failed to retrieve configs:", err)
	}

	NewServer := &Server{
		port: config.PORT,
		db:   db,
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", NewServer.port),
		// register all routes
		Handler: NewServer.RegisterRoutes(),
	}

	return server

}
