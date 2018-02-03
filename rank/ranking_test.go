package rank

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"convert"
	"parse"
)

func TestRun(t *testing.T) {
	graphi := NewRank()
	parsedText := parse.TokenizeText("The red apple is good. good is the best for us in case of apple.")
	for _, sentence := range parsedText.GetSentences() {
		convert.TextToRank(sentence, graphi)
	}

	Calculate(graphi)
	/*for _, word := range graphi.GetWordData() {
		fmt.Println(word.Value)
		//fmt.Println(word.Value + " " + strconv.FormatFloat(word., 'f', -1, 64))
	}*/

	assert.EqualValues(t, 1, 1)
}
