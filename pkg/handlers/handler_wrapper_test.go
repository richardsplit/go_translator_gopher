package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/richardsplit/go_translator_gopher/pkg/handlers/handlers_mocks"
	"github.com/richardsplit/go_translator_gopher/pkg/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/richardsplit/go_translator_gopher/pkg/handlers"
)

var _ = Describe("HandlerWrapper", func() {
	var (
		controller     *gomock.Controller
		translator     *handlers_mocks.MockTranslator
		defaultHandler *handlers_mocks.MockDefaultTranslatorHandler
		history        *handlers_mocks.MockHistory

		testServer *httptest.Server

		request *resty.Request

		expectedOutputWord *model.OutputWord
		actualOutputWord   *model.OutputWord
		body               *model.InputWord
	)

	BeforeEach(func() {

		controller = gomock.NewController(GinkgoT())
		translator = handlers_mocks.NewMockTranslator(controller)
		defaultHandler = handlers_mocks.NewMockDefaultTranslatorHandler(controller)
		history = handlers_mocks.NewMockHistory(controller)

		body = &model.InputWord{
			English: "penguin",
		}

		expectedOutputWord = &model.OutputWord{
			Gopher: "enguinpogo",
		}

	})

	Context("HandlerWrapper", func() {
		BeforeEach(func() {
			router := createHandlerWrapperTestRouter("/", translator, history, defaultHandler, TranslatorHandlerWrapper)
			testServer = httptest.NewServer(router)

			request = resty.New().
				SetHostURL(testServer.URL).R().
				SetResult(&actualOutputWord)
		})

		AfterEach(func() {
			testServer.Close()
		})

		When("handler is successful", func() {
			BeforeEach(func() {

				request.SetBody(body)
				translator.EXPECT().TranslateWord(body.English).Return(expectedOutputWord.Gopher, nil)

				defaultHandler.EXPECT().Handle(gomock.Any(), gomock.Any(), translator, history).Return(nil)
				history.EXPECT().Add(body.English, "adii").Return()
			})

			It("returns expected response", func() {
				response, errRequest := request.Post("")

				Expect(errRequest).ToNot(HaveOccurred())
				Expect(response.StatusCode()).To(Equal(http.StatusOK))
			})
		})

		When("handler fails", func() {
			BeforeEach(func() {
				request.SetBody(body)
				translator.EXPECT().TranslateWord(body.English).Return(expectedOutputWord.Gopher, nil)

				defaultHandler.EXPECT().Handle(gomock.Any(), gomock.Any(), translator, history).Return(errors.New("test-error"))
			})

			It("returns Internal Server Error response", func() {
				response, errRequest := request.Post("")

				Expect(errRequest).ToNot(HaveOccurred())
				Expect(response.StatusCode()).To(Equal(http.StatusInternalServerError))
			})
		})

	})

	Context("HandlerWrapperFunc", func() {

		BeforeEach(func() {
			router := createHandlerWrapperFuncTestRouter("/", translator, history, defaultHandler, TranslatorHandlerWrapperFunc)
			testServer = httptest.NewServer(router)

			request = resty.New().
				SetHostURL(testServer.URL).R().
				SetResult(&actualOutputWord)
		})

		AfterEach(func() {
			testServer.Close()
		})

		When("handler is successful", func() {
			BeforeEach(func() {
				request.SetBody(body)
				translator.EXPECT().TranslateWord(body.English).Return(expectedOutputWord.Gopher, nil)

				defaultHandler.EXPECT().Handle(gomock.Any(), gomock.Any(), translator, history).Return(nil)
			})

			It("returns expected response", func() {
				response, errRequest := request.Post("")

				Expect(errRequest).ToNot(HaveOccurred())
				Expect(response.StatusCode()).To(Equal(http.StatusOK))
			})
		})

		When("handler fails", func() {
			BeforeEach(func() {
				request.SetBody(body)
				translator.EXPECT().TranslateWord(body.English).Return(expectedOutputWord.Gopher, nil)

				defaultHandler.EXPECT().Handle(gomock.Any(), gomock.Any(), translator, history).Return(errors.New("test-error"))
			})

			It("returns Internal Server Error response", func() {
				response, errRequest := request.Post("")

				Expect(errRequest).ToNot(HaveOccurred())
				Expect(response.StatusCode()).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})

func createHandlerWrapperTestRouter(path string, translator Translator, history History, defaultHandler DefaultTranslatorHandler, handleFunc func(w http.ResponseWriter, r *http.Request, translator Translator, defaultHandler DefaultTranslatorHandler, history History)) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {

		handleFunc(rw, r, translator, defaultHandler, history)
	}).
		Methods(http.MethodPost)
	return router
}

func createHandlerWrapperFuncTestRouter(path string, translator Translator, history History, defaultHandler DefaultTranslatorHandler, handleFunc func(translator Translator, defaultHandler DefaultTranslatorHandler, history History) func(rw http.ResponseWriter, r *http.Request)) *mux.Router {
	router := mux.NewRouter()
	router.
		HandleFunc(path, handleFunc(translator, defaultHandler, history)).
		Methods(http.MethodPost)
	return router
}
