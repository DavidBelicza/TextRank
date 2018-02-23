# TextRank on Go

[![GoDoc](https://godoc.org/github.com/DavidBelicza/TextRank?status.svg)](https://godoc.org/github.com/DavidBelicza/TextRank)
[![License: MIT](https://img.shields.io/badge/License-MIT-ee00ee.svg)](https://github.com/DavidBelicza/TextRank/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/DavidBelicza/TextRank.svg?branch=master)](https://travis-ci.org/DavidBelicza/TextRank)
[![Go Report Card](https://goreportcard.com/badge/github.com/DavidBelicza/TextRank)](https://goreportcard.com/report/github.com/DavidBelicza/TextRank)
[![Coverage Status](https://coveralls.io/repos/github/DavidBelicza/TextRank/badge.svg?branch=master)](https://coveralls.io/github/DavidBelicza/TextRank?branch=master)

## TextRank or Automatic summarization
> Automatic summarization is the process of reducing a text document with a computer program in order to create a summary that retains the most important points of the original document. Technologies that can make a coherent summary take into account variables such as length, writing style and syntax. Automatic data summarization is part of machine learning and data mining. The main idea of summarization is to find a representative subset of the data, which contains the information of the entire set. Summarization technologies are used in a large number of sectors in industry today. - Wikipedia

## EXAMEPLES

### Find the most important phrases

```go
package main

import (
	"fmt"
	
	"github.com/DavidBelicza/TextRank"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// ID of the text rank, any number.
	id := 1
	// Default Rule for parsing.
	rule := textrank.CreateDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.CreateDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.CreateDefaultAlgorithm()

	// Add text.
	textrank.Append(rawText, language, rule, id)
	// Run the ranking.
	textrank.Ranking(id, algorithmDef)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(id)

	// Most important phrase.
	fmt.Println(rankedPhrases[0])
	// Second important phrase.
	fmt.Println(rankedPhrases[1])
}
```

### Adding text continuously

```go
package main

import (
	"fmt"
	
	"github.com/DavidBelicza/TextRank"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// ID of the text rank, any number.
	id := 1
	// Default Rule for parsing.
	rule := textrank.CreateDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.CreateDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.CreateDefaultAlgorithm()

	// Add text.
	textrank.Append(rawText, language, rule, id)
	// Run the ranking.
	textrank.Ranking(id, algorithmDef)

	rawText2 := "Another book or article..."
	rawText3 := "Third another book or article..."

	// Add text to the previously added text.
	textrank.Append(rawText2, language, rule, id)
	// Add text to the previously added text.
	textrank.Append(rawText3, language, rule, id)
	// Run the ranking to the whole composed text.
	textrank.Ranking(id, algorithmDef)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(id)

	// Most important phrase.
	fmt.Println(rankedPhrases[0])
	// Second important phrase.
	fmt.Println(rankedPhrases[1])
}
```

### Using different algorithm to ranking text

```go
package main

import (
	"fmt"
	
	"github.com/DavidBelicza/TextRank"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// ID of the text rank, any number.
	id := 1
	// Default Rule for parsing.
	rule := textrank.CreateDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.CreateDefaultLanguage()
	// Using a little bit more complex algorithm to ranking text.
	algorithmMix := textrank.CreateMixedAlgorithm()
	
	// Add text.
	textrank.Append(rawText, language, rule, id)
	// Run the ranking.
	textrank.Ranking(id, algorithmMix)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(id)

	// Most important phrase.
	fmt.Println(rankedPhrases[0])
	// Second important phrase.
	fmt.Println(rankedPhrases[1])
}
```

### Using multiple graphs.

```go
package main

import (
	"fmt"
	
	"github.com/DavidBelicza/TextRank"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// ID of the text rank, any number.
	firstGraphID := 1
	// Default Rule for parsing.
	rule := textrank.CreateDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.CreateDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.CreateDefaultAlgorithm()

	// Add text.
	textrank.Append(rawText, language, rule, firstGraphID)
	// Run the ranking.
	textrank.Ranking(firstGraphID, algorithmDef)

	// ID of the text rank, any number.
	secondGraphID := 2

	// Using a little bit more complex algorithm to ranking text.
	algorithmMix := textrank.CreateMixedAlgorithm()

	// Add text to the second graph.
	textrank.Append(rawText, language, rule, secondGraphID)
	// Run the ranking on the second graph.
	textrank.Ranking(secondGraphID, algorithmMix)

	// Get all phrases by weight from first graph.
	rankedPhrases := textrank.FindPhrases(firstGraphID)

	// Most important phrase from first graph.
	fmt.Println(rankedPhrases[0])
	// Second important phrase from first graph.
	fmt.Println(rankedPhrases[1])

	// Get all phrases by weight from second graph.
	rankedPhrases2 := textrank.FindPhrases(secondGraphID)

	// Most important phrase from second graph.
	fmt.Println(rankedPhrases2[0])
	// Second important phrase from second graph.
	fmt.Println(rankedPhrases2[1])
}
```

### Using different non-English languages.

```go
package main

import (
	"fmt"
	
	"github.com/DavidBelicza/TextRank"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// ID of the text rank, any number.
	id := 1
	// Default Rule for parsing.
	rule := textrank.CreateDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.CreateDefaultLanguage()
	
	// Add Spanish stop words (just some example).
	language.SetWords("es", []string{"uno", "dos", "tres", "yo", "es", "eres"})
	// Active the Spanish.
	language.SetActiveLanguage("es")
	
	// Default algorithm for ranking text.
	algorithmDef := textrank.CreateDefaultAlgorithm()

	// Add text.
	textrank.Append(rawText, language, rule, id)
	// Run the ranking.
	textrank.Ranking(id, algorithmDef)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(id)

	// Most important phrase.
	fmt.Println(rankedPhrases[0])
	// Second important phrase.
	fmt.Println(rankedPhrases[1])
}
```

## A simple visual representation
<img src="http://i.picresize.com/images/2018/01/30/PTn3Y.png" />
