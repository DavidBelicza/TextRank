package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	parsedText := TokenizeText("Hello World! Now, this     is a sentence! ")

	assert.Equal(t, 2, len(parsedText.GetSentences()), "Sentence count")
	assert.Equal(t, 2, len(parsedText.GetSentences()[0].GetWords()), "Word count")
	assert.Equal(t, 5, len(parsedText.GetSentences()[1].GetWords()), "Word count")

	assert.Equal(t, "Hello", parsedText.GetSentences()[0].GetWords()[0])
	assert.Equal(t, "Now", parsedText.GetSentences()[1].GetWords()[0])
	assert.Equal(t, "sentence", parsedText.GetSentences()[1].GetWords()[4])

	parsedText2 := TokenizeText("This is a sentence without mark")

	assert.Equal(t, 6, len(parsedText2.GetSentences()[0].GetWords()))
}
