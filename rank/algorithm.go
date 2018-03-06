package rank

import (
	"math"
	"fmt"
)

// Algorithm interface and its methods make possible the polimorf usage of
// weighting process.
type Algorithm interface {
	WeightingRelation(
		word1ID int,
		word2ID int,
		rank *Rank,
	) float32

	WeightingHits(
		wordID int,
		rank *Rank,
	) float32
}

// AlgorithmDefault struct is the basic implementation of Algorithm. It can
// weight and normalize a word or phrase by comparing them.
type AlgorithmDefault struct{}

// NewAlgorithmDefault constructor retrieves an AlgorithmDefault pointer.
func NewAlgorithmDefault() *AlgorithmDefault {
	return &AlgorithmDefault{}
}

// WeightingRelation method is the traditional algorithm of text rank to
// weighting and normalizing a phrase. It always retrieves a float number
// between 0.00 and 1.00.
func (a *AlgorithmDefault) WeightingRelation(
	word1ID int,
	word2ID int,
	rank *Rank,
) float32 {
	relationQty := rank.Relation.Node[word1ID][word2ID].Qty

	if math.IsNaN(float64(relationQty)) {
		return 0
	}

	return float32(relationQty)
}

// WeightingHits method ranks the words by their number of usage. It always
// retrieves a float number between 0.00 and 1.00.
func (a *AlgorithmDefault) WeightingHits(
	wordID int,
	rank *Rank,
) float32 {
	weight := rank.Words[wordID].Qty

	if math.IsNaN(float64(weight)) {
		return 0
	}

	return float32(weight)
}

// AlgorithmMixed struct is the combined implementation of Algorithm. A good
// example how weighting can be changed by a different implementations. It can
// weight and normalize a word or phrase by comparing them.
type AlgorithmMixed struct{}

// NewAlgorithmMixed constructor retrieves an AlgorithmMixed pointer.
func NewAlgorithmMixed() *AlgorithmMixed {
	return &AlgorithmMixed{}
}

// WeightingRelation method is a combined algorithm of text rank and word
// intensity it weights and normalizes a phrase. It always retrieves a float
// number between 0.00 and 1.00.
func (a *AlgorithmMixed) WeightingRelation(
	word1ID int,
	word2ID int,
	rank *Rank,
) float32 {
	relationQty := rank.Relation.Node[word1ID][word2ID].Qty

	l := false
	if rank.Words[word1ID].Token == "extension" && rank.Words[word2ID].Token == "gnome" {
		fmt.Println("run")
		l = true
	}

	logging := func(word1ID int, word2ID int) {
		if l {
			fmt.Println(rank.Words[word1ID].Token + " - " + rank.Words[word2ID].Token)
		}
	}

	qty := 0;

	for otherW2ID := range rank.Words[word1ID].ConnectionRight {
		if otherW2ID != word2ID {
			if v, ok := rank.Relation.Node[word1ID][otherW2ID]; ok {
				qty += v.Qty
				logging(word1ID, otherW2ID)
			} else if v, ok := rank.Relation.Node[otherW2ID][word1ID]; ok {
				logging(otherW2ID, word1ID)
				qty += v.Qty
			}
		}
	}

	for otherW2ID := range rank.Words[word1ID].ConnectionLeft {
		if otherW2ID != word2ID {
			if v, ok := rank.Relation.Node[word1ID][otherW2ID]; ok {
				qty += v.Qty
				logging(word1ID, otherW2ID)
			} else if v, ok := rank.Relation.Node[otherW2ID][word1ID]; ok {
				qty += v.Qty
				logging(otherW2ID, word1ID)
			}
		}
	}

	for otherW1ID := range rank.Words[word2ID].ConnectionRight {
		if otherW1ID != word1ID {
			if v, ok := rank.Relation.Node[word2ID][otherW1ID]; ok {
				qty += v.Qty
				logging(word2ID, otherW1ID)
			} else if v, ok := rank.Relation.Node[otherW1ID][word2ID]; ok {
				qty += v.Qty
				logging(otherW1ID, word2ID)
			}
		}
	}

	for otherW1ID := range rank.Words[word2ID].ConnectionLeft {
		if otherW1ID != word1ID {
			if v, ok := rank.Relation.Node[word2ID][otherW1ID]; ok {
				qty += v.Qty
				logging(word2ID, otherW1ID)
			} else if v, ok := rank.Relation.Node[otherW1ID][word2ID]; ok {
				qty += v.Qty
				logging(otherW1ID, word2ID)
			}
		}
	}

	if math.IsNaN(float64(relationQty)) {
		return 0
	}

	//@todo word count?
	return float32(relationQty) + (float32(qty)/100)
}

// WeightingHits method ranks the words by their number of usage. It always
// retrieves a float number between 0.00 and 1.00.
func (a *AlgorithmMixed) WeightingHits(
	wordID int,
	rank *Rank,
) float32 {
	weight := rank.Words[wordID].Qty

	if math.IsNaN(float64(weight)) {
		return 0
	}

	return float32(weight)
}
