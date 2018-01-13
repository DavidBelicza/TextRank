package tool

//@todo compare rune with rune instead of string mess
var wordSeparators = [2]string{" ", ","}
var sentenceSeparators = [3]string{"!", ".", "?"}

// Parse Parsing a raw text into a Text struct.
func Parse(rawText string) Text {
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
			words = append(words, word)
			word = ""
		}
	}

	if len(word) > 0 {
		words = append(words, word)
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
