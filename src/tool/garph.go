package tool

// Graph data object carries words and the original sentences.
type Graph struct {
	Sentences []Sentence
	Words     map[string]Word
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
