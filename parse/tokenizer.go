package parse

import (
	"strings"
)

var wordSeparators = [10]string{" ", ",", ")", "(", "[", "]", "{", "}", "\"", ";"}
var sentenceSeparators = [3]string{"!", ".", "?"}

func TokenizeText(rawText string) Text {
	return findSentences(rawText)
}

func findSentences(rawText string) Text {
	text := Text{}

	var sentence string

	for _, chr := range rawText {
		if !isSentenceSeparator(chr) {
			sentence = sentence + string(chr)
		} else if len(sentence) > 0 {
			sentence = sentence + string(chr)

			text.Append(sentence, findWords(sentence))

			sentence = ""
		}
	}

	if len(sentence) > 0 {
		text.Append(sentence, findWords(sentence))
	}

	return text
}

func findWords(rawSentence string) (words []string) {
	words = []string{}

	var word string

	for _, chr := range rawSentence {
		if !isWordSeparator(chr) {
			word = word + string(chr)
		} else if len(word) > 0 {
			words = append(words, strings.ToLower(word))
			word = ""
		}
	}

	if len(word) > 0 {
		words = append(words, strings.ToLower(word))
	}

	return
}

func isWordSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range wordSeparators {
		if chr == val {
			return true
		}
	}

	return isSentenceSeparator(rune)
}

func isSentenceSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range sentenceSeparators {
		if chr == val {
			return true
		}
	}

	return false
}
