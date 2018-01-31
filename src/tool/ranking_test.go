package tool

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	graphi := NewRank()
	parsedText := Parse("We should go to the forest. We need a lot of apple. This is an apple tree. But this apple is purple!")

	for _, sentence := range parsedText.GetSentences() {
		Convert(sentence, graphi)
	}

	Run(graphi)

	for _, word := range graphi.GetWordData() {
		fmt.Println(word.Value)
		//fmt.Println(word.Value + " " + strconv.FormatFloat(word., 'f', -1, 64))
	}

	assert.EqualValues(t, 1, 1)
}
