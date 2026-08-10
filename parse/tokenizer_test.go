package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizeText(t *testing.T) {
	rule := NewRule()

	text := TokenizeText(
		"This is the right sentence. This sentence without end mark",
		rule,
	)

	assert.Equal(t,
		" This sentence without end mark",
		text.parsedSentences[1].original,
	)
}

func TestTokenizeTextWithAbbreviation(t *testing.T) {
	rule := NewRule()

	text := TokenizeText("The U.S.A is big. It is far.", rule)

	assert.Equal(t, 2, len(text.parsedSentences))
	assert.Equal(t, "The U.S.A is big.", text.parsedSentences[0].original)
	assert.Equal(t, " It is far.", text.parsedSentences[1].original)
}

func TestTokenizeTextKeepsAbbreviationAsOneWord(t *testing.T) {
	rule := NewRule()

	text := TokenizeText("The U.S.A is big.", rule)

	assert.Equal(t,
		[]string{"the", "u.s.a", "is", "big"},
		text.parsedSentences[0].words,
	)
}

func TestTokenizeTextWithDecimalNumber(t *testing.T) {
	rule := NewRule()

	text := TokenizeText("Version 3.14 shipped. It works.", rule)

	assert.Equal(t, 2, len(text.parsedSentences))
	assert.Equal(t, "Version 3.14 shipped.", text.parsedSentences[0].original)
}

func TestTokenizeTextSplitsAfterPossessive(t *testing.T) {
	rule := NewRule()

	text := TokenizeText("It was Vasili's. The elderly man left.", rule)

	assert.Equal(t, 2, len(text.parsedSentences))
	assert.Equal(t, "It was Vasili's.", text.parsedSentences[0].original)
}

func TestTokenizeTextSplitsRegularSentences(t *testing.T) {
	rule := NewRule()

	text := TokenizeText("I like it. He came. She left!", rule)

	assert.Equal(t, 3, len(text.parsedSentences))
	assert.Equal(t, "I like it.", text.parsedSentences[0].original)
	assert.Equal(t, " He came.", text.parsedSentences[1].original)
	assert.Equal(t, " She left!", text.parsedSentences[2].original)
}

func TestTokenizeTextWithAbbreviationAtTheStart(t *testing.T) {
	rule := NewRule()

	text := TokenizeText("U.S.A is big.", rule)

	assert.Equal(t, 1, len(text.parsedSentences))
	assert.Equal(t, []string{"u.s.a", "is", "big"}, text.parsedSentences[0].words)
}
