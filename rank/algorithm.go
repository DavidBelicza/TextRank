package rank

import "math"

// Algorithm interface and its methods make possible the polimorf usage of
// weighting process.
type Algorithm interface {
	WeightingRelation(
		word1ID int,
		word2ID int,
		relationQty int,
		relationMin int,
		relationMax int,
		word1Qty int,
		word2Qty int,
		wordQtyMin int,
		wordQtyMax int,
	) float32

	WeightingHits(
		wordID int,
		wordQty int,
		wordMin int,
		wordMax int,
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
	relationQty int,
	relationMin int,
	relationMax int,
	word1Qty int,
	word2Qty int,
	wordQtyMin int,
	wordQtyMax int,
) float32 {
	weight := (float32(relationQty) - float32(relationMin)) / (float32(relationMax) - float32(relationMin))

	if math.IsNaN(float64(weight)) {
		return 0
	}

	return weight
}

// WeightingHits method ranks the words by their number of usage. It always
// retrieves a float number between 0.00 and 1.00.
func (a *AlgorithmDefault) WeightingHits(
	wordID int,
	wordQty int,
	wordMin int,
	wordMax int,
) float32 {
	weight := (float32(wordQty) - float32(wordMin)) / (float32(wordMax) - float32(wordMin))

	if math.IsNaN(float64(weight)) {
		return 0
	}

	return weight
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
	relationQty int,
	relationMin int,
	relationMax int,
	word1Qty int,
	word2Qty int,
	wordQtyMin int,
	wordQtyMax int,
) float32 {
	min := float32(relationMin + wordQtyMin)
	max := float32(relationMax + wordQtyMax)
	qty := float32(relationQty + word1Qty)

	weight := (qty - min) / (max - min)

	if math.IsNaN(float64(weight)) {
		return 0
	}

	return weight
}

// WeightingHits method ranks the words by their number of usage. It always
// retrieves a float number between 0.00 and 1.00.
func (a *AlgorithmMixed) WeightingHits(
	wordID int,
	wordQty int,
	wordMin int,
	wordMax int,
) float32 {
	weight := (float32(wordQty) - float32(wordMin)) / (float32(wordMax) - float32(wordMin))

	if math.IsNaN(float64(weight)) {
		return 0
	}

	return weight
}
