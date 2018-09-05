package rank

// Relation struct contains the phrase data.
//
// Max is the occurrence of the most used phrase.
//
// Min is the occurrence of the less used phrase. It is always greater then 0.
//
// Node is contains the Scores. Firs ID is the word 1, second ID is the word 2,
// and the value is the Score what contains the data about their relation.
type Relation struct {
	Max  float32
	Min  float32
	Node map[int]map[int]Score
}

// Score struct contains data about a relation of two words.
//
// Qty is the occurrence of the phrase.
//
// Weight is the weight of the phrase between 0.00 and 1.00.
//
// SentenceIDs contains all IDs of sentences what contain the phrase.
type Score struct {
	Qty         int
	Weight      float32
	SentenceIDs []int
}

// AddRelation method adds a new relation to Relation object.
func (relation *Relation) AddRelation(wordID int, relatedWordID int, sentenceID int) {
	if relatedWordID == -1 {
		return
	}

	if relation.updateRelation(relatedWordID, wordID, true, sentenceID) {
		return
	}

	if relation.extendRelation(wordID, relatedWordID, true, sentenceID) {
		return
	}

	relation.createRelation(wordID, relatedWordID, sentenceID)
}

func (relation *Relation) updateRelation(x int, y int, r bool, sentenceID int) bool {
	if _, ok := relation.Node[x][y]; ok {
		count := relation.Node[x][y].Qty + 1
		weight := relation.Node[x][y].Weight
		sentenceIDs := append(relation.Node[x][y].SentenceIDs, sentenceID)
		relation.Node[x][y] = Score{count, weight, sentenceIDs}

		return true
	} else if r {
		return relation.updateRelation(y, x, false, sentenceID)
	}

	return false
}

func (relation *Relation) extendRelation(x int, y int, r bool, sentenceID int) bool {
	if _, ok := relation.Node[x]; ok {
		relation.Node[x][y] = Score{1, 0, []int{sentenceID}}

		return true
	} else if r {
		return relation.extendRelation(y, x, false, sentenceID)
	}

	return false
}

func (relation *Relation) createRelation(x int, y int, sentenceID int) {
	relation.Node[x] = map[int]Score{}
	relation.Node[x][y] = Score{1, 0, []int{sentenceID}}
}
