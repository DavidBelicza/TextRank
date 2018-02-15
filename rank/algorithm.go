package rank

import "math"

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

type AlgorithmDefault struct{}

func NewAlgorithmDefault() *AlgorithmDefault {
	return &AlgorithmDefault{}
}

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

type AlgorithmMixed struct{}

func NewAlgorithmMixed() *AlgorithmMixed {
	return &AlgorithmMixed{}
}

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
