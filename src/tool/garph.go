package tool

// Graph data object carries words and the original sentences.
type Graph struct {
	Sentences []Sentence
	Words     map[int]*Word
	WordValID map[string]int
}

// Sentence struct contains the original sentences and their IDs.
type Sentence struct {
	ID   int
	Text string
}

// Word struct contains words without duplication.
// Its properties are:
// - ID: ID of the word (It is unique).
// - SentenceIDs: ID of the sentences where this word is.
// - ConnectionLeft: IDs of the words what are neighbors of the current one at left side.
// - ConnectionRight: IDs of the words what are neighbors of the current one at right side.
// - Value: The word itself.
// - Score: Normalized rank of the word in the whole text.
// - Count: How many times used this word.
type Word struct {
	ID              int
	SentenceIDs     []int
	ConnectionLeft  []int
	ConnectionRight []int
	Value           string
	Score           float32
	Count           int
}

func (graph *Graph) IsWordExist(word string) bool {
	_, find := graph.WordValID[word]

	return find
}

func (graph *Graph) AddNewWord(word string, prevWordIdx int) (wordID int) {
	sentenceID := len(graph.Sentences)
	wordID = len(graph.Words)
	connectionLeft := []int{}

	if prevWordIdx >= 0 {
		connectionLeft = []int{prevWordIdx}
	}

	newWord := &Word{
		ID:              wordID,
		SentenceIDs:     []int{sentenceID},
		ConnectionLeft:  connectionLeft,
		ConnectionRight: []int{},
		Value:           word,
		Score:           0,
		Count:           1,
	}

	graph.Words[wordID] = newWord
	graph.WordValID[word] = wordID

	return
}

func (graph *Graph) UpdateWord(word string, prevWordIdx int) (wordID int) {
	wordID = graph.WordValID[word]

	graph.Words[wordID].Count++

	possibleSentenceID := len(graph.Sentences)
	for _, sentenceID := range graph.Words[wordID].SentenceIDs {
		if sentenceID == possibleSentenceID {
			graph.Words[wordID].SentenceIDs = append(
				graph.Words[wordID].SentenceIDs,
				possibleSentenceID,
			)
			break
		}
	}

	if prevWordIdx >= 0 {
		graph.Words[wordID].ConnectionLeft = append(
			graph.Words[wordID].ConnectionLeft,
			prevWordIdx,
		)
	}

	return
}

func (graph *Graph) UpdateRightConnection(wordID int, rightWordID int) {
	if wordID >= 0 {
		word := graph.Words[wordID]
		word.ConnectionRight = append(word.ConnectionRight, rightWordID)
	}
}
