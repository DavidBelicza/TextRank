package rank

type Algorithm interface {
	WeightingRelation(
		word1ID int,
		word2ID int,
		relationQty int,
		relationMin int,
		relationMax int,
		word1Qty int,
		word2Qty int,
		word1Min int,
		word1Max int,
	) float32

	WeightingHits(
		word1ID int,
		word1Qty int,
		word1Min int,
		word1Max int,
	) float32
}

type AlgorithmDefault struct {}

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
	word1Min int,
	word1Max int,
) float32 {
	return (float32(relationQty) - float32(relationMin)) / (float32(relationMax) - float32(relationMin))
}

func (a *AlgorithmDefault) WeightingHits(
	word1ID int,
	word1Qty int,
	word1Min int,
	word1Max int,
) float32 {
	return (float32(word1Qty) - float32(word1Min)) / (float32(word1Max) - float32(word1Min))
}

type AlgorithmMixed struct {}

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
	word1Min int,
	word1Max int,
) float32 {
	min := float32(relationMin + word1Min)
	max := float32(relationMax + word1Max)
	qty := float32(relationQty + word1Qty)

	return (qty - min) / (max - min)
}

func (a *AlgorithmMixed) WeightingHits(
	word1ID int,
	word1Qty int,
	word1Min int,
	word1Max int,
) float32 {
	return (float32(word1Qty) - float32(word1Min)) / (float32(word1Max) - float32(word1Min))
}
