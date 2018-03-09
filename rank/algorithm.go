package rank

import (
	"math"
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
// weight a word or phrase by comparing them.
type AlgorithmDefault struct{}

// NewAlgorithmDefault constructor retrieves an AlgorithmDefault pointer.
func NewAlgorithmDefault() *AlgorithmDefault {
	return &AlgorithmDefault{}
}

// WeightingRelation method is the traditional algorithm of text rank to
// weighting a phrase.
func (a *AlgorithmDefault) WeightingRelation(
	word1ID int,
	word2ID int,
	rank *Rank,
) float32 {
	relationQty := rank.Relation.Node[word1ID][word2ID].Qty

	return float32(relationQty)
}

// WeightingHits method ranks the words by their occurrence.
func (a *AlgorithmDefault) WeightingHits(
	wordID int,
	rank *Rank,
) float32 {
	weight := rank.Words[wordID].Qty

	return float32(weight)
}

// AlgorithmChain struct is the combined implementation of Algorithm. It is a
// good example how weighting can be changed by a different implementations. It
// can weight a word or phrase by comparing them.
type AlgorithmChain struct{}

// NewAlgorithmChain constructor retrieves an AlgorithmChain pointer.
func NewAlgorithmChain() *AlgorithmChain {
	return &AlgorithmChain{}
}

// WeightingRelation method is a combined algorithm of text rank and word
// occurrence, it weights a phrase.
func (a *AlgorithmChain) WeightingRelation(
	word1ID int,
	word2ID int,
	rank *Rank,
) float32 {
	relationQty := rank.Relation.Node[word1ID][word2ID].Qty
	word1Qty := rank.Words[word1ID].Qty
	word2Qty := rank.Words[word2ID].Qty

	qDiff := float32(math.Abs(float64(word1Qty)-float64(word2Qty))) / 100
	weight := float32(relationQty) + qDiff

	return weight
}

// WeightingHits method ranks the words by their occurrence.
func (a *AlgorithmChain) WeightingHits(
	wordID int,
	rank *Rank,
) float32 {
	word := rank.Words[wordID]
	qty := 0

	for leftWordID, leftWordQty := range word.ConnectionLeft {
		qty += rank.Words[leftWordID].Qty * leftWordQty
	}

	for rightWordID, rightWordQty := range word.ConnectionRight {
		qty += rank.Words[rightWordID].Qty * rightWordQty
	}

	weight := float32(word.Qty) + (float32(qty))

	return float32(weight)
}
