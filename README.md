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
	rawText := ""Your long raw text, it could be a book. Lorem ipsum...""
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

### Using different algorithm to ranking text

```go
package main

import (
	"fmt"
	
	"github.com/DavidBelicza/TextRank"
)

func main() {
	rawText := ""Your long raw text, it could be a book. Lorem ipsum...""
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

### Using different non-English languages.

```go
package main

import (
	"fmt"
	
	"github.com/DavidBelicza/TextRank"
)

func main() {
	rawText := ""Your long raw text, it could be a book. Lorem ipsum...""
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

## A simple visual representation
<img src="http://i.picresize.com/images/2018/01/30/PTn3Y.png" />
