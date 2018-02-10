package rank

type Rank struct {
	Max         int
	Min         int
	Relation    Relation
	SentenceMap map[int]string
	Words       map[int]*Word
	WordValID   map[string]int
}

type Word struct {
	ID              int
	SentenceIDs     []int
	ConnectionLeft  map[int]int
	ConnectionRight map[int]int
	Token           string
	Qty             int
	Weight          float32
}

func NewRank() *Rank {
	return &Rank{
		0,
		0,
		Relation{
			0,
			0,
			make(map[int]map[int]Score),
		},
		make(map[int]string),
		make(map[int]*Word),
		make(map[string]int),
	}
}

func (rank *Rank) IsWordExist(word string) bool {
	_, find := rank.WordValID[word]

	return find
}

func (rank *Rank) AddNewWord(word string, prevWordIdx int, sentenceID int) (wordID int) {
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
		Token:           word,
		Qty:             1,
		Weight:          0,
	}

	rank.Words[wordID] = newWord
	rank.WordValID[word] = wordID

	return
}

func (rank *Rank) UpdateWord(word string, prevWordIdx int, sentenceID int) (wordID int) {
	wordID = rank.WordValID[word]

	found := false

	for _, oldSentenceID := range rank.Words[wordID].SentenceIDs {
		if sentenceID == oldSentenceID {
			found = true
			break
		}
	}

	if !found {
		rank.Words[wordID].SentenceIDs = append(
			rank.Words[wordID].SentenceIDs,
			sentenceID,
		)
	}

	rank.Words[wordID].Qty++

	if prevWordIdx >= 0 {
		rank.Words[wordID].ConnectionLeft[prevWordIdx]++
	}

	return
}

func (rank *Rank) UpdateRightConnection(wordID int, rightWordID int) {
	if wordID >= 0 {
		rank.Words[wordID].ConnectionRight[rightWordID]++
	}
}

func (rank *Rank) GetWordData() map[int]*Word {
	return rank.Words
}
