package main

import (
	"embed"
	"goWebUIStarter/templates"
	"html/template"
	"net/http"
)

//go:embed static
var static embed.FS

type StarterServer struct {
	Mux       *http.ServeMux
	Templates map[string]*template.Template
}

func NewStarterServer() *StarterServer {
	server := &StarterServer{
		Mux:       http.NewServeMux(),
		Templates: templates.ParseTemplates(),
	}

	server.registerRoutes()

	return server
}

func (s *StarterServer) registerRoutes() {
	// View Handler
	s.Mux.HandleFunc("GET /", s.homeView)
	s.Mux.HandleFunc("GET /about", s.aboutView)

	// API Handler
	s.Mux.HandleFunc("GET /api/check-number", s.checkNumber)

	// Static
	s.Mux.Handle("GET /static/*", http.FileServer(http.FS(static)))
}
