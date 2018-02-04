package textrank

import (
	"rank"
	"parse"
	"convert"
)


var provider = make(map[int]*rank.Rank)

func AddText(text string, lang string, id int) {
	var textRank *rank.Rank

	if savedTextRank, ok := provider[id]; ok {
		textRank = savedTextRank
	} else {
		textRank = rank.NewRank()
		provider[id] = textRank
	}

	language := convert.NewLanguage()
	language.SetDefaultLanguage(lang)

	parsedText := parse.TokenizeText(text)

	for _, sentence := range parsedText.GetSentences() {
		convert.TextToRank(sentence, language, provider[id])
	}

	rank.Calculate(provider[id])
}

func GetPhrases(id int) []rank.Phrase  {
	return rank.GetPhrases(provider[id])
}
