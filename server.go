package main

import (
	"goWebUIStarter/templates"
	"html/template"
	"net/http"
)

type StarterServer struct {
	Mux       *http.ServeMux
	Templates map[string]*template.Template
}

func NewStarterServer() *StarterServer {
	server := &StarterServer{Mux: http.NewServeMux(), Templates: templates.ParseTemplates()}

	server.registerRoutes()

	return server
}

func (s *StarterServer) registerRoutes() {
	s.Mux.HandleFunc("GET /", s.homeView)
	s.Mux.HandleFunc("GET /about", s.aboutView)
}
