package main

import (
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
	slog.Info("Check Number called")
	var response guessResponse
	guessedNumber := req.PostFormValue("numberGuess")
	guess, _ := strconv.Atoi(guessedNumber)
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

	err := s.Templates["home"].ExecuteTemplate(w, indexTemplateName, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
