package translation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/richardsplit/go_translator_gopher/pkg/translation"
)

var _ = Describe("Translator", func() {

	var translator *Translator

	BeforeEach(func() {
		translator = NewTranslator()
	})

	Context("Tranlate", func() {

		It("adds 'g' in front of vowel", func() {
			Expect(translator.TranslateWord("apple")).To(Equal("gapple"))
		})

		When("a word starts with consonant", func() {

			Context("and first two letters are is 'xr'", func() {
				It("adds 'ge' in front of 'xr'", func() {
					Expect(translator.TranslateWord("xray")).To(Equal("gexray"))
				})
			})

			It("moves consonant sound (one letter) to end of word and adds 'ogo' suffix", func() {
				Expect(translator.TranslateWord("yellow")).To(Equal("ellowyogo"))
			})

			It("moves consonant sound (multiple letters) to end of word and adds 'ogo' suffix", func() {
				Expect(translator.TranslateWord("chair")).To(Equal("airchogo"))
			})

			Context("and is followed by 'qu'", func() {
				It("moves consonant sound and 'qu' to end of word and adds 'ogo' suffix", func() {
					Expect(translator.TranslateWord("square")).To(Equal("aresquogo"))
				})
			})

			It("removes multiple apostrophes", func() {
				Expect(translator.TranslateWord("yell'ow's")).To(Equal("ellowsyogo"))
			})

		})

	})

	Context("TranslateSentence", func() {

		When("sentence does not end in punctuation mark", func() {
			It("returns error", func() {
				_, err := translator.TranslateSentence("I'm Batman")
				Expect(err).To(HaveOccurred())
			})
		})

		It("translates all words in a sentence", func() {
			result, err := translator.TranslateSentence("You either die a hero, or you live long enough to see yourself become the villain.")
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("ouYogo geither iedogo ga erohogo, gor ouyogo ivelogo onglogo genough otogo eesogo ourselfyogo ecomebogo ethogo illainvogo."))
		})

	})
})
