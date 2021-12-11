package handlers

import (
	"net/http"

	"github.com/richardsplit/translator_go/pkg/server"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen --source=handler_wrapper.go --destination handlers_mocks/handler_wrapper.go --package handlers_mocks

type Translator interface {
	TranslateWord(original string) (string, error)
	TranslateSentence(original string) (string, error)
}

type DefaultTranslatorHandler interface {
	Handle(response server.Response, request server.Request, translator Translator, history History) error
}

type DefaultHandler interface {
	Handle(response server.Response, request server.Request, history History) error
}

type History interface {
	Add(key, value string)
	GetArranged() []map[string]string
}

func TranslatorHandlerWrapper(w http.ResponseWriter, r *http.Request, translator Translator, defaultHandler DefaultTranslatorHandler, history History) {
	request := server.Request{
		Request: r,
	}

	response := server.Response{
		ResponseWriter: w,
	}

	errHandle := defaultHandler.Handle(response, request, translator, history)
	if errHandle != nil {
		errWriteInternalServerError := response.WriteJSON(http.StatusInternalServerError, &errorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		if errWriteInternalServerError != nil {
			logrus.Warn("unable to write response")
		}
	}
}

func HandlerWrapper(w http.ResponseWriter, r *http.Request, defaultHandler DefaultHandler, history History) {
	request := server.Request{
		Request: r,
	}

	response := server.Response{
		ResponseWriter: w,
	}

	errHandle := defaultHandler.Handle(response, request, history)
	if errHandle != nil {
		errWriteInternalServerError := response.WriteJSON(http.StatusInternalServerError, &errorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		if errWriteInternalServerError != nil {
			logrus.Warn("unable to write response")
		}
	}
}

func TranslatorHandlerWrapperFunc(translator Translator, defaultHandler DefaultTranslatorHandler, history History) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		TranslatorHandlerWrapper(rw, r, translator, defaultHandler, history)
	}
}
func HandlerWrapperFunc(defaultHandler DefaultHandler, history History) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		HandlerWrapper(rw, r, defaultHandler, history)
	}
}

type errorResponse struct {
	Message string `json:"message"`
}
