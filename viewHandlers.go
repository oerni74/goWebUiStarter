package main

import (
	"net/http"
)

const indexTemplateName = "index.gohtml"

func (s *StarterServer) homeView(w http.ResponseWriter, req *http.Request) {
	err := s.Templates["home"].ExecuteTemplate(w, indexTemplateName, "Data String")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *StarterServer) aboutView(w http.ResponseWriter, req *http.Request) {
	err := s.Templates["about"].ExecuteTemplate(w, indexTemplateName, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
