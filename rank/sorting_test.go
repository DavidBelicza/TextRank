package rank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSentences(t *testing.T) {
	ranking := NewRank()

	sentences := FindSentences(ranking, 999, 1)

	assert.Equal(t, 0, len(sentences))
}
