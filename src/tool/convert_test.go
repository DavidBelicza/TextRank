package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	text := Text{}
	text.Append(
		"There is a tree in the forest.",
		[]string{"There", "is", "a", "tree", "in", "the", "forest"},
	)

	text.Append(
		"It has an apple tree",
		[]string{"It", "has", "an", "apple", "tree"},
	)

	graph := Graph{
		[]Sentence{},
		make(map[int]*Word),
		make(map[string]int),
	}

	Convert(text.GetSentences()[0], &graph)
	Convert(text.GetSentences()[1], &graph)

	id := graph.WordValID["tree"]

	assert.True(t, id > 0)
	assert.EqualValues(t, 2, graph.Words[id].Count)
	assert.EqualValues(t, 2, len(graph.Sentences))
}
