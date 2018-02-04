package rank

type Rank struct {
	Max       int
	Min       int
	Relation  Relation
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
		[]Sentence{},
		make(map[int]*Word),
		make(map[string]int),
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
		Qty:             1,
		Weight:          0,
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
