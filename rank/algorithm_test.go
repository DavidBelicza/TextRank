package rank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightingRelation(t *testing.T) {
	rank := createRank()
	def := NewAlgorithmDefault()
	weightDef := def.WeightingRelation(0, 1, rank)

	assert.Equal(t, float32(2), weightDef)

	chain := NewAlgorithmChain()
	weightChain := chain.WeightingRelation(0, 1, rank)

	assert.Equal(t, float32(2.01), weightChain)

	weightChain = chain.WeightingRelation(2, 3, rank)

	assert.Equal(t, float32(1), weightChain)
}

func TestWeightingHits(t *testing.T) {
	rank := createRank()

	def := NewAlgorithmDefault()
	weightDef := def.WeightingHits(0, rank)

	assert.Equal(t, float32(2), weightDef)

	chain := NewAlgorithmChain()
	weightChain := chain.WeightingHits(0, rank)

	assert.Equal(t, float32(3), weightChain)

	weightChain = chain.WeightingHits(2, rank)

	assert.Equal(t, float32(3), weightChain)
}

func createRank() *Rank {
	rank := NewRank()
	rank.AddNewWord("word1", -1, 0)
	rank.AddNewWord("word2", 0, 0)
	rank.UpdateWord("word1", 1, 0)
	rank.AddNewWord("word3", 0, 0)
	rank.AddNewWord("word4", 2, 0)

	rank.Relation.AddRelation(0, 1, 0)
	rank.Relation.AddRelation(1, 0, 0)
	rank.Relation.AddRelation(0, 2, 0)
	rank.Relation.AddRelation(2, 3, 0)

	rank.Relation.Max = 3
	rank.Relation.Min = 1

	return rank
}
