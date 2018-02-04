package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/DavidBelicza/TextRank/parse"
	"github.com/DavidBelicza/TextRank/rank"
)

func TestConvert(t *testing.T) {
	text := parse.Text{}
	text.Append(
		"There is a tree in the forest.",
		[]string{"There", "is", "apple", "tree", "in", "the", "forest"},
	)

	text.Append(
		"It has an apple tree",
		[]string{"It", "has", "an", "apple", "tree"},
	)

	curRank := rank.Rank{
		0,
		0,
		rank.Relation{
			0,
			0,
			make(map[int]map[int]rank.Score),
		},
		[]rank.Sentence{},
		make(map[int]*rank.Word),
		make(map[string]int),
	}

	id := curRank.WordValID["tree"]

	assert.True(t, id > 0)
	assert.EqualValues(t, 2, len(curRank.Sentences))

}