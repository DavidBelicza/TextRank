package rank

type Relation struct {
	Max    int
	Min    int
	Scores map[int]map[int]Score
}

type Score struct {
	Qty    int
	Weight float32
}

func (relation *Relation) AddRelation(wordID int, relatedWordID int) {
	if relatedWordID == -1 {
		return
	}

	if relation.updateRelation(relatedWordID, wordID, true) {
		return
	}

	if relation.extendRelation(wordID, relatedWordID, true) {
		return
	}

	relation.createRelation(wordID, relatedWordID)
}

func (relation *Relation) updateRelation(x int, y int, r bool) bool {
	if _, ok := relation.Scores[x][y]; ok {
		count := relation.Scores[x][y].Qty + 1
		weight := relation.Scores[x][y].Weight
		relation.Scores[x][y] = Score{count, weight}

		return true
	} else if r {
		return relation.updateRelation(y, x, false)
	}

	return false
}

func (relation *Relation) extendRelation(x int, y int, r bool) bool {
	if _, ok := relation.Scores[x]; ok {
		relation.Scores[x][y] = Score{1, 0}

		return true
	} else if r {
		return relation.extendRelation(y, x, false)
	}

	return false
}

func (relation *Relation) createRelation(x int, y int) {
	relation.Scores[x] = map[int]Score{}
	relation.Scores[x][y] = Score{1, 0}
}
