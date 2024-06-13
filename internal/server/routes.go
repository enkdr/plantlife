package server

import (
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

// CustomFileServer serves static files and sets the correct MIME types
func CustomFileServer(root http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the file extension
		ext := filepath.Ext(r.URL.Path)

		// Determine the MIME type
		mimeType := mime.TypeByExtension(ext)
		if mimeType != "" {
			w.Header().Set("Content-Type", mimeType)
		} else if strings.HasSuffix(r.URL.Path, "/") {
			// Default to text/html for directory requests
			w.Header().Set("Content-Type", "text/html")
		}

		// Serve the file
		http.FileServer(root).ServeHTTP(w, r)
	})
}

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.IndexHandler)
	mux.HandleFunc("/home", s.HomeHandler)
	mux.HandleFunc("/plants", s.PlantHandler)

	mux.Handle("/assets/", http.StripPrefix("/assets/", CustomFileServer(http.Dir("./plantlife-fe/dist/assets"))))
	return mux
}
