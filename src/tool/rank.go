package tool

type Rank struct {
	Max       int
	Min       int
	Relation  Relation
	Sentences []Sentence
	Words     map[int]*Word
	WordValID map[string]int
}

type Relation struct {
	Max    int
	Min    int
	Scores map[int]map[int]Score
}

type Score struct {
	Qty    int
	Weight float32
}

type Sentence struct {
	ID   int
	Text string
	//@todo store words string because rebuild phrases
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

func (rank *Rank) AddRelation(wordID int, relatedWordID int) {
	count := 0

	if relatedWordID == -1 {
		return
	}

	if _, ok := rank.Relation.Scores[relatedWordID][wordID]; ok {
		count = rank.Relation.Scores[relatedWordID][wordID].Qty + 1
		rank.Relation.Scores[relatedWordID][wordID] = Score{count, 0}

		return
	}

	if _, ok := rank.Relation.Scores[wordID][relatedWordID]; ok {
		count = rank.Relation.Scores[wordID][relatedWordID].Qty + 1
		rank.Relation.Scores[wordID][relatedWordID] = Score{count, 0}

		return

	}

	if _, ok := rank.Relation.Scores[wordID]; ok {
		count = 1
		rank.Relation.Scores[wordID][relatedWordID] = Score{count, 0}

		return
	}

	if _, ok := rank.Relation.Scores[relatedWordID]; ok {
		count = 1
		rank.Relation.Scores[relatedWordID][wordID] = Score{count, 0}

		return
	}

	count = 1
	rank.Relation.Scores[wordID] = map[int]Score{}
	rank.Relation.Scores[wordID][relatedWordID] = Score{count, 0}

	return
}

func (rank *Rank) UpdateRelatedMinMax(value int) {
	if rank.Relation.Max < value {
		rank.Relation.Max = value
	} else if rank.Relation.Min > value {
		rank.Relation.Min = value
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
