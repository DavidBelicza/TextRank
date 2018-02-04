package convert

import (
	"rank"
	"parse"
)

var curRank *rank.Rank

func TextToRank(sentence parse.ParsedSentence, lang *Language, currentRank *rank.Rank) {
	curRank = currentRank

	addSentence(sentence)
	addWord(sentence.GetWords(), lang)
}

func addWord(words []string, lang *Language) {
	prevWordID := -1
	curWordID := -1

	for _, word := range words {
		if !lang.IsStopWord(word) {
			if !curRank.IsWordExist(word) {
				curWordID = curRank.AddNewWord(word, prevWordID)
			} else {
				curWordID = curRank.UpdateWord(word, prevWordID)
			}

			curRank.Relation.AddRelation(curWordID, prevWordID)
			curRank.UpdateRightConnection(prevWordID, curWordID)

			prevWordID = curWordID
		}
	}
}

func addSentence(sentence parse.ParsedSentence) {
	newSentence := rank.Sentence{
		ID:   len(curRank.Sentences),
		Text: sentence.GetOriginal(),
	}

	curRank.Sentences = append(curRank.Sentences, newSentence)
}
