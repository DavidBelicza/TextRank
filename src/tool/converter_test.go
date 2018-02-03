package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	text := Text{}
	text.Append(
		"There is a tree in the forest.",
		[]string{"There", "is", "apple", "tree", "in", "the", "forest"},
	)

	text.Append(
		"It has an apple tree",
		[]string{"It", "has", "an", "apple", "tree"},
	)

	rank := Rank{
		0,
		0,
		Relation{
			0,
			0,
			make(map[int]map[int]Score),
		},
		[]Sentence{},
		make(map[int]*Word),
		make(map[string]int),
	}

	Convert(text.GetSentences()[0], &rank)
	Convert(text.GetSentences()[1], &rank)

	id := rank.WordValID["tree"]

	assert.True(t, id > 0)
	assert.EqualValues(t, 2, len(rank.Sentences))

}
