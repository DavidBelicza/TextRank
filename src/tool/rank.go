package tool

type Rank struct {
	Scores    map[int]map[int]float64
	Sentences []Sentence
	Words     map[int]*Word
	WordValID map[string]int
}

type Sentence struct {
	ID   int
	Text string
}

type Word struct {
	ID              int
	SentenceIDs     []int
	ConnectionLeft  map[int]int
	ConnectionRight map[int]int
	Value           string
}

func NewRank() *Rank {
	return &Rank{
		make(map[int]map[int]float64),
		[]Sentence{},
		make(map[int]*Word),
		make(map[string]int),
	}
}

func (rank *Rank) ResetItem(wordID int, relatedWordID int) {
	if _, ok := rank.Scores[relatedWordID][wordID]; ok {
		rank.Scores[relatedWordID][wordID] = -1
	} else {
		rank.Scores[wordID][relatedWordID] = -1
	}
}

func (rank *Rank) IsWordExist(word string) bool {
	_, find := rank.WordValID[word]

	return find
}

func (rank *Rank) AddNewWord(word string, prevWordIdx int) (wordID int) {
	sentenceID := len(rank.Sentences)
	wordID = len(rank.Words)
	connectionLeft := make(map[int]int)

	if prevWordIdx >= 0 {
		connectionLeft[prevWordIdx] = 1
	}

	newWord := &Word{
		ID:              wordID,
		SentenceIDs:     []int{sentenceID},
		ConnectionLeft:  connectionLeft,
		ConnectionRight: make(map[int]int),
		Value:           word,
	}

	rank.Words[wordID] = newWord
	rank.WordValID[word] = wordID

	return
}

func (rank *Rank) UpdateWord(word string, prevWordIdx int) (wordID int) {
	wordID = rank.WordValID[word]

	possibleSentenceID := len(rank.Sentences)
	for _, sentenceID := range rank.Words[wordID].SentenceIDs {
		if sentenceID == possibleSentenceID {
			rank.Words[wordID].SentenceIDs = append(
				rank.Words[wordID].SentenceIDs,
				possibleSentenceID,
			)
			break
		}
	}

	if prevWordIdx >= 0 {
		rank.Words[wordID].ConnectionLeft[prevWordIdx] += 1
	}

	return
}

func (rank *Rank) UpdateRightConnection(wordID int, rightWordID int) {
	if wordID >= 0 {
		rank.Words[wordID].ConnectionRight[rightWordID] += 1
	}
}

func (rank *Rank) GetWordData() map[int]*Word {
	return rank.Words
}
