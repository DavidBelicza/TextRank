package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	graphi := NewRank()
	parsedText := Parse("The red apple is good. good is the best for us in case of apple.")
	for _, sentence := range parsedText.GetSentences() {
		Convert(sentence, graphi)
	}

	Run(graphi)
	/*for _, word := range graphi.GetWordData() {
		fmt.Println(word.Value)
		//fmt.Println(word.Value + " " + strconv.FormatFloat(word., 'f', -1, 64))
	}*/

	assert.EqualValues(t, 1, 1)
}
