# TextRank on Go

[![GoDoc](https://godoc.org/github.com/DavidBelicza/TextRank?status.svg)](https://godoc.org/github.com/DavidBelicza/TextRank)
[![License: MIT](https://img.shields.io/badge/License-MIT-ee00ee.svg)](https://github.com/DavidBelicza/TextRank/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/DavidBelicza/TextRank.svg?branch=master)](https://travis-ci.org/DavidBelicza/TextRank)
[![Go Report Card](https://goreportcard.com/badge/github.com/DavidBelicza/TextRank)](https://goreportcard.com/report/github.com/DavidBelicza/TextRank)
[![Coverage Status](https://coveralls.io/repos/github/DavidBelicza/TextRank/badge.svg?branch=master)](https://coveralls.io/github/DavidBelicza/TextRank?branch=master)

This source code is an implementation of textrank algorithm up to Go 1.8, under MIT licence.

## MOTIVATION

If there was a program what could rank book size text's words, phrases and sentences continuously on multiple threads and it would be opened to modifing by objects, written in a simple, secure, static language and if it would be very well documented... Now, there is.

## FEATURES

* Find the most important phrases.
* Find the most important words.
* Find the most important N sentences. 
	* Importance by phrase weights.
	* Importance by phrase weights.
* Find the first N sentences, start from Xth sentence.
* Find sentences by phrase chains ordered by position in text.
* Access to the whole ranked data.
* Support more languages.
* Algorithm for weighting can be modified by interface implementation.
* Parser can be modified by interface implementation.
* Multi thread safe.


## TEXTRANK OR AUTOMATIC SUMMARIZATION
> Automatic summarization is the process of reducing a text document with a computer program in order to create a summary that retains the most important points of the original document. Technologies that can make a coherent summary take into account variables such as length, writing style and syntax. Automatic data summarization is part of machine learning and data mining. The main idea of summarization is to find a representative subset of the data, which contains the information of the entire set. Summarization technologies are used in a large number of sectors in industry today. - Wikipedia

## EXAMPLES

### Find the most important phrases

This is the most basic and simplest usage of textrank.

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

### All possible pre-defined finder queries

After ranking, the graph contains a lot of valuable data. There are functions in textrank package what contains logic to retrieve those data from the graph.

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

	// Get all phrases order by weight.
	rankedPhrases := textrank.FindPhrases(id)
	// Most important phrase.
	fmt.Println(rankedPhrases[0])

	// Get all words order by weight.
	words := textrank.FindSingleWords(id)
	// Most important word.
	fmt.Println(words[0])

	// Get the most important 10 sentences. Importance by phrase weights.
	sentences := textrank.FindSentencesByRelationWeight(id, 10)
	// Found sentences
	fmt.Println(sentences)

	// Get the most important 10 sentences. Importance by word occurrence.
	sentences = textrank.FindSentencesByWordQtyWeight(id, 10)
	// Found sentences
	fmt.Println(sentences)

	// Get the first 10 sentences, start from 5th sentence.
	sentences = textrank.FindSentencesFrom(id, 5, 10)
	// Found sentences
	fmt.Println(sentences)

	// Get sentences by phrase/word chains order by position in text.
	sentencesPh := textrank.FindSentencesByPhraseChain(id, []string{"gnome", "shell", "extension"})
	// Found sentence.
	fmt.Println(sentencesPh[0])
}
```

### Access to everything

After ranking, the graph contains a lot of valuable data. The GetRank function allows access to the graph and every data can be retrieved from this structure.

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

	// Get the rank graph.
	rankData := textrank.GetRank(id)

	// Get word ID by token/word.
	wordId := rankData.WordValID["gnome"]

	// Word's weight.
	fmt.Println(rankData.Words[wordId].Weight)
	// Word's quantity/occurrence.
	fmt.Println(rankData.Words[wordId].Qty)
	// All sentences what contain the this word.
	fmt.Println(rankData.Words[wordId].SentenceIDs)
	// All other words what are related to this word on left side.
	fmt.Println(rankData.Words[wordId].ConnectionLeft)
	// All other words what are related to this word on right side.
	fmt.Println(rankData.Words[wordId].ConnectionRight)
	// The node of this word, it contains the related words and the relation weight.
	fmt.Println(rankData.Relation.Node[wordId])
}
```

### Adding text continuously

It is possibe to add more text after another texts already have been added. The Ranking function can merge these multiple texts and it can recalculate the weights and all related data.

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

There are two algorithm has implemented, it is possible to write custom algorithm by Algorithm interface and use it instead of defaults.

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

### Using multiple graphs

Graph ID exists because it is possible run multiple independent text ranking processes.

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

### Using different non-English languages

Engish is used by default but it is possible to add any language. To use other languages a stop word list is required what you can find here: https://github.com/stopwords-iso

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

### Asynchronous usage by goroutines

It is thread safe. Independent graphs can recieve texts in the same time and can be extended by more text also in the same time.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/DavidBelicza/TextRank"
)

func main() {
// A flag when program has to stop.
	stopProgram := false
	// Channel.
	stream := make(chan string)
	// The text rank graph ID.
	id := 1

	// Open new thread/routine
	go func() {
		// 3 texts.
		rawTexts := []string{
			"Very long text...",
			"Another very long text...",
			"Second another very long text...",
		}

		// Add 3 texts to the stream channel, one by one.
		for _, rawText := range rawTexts {
			stream <- rawText
		}
	}()

	// Open new thread/routine
	go func() {
		// Counter how many times texts added to the ranking.
		i := 1

		for {
			// Get text from stream channel when it got a new one.
			rawText := <-stream

			// Default Rule for parsing.
			rule := textrank.CreateDefaultRule()
			// Default Language for filtering stop words.
			language := textrank.CreateDefaultLanguage()
			// Default algorithm for ranking text.
			algorithm := textrank.CreateDefaultAlgorithm()

			// Add text.
			textrank.Append(rawText, language, rule, id)
			// Run the ranking.
			textrank.Ranking(id, algorithm)

			// Set stopProgram flag to true when all 3 text have been added.
			if i == 3 {
				stopProgram = true
			}

			i++
		}
	}()

	// The main thread has to run while go-routines run. When stopProgram is
	// true then the loop has finish.
	for !stopProgram {
		time.Sleep(time.Second * 1)
	}

	// Most important phrase.
	phrases := textrank.FindPhrases(id)
	// Second important phrase.
	fmt.Println(phrases[0])
}
```

## A simple visual representation

The below image is a representation how works the simplest text ranking algorithm. This algorithm can be replaced by an another one by inject different Algorithm interface implementation.

<img src="http://i.picresize.com/images/2018/01/30/PTn3Y.png" />
