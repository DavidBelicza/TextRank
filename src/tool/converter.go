package tool

var rank *Rank

func Convert(sentence ParsedSentence, currentRank *Rank) {
	rank = currentRank

	addSentence(sentence)
	addWord(sentence.GetWords())
}

func addWord(words []string) {
	prevWordID := -1
	curWordID := -1

	for _, word := range words {
		if !rank.IsWordExist(word) {
			curWordID = rank.AddNewWord(word, prevWordID)
		} else {
			curWordID = rank.UpdateWord(word, prevWordID)
		}

		rank.UpdateRightConnection(prevWordID, curWordID)
		curWordID = prevWordID
	}
}

func addSentence(sentence ParsedSentence) {
	newSentence := Sentence{
		ID:   len(rank.Sentences),
		Text: sentence.GetOriginal(),
	}

	rank.Sentences = append(rank.Sentences, newSentence)
}
