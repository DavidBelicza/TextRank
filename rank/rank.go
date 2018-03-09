package rank

// Rank struct contains every original raw sentences, words, tokens, phrases,
// indexes, word hits, phrase hits and minimum-maximum values.
//
// Max is the occurrence of the most used word.
//
// Min is the occurrence of the less used word. It is always greater then 0.
//
// Relation is the Relation object, contains phrases.
//
// SentenceMap contains raw sentences. Index is the sentence ID, value is the
// sentence itself.
//
// Words contains Word objects. Index is the word ID, value is the word/token
// itself.
//
// WordValID contains words. Index is the word/token, value is the ID.
type Rank struct {
	Max         float32
	Min         float32
	Relation    Relation
	SentenceMap map[int]string
	Words       map[int]*Word
	WordValID   map[string]int
}

// Word struct contains all data about the words.
//
// If a word is multiple times in the text then the multiple words point to the
// same ID. So Word is unique.
//
// SentenceIDs contains all IDs of sentences what contain the word.
//
// ConnectionLeft contains all words what are connected to this word on the left
// side. The map index is the ID of the related word and its value is the
// occurrence.
//
// ConnectionRight contains all words what are connected to this word on the
// right side. The map index is the ID of the related word and its value is the
// occurrence.
//
// Token is the word itself, but not the original, it is tokenized.
//
// Qty is the number of occurrence of the word.
//
// Weight is the weight of the word between 0.00 and 1.00.
type Word struct {
	ID              int
	SentenceIDs     []int
	ConnectionLeft  map[int]int
	ConnectionRight map[int]int
	Token           string
	Qty             int
	Weight          float32
}

// NewRank constructor retrieves a Rank pointer.
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

// IsWordExist method retrieves true when the given word is already in the rank.
func (rank *Rank) IsWordExist(word string) bool {
	_, find := rank.WordValID[word]

	return find
}

// AddNewWord method adds a new word to the rank object and it defines its ID.
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

// UpdateWord method update a word what already exists in the rank object. It
// retrieves its ID.
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

// UpdateRightConnection method adds the right connection to the word. It always
// can be used after a word has added and the next word is known.
func (rank *Rank) UpdateRightConnection(wordID int, rightWordID int) {
	if wordID >= 0 {
		rank.Words[wordID].ConnectionRight[rightWordID]++
	}
}

// GetWordData method retrieves all words as a pointer.
func (rank *Rank) GetWordData() map[int]*Word {
	return rank.Words
}
