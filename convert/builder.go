package convert

import (
	"rank"
	"parse"
)

func TextToRank(sentence parse.ParsedSentence, lang *Language, ranks *rank.Rank) {
	addSentence(ranks, sentence)
	addWord(ranks, sentence.GetWords(), lang)
}

func addWord(ranks *rank.Rank, words []string, lang *Language) {
	prevWordID := -1
	curWordID := -1

	for _, word := range words {
		if !lang.IsStopWord(word) {
			if !ranks.IsWordExist(word) {
				curWordID = ranks.AddNewWord(word, prevWordID)
			} else {
				curWordID = ranks.UpdateWord(word, prevWordID)
			}

			ranks.Relation.AddRelation(curWordID, prevWordID)
			ranks.UpdateRightConnection(prevWordID, curWordID)

			prevWordID = curWordID
		}
	}
}

func addSentence(ranks *rank.Rank, sentence parse.ParsedSentence) {
	newSentence := rank.Sentence{
		ID:   len(ranks.Sentences),
		Text: sentence.GetOriginal(),
	}

	ranks.Sentences = append(ranks.Sentences, newSentence)
}
