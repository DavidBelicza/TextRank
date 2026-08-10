package parse

import (
	"strings"
	"unicode"
	"unicode/utf8"
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
		if (rule.IsSentenceSeparator(chr) && !isAbbreviationDot(rawText, chr, j)) ||
			j == slen {
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

func isAbbreviationDot(text string, chr rune, end int) bool {
	if chr != '.' {
		return false
	}

	next, _ := utf8.DecodeRuneInString(text[end:])

	if isSingleLetterWord(text[:end-1]) && unicode.IsLetter(next) {
		return true
	}

	previous, _ := utf8.DecodeLastRuneInString(text[:end-1])

	return unicode.IsDigit(previous) && unicode.IsDigit(next)
}

func isSingleLetterWord(text string) bool {
	last, size := utf8.DecodeLastRuneInString(text)

	if !unicode.IsLetter(last) {
		return false
	}

	if len(text) == size {
		return true
	}

	before, _ := utf8.DecodeLastRuneInString(text[:len(text)-size])

	return unicode.IsSpace(before) || before == '.'
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
		sep := rule.IsWordSeparator(chr) &&
			!isAbbreviationDot(rawSentence, chr, j)

		if sep || j == slen {
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
