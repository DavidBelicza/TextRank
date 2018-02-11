package textrank

import (
	"github.com/DavidBelicza/TextRank/convert"
	"github.com/DavidBelicza/TextRank/parse"
	"github.com/DavidBelicza/TextRank/rank"
)

var provider = make(map[int]*rank.Rank)

func Append(text string, lang convert.Language, rule parse.Rule, id int) {
	var ranks *rank.Rank

	if savedTextRank, ok := provider[id]; ok {
		ranks = savedTextRank
	} else {
		ranks = rank.NewRank()
		provider[id] = ranks
	}

	parsedText := parse.TokenizeText(text, rule)

	for _, sentence := range parsedText.GetSentences() {
		convert.TextToRank(sentence, lang, provider[id])
	}
}

func Ranking(
	id int,
	algorithm rank.Algorithm,
) {
	rank.Calculate(provider[id], algorithm)
}

func CreateDefaultAlgorithm() *rank.AlgorithmDefault {
	return rank.NewAlgorithmDefault()
}

func CreateMixedAlgorithm() *rank.AlgorithmMixed {
	return rank.NewAlgorithmMixed()
}

func CreateDefaultLanguage() *convert.LanguageDefault {
	return convert.NewLanguage()
}

func CreateDefaultRule() *parse.RuleDefault {
	return parse.NewRule()
}

func GetRank(id int) *rank.Rank {
	return provider[id]
}

func FindPhrases(id int) []rank.Phrase {
	return rank.GetPhrases(provider[id])
}

func FindSingleWords(id int) []rank.SingleWord {
	return rank.GetSingleWords(provider[id])
}

func FindSentencesByRelationScore(id int, limit int) []rank.Sentence {
	return rank.GetSentences(provider[id], rank.ByRelation, limit)
}

func FindSentencesByWordQtyScore(id int, limit int) []rank.Sentence {
	return rank.GetSentences(provider[id], rank.ByQty, limit)
}

func FindSentencesByPhrases(id int, phrases []string) []rank.Sentence {
	return rank.GetSentencesByPhrases(provider[id], phrases)
}

func FindSentencesFrom(id int, sentenceID int, limit int) []rank.Sentence {
	return rank.GetSentencesFrom(provider[id], sentenceID, limit)
}
