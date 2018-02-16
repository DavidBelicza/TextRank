package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizeText(t *testing.T) {
	lang := NewLanguage()

	lang.SetActiveLanguage("hu")
	lang.SetWords("hu", []string{"word1"})

	assert.Equal(t, true, lang.IsStopWord("word1"))
	assert.Equal(t, false, lang.IsStopWord("word2"))
}
