package server_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/richardsplit/translator_go/pkg/server"
)

var _ = Describe("Response", func() {
	const statusCode = http.StatusTeapot

	var fakeResponse *httptest.ResponseRecorder

	BeforeEach(func() {
		fakeResponse = httptest.NewRecorder()
	})

	Describe("WriteJSON", func() {
		const validJsonString = "{\"key\":\"value\"}"

		It("writes json successfully", func() {
			object := struct {
				Key string `json:"key"`
			}{
				Key: "value",
			}
			response := server.Response{
				ResponseWriter: fakeResponse,
			}
			if err := response.WriteJSON(statusCode, object); err != nil {
				Expect(err).ToNot(HaveOccurred())
			}

			responseResult := fakeResponse.Result()
			body, err := ioutil.ReadAll(responseResult.Body)
			Expect(err).ToNot(HaveOccurred())

			Expect(responseResult.StatusCode).To(Equal(statusCode))
			Expect(responseResult.Header.Get("Content-Type")).To(Equal("application/json"))
			Expect(strings.TrimSpace(string(body))).To(Equal(validJsonString))
		})
	})

	Describe("WriteOctet", func() {
		It("writes octet-stream successfully", func() {
			response := server.Response{
				ResponseWriter: fakeResponse,
			}
			if err := response.WriteOctet(statusCode, []byte("test")); err != nil {
				Expect(err).ToNot(HaveOccurred())
			}

			responseResult := fakeResponse.Result()
			body, err := ioutil.ReadAll(responseResult.Body)
			Expect(err).ToNot(HaveOccurred())

			Expect(responseResult.StatusCode).To(Equal(statusCode))
			Expect(responseResult.Header.Get("Content-Type")).To(Equal("application/octet-stream"))
			Expect(strings.TrimSpace(string(body))).To(Equal("test"))
		})
	})
})
