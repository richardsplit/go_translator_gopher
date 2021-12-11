package env_test

import (
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/richardsplit/translator_go/pkg/env"
)

var _ = Describe("Env", func() {

	const envName = "ENV_NAME"

	setEnv := func(value string) {
		Expect(os.Setenv(envName, value)).To(Succeed())
	}

	unsetEnv := func() {
		Expect(os.Unsetenv(envName)).To(Succeed())
	}

	Context("OptionalString", func() {
		const (
			defaultValue = "default"
			validString  = "value"
		)

		When("environment variable is set", func() {
			BeforeEach(func() {
				setEnv(validString)
			})

			It("extracts value successfully", func() {
				Expect(env.OptionalString(envName, defaultValue)).To(Equal(validString))
			})
		})

		When("environment variable is not set", func() {
			BeforeEach(func() {
				unsetEnv()
			})

			It("returns default value", func() {
				Expect(env.OptionalString(envName, defaultValue)).To(Equal(defaultValue))
			})
		})
	})

	Describe("OptionalInt", func() {
		const (
			defaultValue = 42
			validInt     = 17
		)

		When("environment variable is set", func() {
			BeforeEach(func() {
				setEnv(strconv.Itoa(validInt))
			})

			It("extracts value successfully", func() {
				result, err := env.OptionalInt(envName, defaultValue)
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(validInt))
			})
		})

		When("environment variable is not set", func() {
			BeforeEach(func() {
				unsetEnv()
			})

			It("returns default value", func() {
				result, err := env.OptionalInt(envName, defaultValue)
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(defaultValue))
			})
		})

		When("environment variable is not a number", func() {
			BeforeEach(func() {
				setEnv("ABCD")
			})

			It("returns an error", func() {
				_, err := env.OptionalInt(envName, defaultValue)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
