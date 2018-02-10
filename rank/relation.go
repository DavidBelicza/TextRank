package rank

type Relation struct {
	Max  int
	Min  int
	Node map[int]map[int]Score
}

type Score struct {
	Qty         int
	Weight      float32
	SentenceIDs []int
}

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
