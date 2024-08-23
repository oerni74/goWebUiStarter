package main

import (
	"goWebUIStarter/templates"
	"log/slog"
	"net/http"
	"strconv"
)

const secretNumber = 42

type guessResponse struct {
	IsCorrect bool
	Hint      string
	Guess     int
}

func (s *StarterServer) checkNumber(w http.ResponseWriter, req *http.Request) {
	var response guessResponse
	guessedNumber := req.FormValue("numberGuess")
	guess, _ := strconv.Atoi(guessedNumber)

	slog.Info("checkNumber called with guess ", guess)

	if guess > secretNumber {
		response = guessResponse{
			IsCorrect: false,
			Hint:      "Go lower",
			Guess:     guess,
		}
	} else if guess < secretNumber {
		response = guessResponse{
			IsCorrect: false,
			Hint:      "Go higher",
			Guess:     guess,
		}
	} else {
		response = guessResponse{
			IsCorrect: true,
			Hint:      "",
			Guess:     guess,
		}
	}

	err := s.Templates[templates.HomeTemplate].ExecuteTemplate(w, "guess-response", response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
