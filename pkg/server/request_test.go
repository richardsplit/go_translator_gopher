package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/richardsplit/go_translator_gopher/pkg/server"
)

var _ = Describe("Request", func() {

	var fakeRequest *http.Request

	Describe("ReadJSON", func() {
		const validJsonString = "{\"key\":\"value\"}"

		type Object struct {
			Key string `json:"key"`
		}

		var receivedObject Object
		var errRequest error

		BeforeEach(func() {
			fakeRequest = httptest.NewRequest(http.MethodGet, "http://example.org", strings.NewReader(validJsonString))
		})

		JustBeforeEach(func() {
			request := server.Request{Request: fakeRequest}
			errRequest = request.ReadJSON(&receivedObject)
		})

		It("reads json successfully", func() {
			Expect(errRequest).ToNot(HaveOccurred())
			Expect(receivedObject).To(Equal(Object{Key: "value"}))
		})

		When("json in body is invalid", func() {
			const invalidJsonString = "invalid-json"
			BeforeEach(func() {
				fakeRequest = httptest.NewRequest(http.MethodGet, "http://example.org", strings.NewReader(invalidJsonString))
			})

			It("fails", func() {
				Expect(errRequest).To(HaveOccurred())
			})
		})
	})
})
