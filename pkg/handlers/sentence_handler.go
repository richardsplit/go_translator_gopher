package handlers

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/richardsplit/go_translator_gopher/pkg/model"
	"github.com/richardsplit/go_translator_gopher/pkg/server"
	"github.com/richardsplit/go_translator_gopher/pkg/translation"
)

type SentenceHandler struct{}

func NewSentenceHandler() *SentenceHandler {
	return &SentenceHandler{}
}

func (h *SentenceHandler) Handle(response server.Response, request server.Request, translator Translator, history History) error {

	sentence := model.InputSentence{}

	if err := request.ReadJSON(&sentence); err != nil {
		if err := response.WriteJSON(http.StatusBadRequest, &errorResponse{
			Message: "please provide valid json input",
		}); err != nil {
			return errors.Wrap(err, "failed to write response")
		}
		return nil
	}

	if strings.TrimSpace(sentence.English) == "" {
		if err := response.WriteJSON(http.StatusBadRequest, &errorResponse{
			Message: "please provide non-empty input",
		}); err != nil {
			return errors.Wrap(err, "failed to write response")
		}
		return nil
	}

	gopherSentence, err := translator.TranslateSentence(sentence.English)

	if err != nil {
		translationErr, ok := err.(translation.TranslationError)
		if !ok {
			if err := response.WriteJSON(http.StatusInternalServerError, &errorResponse{
				Message: http.StatusText(http.StatusInternalServerError),
			}); err != nil {
				return errors.Wrap(err, "failed to write response")
			}
			return nil
		}

		if err := response.WriteJSON(http.StatusBadRequest, &errorResponse{
			Message: translationErr.UserMessage(),
		}); err != nil {
			return errors.Wrap(err, "failed to write response")
		}
		return nil
	}

	history.Add(sentence.English, gopherSentence)

	if err := response.WriteJSON(http.StatusOK, model.OutputSentence{
		Gopher: gopherSentence,
	}); err != nil {
		return errors.Wrap(err, "failed to write response")
	}
	return nil
}
