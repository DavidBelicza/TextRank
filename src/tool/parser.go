package tool

//@todo sentences
//@todo compare rune with rune instead of string mess
var wordSeparators = []string{" ", ","}

// Parse @todo
func Parse(text string) (words []string) {
	words = []string{}

	var word string

	for _, chr := range text {
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

	return false
}
