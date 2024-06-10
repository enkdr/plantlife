package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	var templatesPath string
	templatesPath = "templates/index.html"

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	NewServer := &Server{
		templates: template.Must(template.ParseFiles(templatesPath)),
		port:      config.PORT,
		db:        db,
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", NewServer.port),
		// register all routes
		Handler: NewServer.RegisterRoutes(),
	}

	return server

}
