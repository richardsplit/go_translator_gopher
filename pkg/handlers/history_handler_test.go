package handlers_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/richardsplit/translator_go/pkg/handlers/handlers_mocks"
	"github.com/richardsplit/translator_go/pkg/model"
	"github.com/richardsplit/translator_go/pkg/server"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/richardsplit/translator_go/pkg/handlers"
)

var _ = Describe("Handler", func() {
	var (
		controller *gomock.Controller
		history    *handlers_mocks.MockHistory

		testServer *httptest.Server

		request *resty.Request

		expectedOutputPayload model.History
		actualOutputPayload   model.History
	)

	BeforeEach(func() {
		controller = gomock.NewController(GinkgoT())
		history = handlers_mocks.NewMockHistory(controller)

		handler := NewHistoryHandler()
		router := createHandlerHistoryRouter("/", history, handler.Handle)
		testServer = httptest.NewServer(router)

		expectedOutputPayload = model.History{
			History: []map[string]string{
				{"antman": "wasp"},
				{"batman": "robin"},
				{"Ironman": "Pepper Potts"},
			},
		}

		request = resty.New().
			SetHostURL(testServer.URL).R().
			SetResult(&actualOutputPayload)

		history.EXPECT().GetArranged().Return(expectedOutputPayload.History)
	})

	AfterEach(func() {
		testServer.Close()
	})

	It("returns expected response", func() {

		response, errRequest := request.Post("")

		Expect(errRequest).ToNot(HaveOccurred())
		Expect(response.StatusCode()).To(Equal(http.StatusOK))
		Expect(actualOutputPayload).To(Equal(expectedOutputPayload))

	})

})

func createHandlerHistoryRouter(path string, history History, handleFunc func(response server.Response, request server.Request, history History) error) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
		request := server.Request{
			Request: r,
		}

		response := server.Response{
			ResponseWriter: rw,
		}
		handleFunc(response, request, history)
	}).
		Methods(http.MethodPost)
	return router
}
