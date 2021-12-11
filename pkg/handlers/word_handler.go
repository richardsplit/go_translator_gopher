package handlers

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/richardsplit/go_translator_gopher/pkg/model"
	"github.com/richardsplit/go_translator_gopher/pkg/server"
	"github.com/richardsplit/go_translator_gopher/pkg/translation"
)

type WordHandler struct{}

func NewWordHandler() *WordHandler {
	return &WordHandler{}
}

func (h *WordHandler) Handle(response server.Response, request server.Request, translator Translator, history History) error {

	word := model.InputWord{}

	if err := request.ReadJSON(&word); err != nil {
		if err := response.WriteJSON(http.StatusBadRequest, &errorResponse{
			Message: "please provide valid json input",
		}); err != nil {
			return errors.Wrap(err, "failed to write response")
		}
		return nil
	}

	if strings.TrimSpace(word.English) == "" {
		if err := response.WriteJSON(http.StatusBadRequest, &errorResponse{
			Message: "please provide non-empty word",
		}); err != nil {
			return errors.Wrap(err, "failed to write response")
		}
		return nil
	}

	gopherWord, err := translator.TranslateWord(word.English)
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

	history.Add(word.English, gopherWord)

	if err := response.WriteJSON(http.StatusOK, model.OutputWord{
		Gopher: gopherWord,
	}); err != nil {
		return errors.Wrap(err, "failed to write response")
	}
	return nil
}
