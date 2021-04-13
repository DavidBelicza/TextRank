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
	var i int
	slen := len(rawText)

	for j, chr := range rawText {
		j += len(string(chr))
		//when separator or the last
		if rule.IsSentenceSeparator(chr) || j == slen {
			sentence = rawText[i:j]
			if len(sentence) > 0 {
				text.Append(sentence, findWords(sentence, rule))
			}

			sentence = ""
			i = j
		}
	}

	return text
}

func findWords(rawSentence string, rule Rule) (words []string) {
	words = []string{}

	var word string
	var i int
	slen := len(rawSentence)

	for j, chr := range rawSentence {
		chrlen := len(string(chr))
		j += chrlen
		//when separator or the last
		if sep := rule.IsWordSeparator(chr); sep || j == slen {
			if sep {
				word = rawSentence[i : j-chrlen]
			} else {
				word = rawSentence[i:j]
			}
			if len(word) > 0 {
				words = append(words, strings.ToLower(word))
			}
			word = ""
			i = j
		}
	}

	return
}
