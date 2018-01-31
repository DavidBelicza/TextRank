package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	connectionLeft := make(map[int]int)
	connectionLeft[2] = 1
	connectionLeft[5] = 3

	connectionRight := make(map[int]int)
	connectionRight[65] = 4
	connectionRight[74] = 12

	word := &Word{
		0,
		[]int{1},
		connectionLeft,
		connectionRight,
		"apple",
	}

	words := make(map[int]*Word)
	words[0] = word

	wordValIDs := make(map[string]int)
	wordValIDs["apple"] = 0

	sentence := Sentence{
		1,
		"Old apple tree in sunshine.",
	}

	sentences := []Sentence{}
	sentences = append(sentences, sentence)

	rank := Rank{
		make(map[int]map[int]float64),
		sentences,
		words,
		wordValIDs,
	}

	assert.EqualValues(t, "apple", rank.Words[0].Value)
	assert.EqualValues(t, rank.Sentences[0].ID, rank.Words[0].SentenceIDs[0])
	assert.EqualValues(t, 0, rank.WordValID["apple"])
}
