package handlers

import (
	"net/http"

	"github.com/richardsplit/go_translator_gopher/pkg/server"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := server.Response{
		ResponseWriter: w,
	}

	response.WriteJSON(http.StatusOK, homeResponse{
		Message:  "please use one of the following endpoints",
		History:  "/history",
		Word:     "/word",
		Sentence: "/sentence",
	})
}

type homeResponse struct {
	Message  string `json:"message"`
	History  string `json:"history"`
	Word     string `json:"word"`
	Sentence string `json:"sentence"`
}
