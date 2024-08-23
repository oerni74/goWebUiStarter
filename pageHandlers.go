package main

import (
	"goWebUIStarter/templates"
	"net/http"
)

func (s *StarterServer) homeView(w http.ResponseWriter, req *http.Request) {
	err := s.Templates[templates.HomeTemplate].ExecuteTemplate(w, templates.IndexTemplateName, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *StarterServer) aboutView(w http.ResponseWriter, req *http.Request) {
	err := s.Templates[templates.AboutTemplate].ExecuteTemplate(w, templates.IndexTemplateName, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
