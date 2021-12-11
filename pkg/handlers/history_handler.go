package handlers

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/richardsplit/go_translator_gopher/pkg/model"
	"github.com/richardsplit/go_translator_gopher/pkg/server"
)

type HistoryHandler struct{}

func NewHistoryHandler() *HistoryHandler {
	return &HistoryHandler{}
}

func (h *HistoryHandler) Handle(response server.Response, request server.Request, history History) error {

	if err := response.WriteJSON(http.StatusOK, model.History{
		History: history.GetArranged(),
	}); err != nil {
		return errors.Wrap(err, "failed to write response")
	}
	return nil
}
