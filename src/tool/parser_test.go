package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	parsedText := Parse("Hello World! Now, this     is a sentence! ")
	assert.Equal(t, 7, len(parsedText))
}
