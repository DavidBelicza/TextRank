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
