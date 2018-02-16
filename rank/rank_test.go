package rank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWordData(t *testing.T) {
	ranking := NewRank()

	words := ranking.GetWordData()

	assert.Equal(t, 0, len(words))

	ranking.AddNewWord("word1", 0, 1)
	ranking.AddNewWord("word2", 1, 1)

	assert.Equal(t, 2, len(words))
}
