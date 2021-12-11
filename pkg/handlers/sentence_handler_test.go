package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/richardsplit/go_translator_gopher/pkg/handlers"
	"github.com/richardsplit/go_translator_gopher/pkg/handlers/handlers_mocks"
	"github.com/richardsplit/go_translator_gopher/pkg/model"
	"github.com/richardsplit/go_translator_gopher/pkg/translation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	var (
		controller *gomock.Controller
		translator *handlers_mocks.MockTranslator
		history    *handlers_mocks.MockHistory

		testServer *httptest.Server

		request *resty.Request

		expectedOutputPayload model.OutputSentence
		actualOutputPayload   model.OutputSentence
		body                  model.InputSentence
	)

	BeforeEach(func() {
		controller = gomock.NewController(GinkgoT())
		translator = handlers_mocks.NewMockTranslator(controller)
		history = handlers_mocks.NewMockHistory(controller)

		handler := handlers.NewSentenceHandler()
		router := createHandlerTestRouter("/", translator, history, handler.Handle)
		testServer = httptest.NewServer(router)

		body = model.InputSentence{
			English: "You either die a hero, or you live long enough to see yourself become the villain.",
		}

		expectedOutputPayload = model.OutputSentence{
			Gopher: "ouYogo geither iedogo ga erohogo, gor ouyogo ivelogo onglogo genough otogo eesogo ourselfyogo ecomebogo ethogo illainvogo.",
		}

		request = resty.New().
			SetHostURL(testServer.URL).R().
			SetResult(&actualOutputPayload)
	})

	AfterEach(func() {
		testServer.Close()
	})

	When("translation is successful", func() {
		BeforeEach(func() {
			request.SetBody(body)

			history.EXPECT().Add(body.English, expectedOutputPayload.Gopher).Return()
			translator.EXPECT().TranslateSentence(body.English).Return(expectedOutputPayload.Gopher, nil)
		})

		It("returns expected response", func() {
			response, errRequest := request.Post("")

			Expect(errRequest).ToNot(HaveOccurred())
			Expect(response.StatusCode()).To(Equal(http.StatusOK))
			Expect(actualOutputPayload).To(Equal(expectedOutputPayload))
		})
	})

	When("input is invalid", func() {
		Context("and proper error is returned", func() {
			BeforeEach(func() {
				request.SetBody(body)

				history.EXPECT().Add(gomock.Any(), gomock.Any()).Times(0)
				translator.EXPECT().TranslateSentence(body.English).Return(expectedOutputPayload.Gopher, translation.TranslationError{
					Cause: "expected error",
				})
			})

			It("returns Bad Request response", func() {
				response, errRequest := request.Post("")

				Expect(errRequest).ToNot(HaveOccurred())
				Expect(response.StatusCode()).To(Equal(http.StatusBadRequest))
			})
		})

		Context("and unexpected error is returned", func() {
			BeforeEach(func() {
				request.SetBody(body)

				history.EXPECT().Add(gomock.Any(), gomock.Any()).Times(0)
				translator.EXPECT().TranslateSentence(body.English).Return(expectedOutputPayload.Gopher, errors.New("test-error"))
			})

			It("returns Internal Server Error response", func() {
				response, errRequest := request.Post("")

				Expect(errRequest).ToNot(HaveOccurred())
				Expect(response.StatusCode()).To(Equal(http.StatusInternalServerError))
			})
		})
	})

	When("input is empty", func() {
		BeforeEach(func() {
			emptyBody := model.InputSentence{
				English: "",
			}
			request.SetBody(emptyBody)

			history.EXPECT().Add(gomock.Any(), gomock.Any()).Times(0)
			translator.EXPECT().TranslateSentence(gomock.Any()).Times(0)
		})

		It("returns Bad Request response", func() {
			response, errRequest := request.Post("")

			Expect(errRequest).ToNot(HaveOccurred())
			Expect(response.StatusCode()).To(Equal(http.StatusBadRequest))
		})
	})

	When("input is invlaid json", func() {
		BeforeEach(func() {
			request.SetBody("invalid JSON")

			history.EXPECT().Add(gomock.Any(), gomock.Any()).Times(0)
			translator.EXPECT().TranslateSentence(gomock.Any()).Times(0)
		})

		It("returns Bad Request response", func() {
			response, errRequest := request.Post("")

			Expect(errRequest).ToNot(HaveOccurred())
			Expect(response.StatusCode()).To(Equal(http.StatusBadRequest))
		})
	})
})
