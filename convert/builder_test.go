package convert

import (
	"testing"

	"github.com/DavidBelicza/TextRank/v2/parse"
	"github.com/DavidBelicza/TextRank/v2/rank"
	"github.com/stretchr/testify/assert"
)

type languageWithRootWords struct {
	LanguageDefault
}

func (lang *languageWithRootWords) FindRootWord(word string) (bool, string) {
	if word == "apples" {
		return true, "apple"
	}

	return false, ""
}

func TestTextToRankReplacesWordsByRootWords(t *testing.T) {
	var text parse.Text
	text.Append("Apples are red.", []string{"apples", "are", "red"})

	lang := &languageWithRootWords{*NewLanguage()}
	ranks := rank.NewRank()

	for _, sentence := range text.GetSentences() {
		TextToRank(sentence, lang, ranks)
	}

	assert.True(t, ranks.IsWordExist("apple"))
	assert.False(t, ranks.IsWordExist("apples"))
}
