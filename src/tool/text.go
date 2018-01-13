package tool

// Text The container of the parsed sentences.
type Text struct {
	parsedSentences []ParsedSentence
}

// ParsedSentence Original sentence and its parsed words.
type ParsedSentence struct {
	original string
	words    []string
}

// Append It creates a Sentence and Words and append it to the Text.
func (text *Text) Append(rawSentence string, words []string) {
	if len(words) > 0 {
		parsedSentence := ParsedSentence{
			original: rawSentence,
			words:    words,
		}

		text.parsedSentences = append(
			text.parsedSentences,
			parsedSentence,
		)
	}
}

// GetSentences returns ParsedSentence slice from Text struct.
func (text *Text) GetSentences() []ParsedSentence {
	return text.parsedSentences
}

// GetWords returns the words string slice of ParsedSentence.
func (parsedSentence *ParsedSentence) GetWords() []string {
	return parsedSentence.words
}

// GetOriginal return the original sentence as a string from a ParsedSentence.
func (parsedSentence *ParsedSentence) GetOriginal() string {
	return parsedSentence.original
}
