package tool

var graph *Graph

func Convert(sentence ParsedSentence, currenGraph *Graph) {
	graph = currenGraph

	addSentence(sentence)
	addWord(sentence.GetWords())
}

func addWord(words []string) {
	prevWordID := -1
	curWordID := -1

	for _, word := range words {
		if !graph.IsWordExist(word) {
			curWordID = graph.AddNewWord(word, prevWordID)
		} else {
			curWordID = graph.UpdateWord(word, prevWordID)
		}

		graph.UpdateRightConnection(prevWordID, curWordID)
		curWordID = prevWordID
	}

	updateRanks()
}

func addSentence(sentence ParsedSentence) {
	newSentence := Sentence{
		ID:   len(graph.Sentences),
		Text: sentence.GetOriginal(),
	}

	graph.Sentences = append(graph.Sentences, newSentence)
}
