package textrank

import (
	"github.com/DavidBelicza/TextRank/rank"
	"github.com/DavidBelicza/TextRank/parse"
	"github.com/DavidBelicza/TextRank/convert"
)


var provider = make(map[int]*rank.Rank)

func AddText(text string, lang string, id int) {
	var ranks *rank.Rank

	if savedTextRank, ok := provider[id]; ok {
		ranks = savedTextRank
	} else {
		ranks = rank.NewRank()
		provider[id] = ranks
	}

	language := convert.NewLanguage()
	language.SetDefaultLanguage(lang)

	parsedText := parse.TokenizeText(text)

	for _, sentence := range parsedText.GetSentences() {
		convert.TextToRank(sentence, language, provider[id])
	}
}

func Calculate(id int) {
	rank.Calculate(provider[id])
}

func GetPhrases(id int) []rank.Phrase  {
	return rank.GetPhrases(provider[id])
}
