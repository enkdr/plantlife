package server

import (
	"fmt"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.IndexHandler)
	// sm.HandleFunc("/home", s.HomeHandler)
	return mux

}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index handler")
}
