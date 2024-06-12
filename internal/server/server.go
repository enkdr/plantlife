package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"plantlife/config"
	"plantlife/internal/database"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	templates *template.Template
	port      int
	db        *sqlx.DB
}

func NewServer() *http.Server {

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Failed to retrieve configs:", err)
	}

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	currentDir, _ := os.Getwd()
	templatesPath := currentDir + config.TEMPLATE_PATH

	NewServer := &Server{
		templates: template.Must(template.ParseFiles(templatesPath)),
		port:      config.PORT,
		db:        db,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", NewServer.port),
		Handler: NewServer.RegisterRoutes(),
	}

	return server

}
