package parse

// Text struct contains a parsed text.
type Text struct {
	parsedSentences []ParsedSentence
}

// ParsedSentence struct contains the original raw sentences and their words.
type ParsedSentence struct {
	original string
	words    []string
}

// Append method creates a sentence and its words and append them to the Text
// object.
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

// GetSentences method returns ParsedSentence slice from Text struct.
func (text *Text) GetSentences() []ParsedSentence {
	return text.parsedSentences
}

// GetWords methods returns the words string slice of ParsedSentence struct.
func (parsedSentence *ParsedSentence) GetWords() []string {
	return parsedSentence.words
}

// GetOriginal method returns the original sentence as a string from a
// ParsedSentence struct.
func (parsedSentence *ParsedSentence) GetOriginal() string {
	return parsedSentence.original
}
