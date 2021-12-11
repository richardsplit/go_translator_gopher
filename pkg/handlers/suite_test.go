package handlers_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/richardsplit/translator_go/pkg/server"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/richardsplit/translator_go/pkg/handlers"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}

func createHandlerTestRouter(path string, translator Translator, history History, handleFunc func(response server.Response, request server.Request, translator Translator, history History) error) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
		request := server.Request{
			Request: r,
		}

		response := server.Response{
			ResponseWriter: rw,
		}
		handleFunc(response, request, translator, history)
	}).
		Methods(http.MethodPost)
	return router
}
