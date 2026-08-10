package rank

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNormalizeSingleWord(t *testing.T) {
	r := NewRank()
	r.AddNewWord("hello", -1, 0)

	Calculate(r, NewAlgorithmDefault())

	id := r.WordValID["hello"]
	weight := r.Words[id].Weight

	assert.False(t, math.IsNaN(float64(weight)))
	assert.Equal(t, float32(1), weight)
}

func TestNormalizeWithoutRange(t *testing.T) {
	assert.Equal(t, float32(1), normalize(5, 5, 5))
	assert.Equal(t, float32(0), normalize(0, 0, 0))
}
