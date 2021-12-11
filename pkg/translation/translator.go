package translation

import (
	"fmt"
	"strings"
)

type TranslationError struct {
	Cause string
}

func (e TranslationError) Error() string {
	return fmt.Sprintf("invalid input: %s", e.Cause)
}

func (e TranslationError) UserMessage() string {
	return fmt.Sprintf("provided input is invalid [%s]. please check and try again.", e.Cause)
}

type Translator struct {
}

func NewTranslator() *Translator {
	return &Translator{}
}

var vowels = []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}

// consonants sounds list could be extended.
var consonantSounds = []string{"th", "Th", "sh", "Sh", "ch", "Ch", "p", "P", "b", "B", "t", "T", "d", "D", "c", "C", "k", "K", "g", "G", "f", "F", "v", "V", "s", "S", "z", "Z", "h", "H", "j", "J", "m", "M", "n", "N", "l", "L", "r", "R", "v", "V", "w", "W", "y", "Y"}

func (t *Translator) TranslateWord(original string) (string, error) {

	original = strings.ReplaceAll(original, "'", "")

	if contains(vowels, []rune(original)[0]) {
		return "g" + original, nil
	}

	if strings.HasPrefix(original, "xr") {
		return "ge" + original, nil
	}

	for _, sound := range consonantSounds {
		if strings.HasPrefix(original, sound) {

			trimmed := strings.TrimPrefix(original, sound)

			if strings.HasPrefix(trimmed, "qu") {
				return strings.TrimPrefix(trimmed, "qu") + sound + "qu" + "ogo", nil
			}

			return trimmed + sound + "ogo", nil
		}
	}

	return "", TranslationError{
		Cause: "invalid word",
	}
}

// punctuations list could be extended.
var allPunctuations = []string{",", ":", ";", "-", ".", "?", "!"}
var endPunctuations = []string{".", "?", "!"}

func (t *Translator) TranslateSentence(original string) (string, error) {

	endsProperly := false
	for _, punc := range endPunctuations {
		if strings.HasSuffix(original, punc) {
			endsProperly = true
		}
	}

	if !endsProperly {
		return "", TranslationError{
			Cause: "sentence does not end in a punctuation marks",
		}
	}

	words := strings.Fields(original)
	result := ""

	for _, word := range words {

		currentWordPunc := ""
		for _, punc := range allPunctuations {
			if strings.HasSuffix(word, punc) {
				currentWordPunc = punc
			}
		}

		var translated string
		var err error

		if currentWordPunc != "" {
			translated, err = t.TranslateWord(strings.TrimSuffix(word, currentWordPunc))

		} else {
			translated, err = t.TranslateWord(word)
		}

		if err != nil {
			return "", TranslationError{
				Cause: "sentence contains invalid word",
			}
		}
		if len(result) > 0 {
			result = result + " "
		}
		result = result + translated + currentWordPunc
	}

	return result, nil
}

func contains(list []rune, r rune) bool {
	for _, a := range list {
		if a == r {
			return true
		}
	}
	return false
}
