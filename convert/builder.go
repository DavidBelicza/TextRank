package convert

import (
	"github.com/DavidBelicza/TextRank/parse"
	"github.com/DavidBelicza/TextRank/rank"
)

// TextToRank function converts a ParsedSentence object to Rank object, it is
// the preparing process to later text ranking.
func TextToRank(sentence parse.ParsedSentence, lang Language, ranks *rank.Rank) {
	sentenceId := addSentence(ranks, sentence)
	addWord(ranks, sentence.GetWords(), lang, sentenceId)
}

func addWord(ranks *rank.Rank, words []string, lang Language, sentenceID int) {
	prevWordID := -1
	curWordID := -1

	for _, word := range words {
		if !lang.IsStopWord(word) {
			if !ranks.IsWordExist(word) {
				curWordID = ranks.AddNewWord(word, prevWordID, sentenceID)
			} else {
				curWordID = ranks.UpdateWord(word, prevWordID, sentenceID)
			}

			ranks.Relation.AddRelation(curWordID, prevWordID, sentenceID)
			ranks.UpdateRightConnection(prevWordID, curWordID)

			prevWordID = curWordID
		}
	}
}

func addSentence(ranks *rank.Rank, sentence parse.ParsedSentence) int {
	ranks.SentenceMap[len(ranks.SentenceMap)] = sentence.GetOriginal()

	return len(ranks.SentenceMap) - 1
}
