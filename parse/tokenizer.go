package parse

import (
	"strings"
)

// TokenizeText function use the given raw text and parses by a Rule object and
// retrieves the parsed text in a Text struct object.
func TokenizeText(rawText string, rule Rule) Text {
	return findSentences(rawText, rule)
}

func findSentences(rawText string, rule Rule) Text {
	text := Text{}

	var sentence string

	for _, chr := range rawText {
		if !rule.IsSentenceSeparator(chr) {
			sentence = sentence + string(chr)
		} else if len(sentence) > 0 {
			sentence = sentence + string(chr)

			text.Append(sentence, findWords(sentence, rule))

			sentence = ""
		}
	}

	if len(sentence) > 0 {
		text.Append(sentence, findWords(sentence, rule))
	}

	return text
}

func findWords(rawSentence string, rule Rule) (words []string) {
	words = []string{}

	var word string

	for _, chr := range rawSentence {
		if !rule.IsWordSeparator(chr) {
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
