<h1 align="center">
TextRank on Go
</h1>

<p align="center">
	<a href="https://godoc.org/github.com/DavidBelicza/TextRank">
		<img src="https://godoc.org/github.com/DavidBelicza/TextRank?status.svg" alt="GoDoc" />
	</a>
	<a href="https://github.com/DavidBelicza/TextRank/blob/master/LICENSE">
		<img src="https://img.shields.io/badge/License-MIT-ee00ee.svg" alt="License: MIT" />
	</a>
	<a href="https://travis-ci.org/DavidBelicza/TextRank">
		<img src="https://travis-ci.org/DavidBelicza/TextRank.svg?branch=master" alt="Build Status" />
	</a>
	<a href="https://goreportcard.com/report/github.com/DavidBelicza/TextRank">
		<img src="https://goreportcard.com/badge/github.com/DavidBelicza/TextRank" alt="Go Report Card" />
	</a>
	<a href="https://coveralls.io/github/DavidBelicza/TextRank?branch=master">
		<img src="https://coveralls.io/repos/github/DavidBelicza/TextRank/badge.svg?branch=master" alt="Coverage Status" />
	</a>
	<a href="https://github.com/DavidBelicza/TextRank/releases/latest">
		<img src="https://img.shields.io/github/release/DavidBelicza/TextRank.svg?colorB=269aca" alt="Release" />
	</a>
	
</p>

<p align="center">
This source code is an implementation of textrank algorithm, under MIT licence.
<br />The minimum requred Go version is 1.8.
<p align="center">
<br />	
	
## MOTIVATION

If there was a program what could rank book size text's words, phrases and sentences continuously on multiple threads and it would be opened to modifing by objects, written in a simple, secure, static language and if it would be very well documented... Now, here it is.

## FEATURES

* Find the most important phrases.
* Find the most important words.
* Find the most important N sentences. 
	* Importance by phrase weights.
	* Importance by word occurrence.
* Find the first N sentences, start from Xth sentence.
* Find sentences by phrase chains ordered by position in text.
* Access to the whole ranked data.
* Support more languages.
* Algorithm for weighting can be modified by interface implementation.
* Parser can be modified by interface implementation.
* Multi thread support.

## INSTALL

You can install TextRank by Go's get:

```go get github.com/DavidBelicza/TextRank```

TextRank uses the default Go *mod* as vendoring tool, so you can install the dependencies with this command:

```go mod vendor```

## DOCKER

Using Docker to TextRank isn't necessary, it's just an option.

Build image from the repository's root directory:

```docker build -t go_text_rank_image .```

Create container from the image:

```docker run -dit --name textrank go_text_rank_image:latest```

Run the **go test -v .** code inside the container:

```docker exec -i -t textrank go test -v .```

Stop, start or remove the container:

* ```docker stop textrank```
* ```docker start textrank```
* ```docker rm textrank```

## HOW DOES IT WORK

Too see how does it work, the easiest way is to use the sample text. Sample text can be found in the [textrank_test.go file at this line](https://github.com/DavidBelicza/TextRank/blob/master/textrank_test.go#L12). It's a short size text about Gnome Shell.

* TextRank reads the text, 
    * parse it, 
    * remove the unnecessary stop words,
    * tokenize it 
* and counting the occurrence of the words and phrases 
* and then it starts weighting
    * by the occurrence of words and phrases and their relations. 
* After weights are done, TextRank normalize weights to between 1 and 0.
* Then the different finder methods capable to find the most important words, phrases or sentences.

The most important phrases from the sample text are:

Phrase | Occurrence | Weight
--- | --- | ---
gnome - shell | 5 | 1
extension - gnome | 3 | 0.50859946
icons - tray | 3 | 0.49631447
gnome - caffeine | 2 | 0.27027023

The **gnome** is the most often used word in this text and **shell** is also used multiple times. Two of them are used together as a phrase 5 times. This is the highest occurrence in this text, so this is the most important phrase.

The following two important phrases have same occurrence 3, however they are not equal. This is because the **extension gnome** phrase contains the word **gnome**, the most popular word in the text, and it increases the phrase's weight. It increases the weight of any word what is related to it, but not too much to overcome other important phrases what don't contain the **gnome** word.

The exact algorithm can be found in the [algorithm.go file at this line](https://github.com/DavidBelicza/TextRank/blob/master/rank/algorithm.go#L65).

## TEXTRANK OR AUTOMATIC SUMMARIZATION
> Automatic summarization is the process of reducing a text document with a computer program in order to create a summary that retains the most important points of the original document. Technologies that can make a coherent summary take into account variables such as length, writing style and syntax. Automatic data summarization is part of machine learning and data mining. The main idea of summarization is to find a representative subset of the data, which contains the information of the entire set. Summarization technologies are used in a large number of sectors in industry today. - Wikipedia

## EXAMPLES

### Find the most important phrases

This is the most basic and simplest usage of textrank.

```go
package main

import (
	"fmt"
	
	"github.com/DavidBelicza/TextRank/v2"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.NewDefaultAlgorithm()

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(tr)

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
	
	"github.com/DavidBelicza/TextRank/v2"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.NewDefaultAlgorithm()

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get all phrases order by weight.
	rankedPhrases := textrank.FindPhrases(tr)
	// Most important phrase.
	fmt.Println(rankedPhrases[0])

	// Get all words order by weight.
	words := textrank.FindSingleWords(tr)
	// Most important word.
	fmt.Println(words[0])

	// Get the most important 10 sentences. Importance by phrase weights.
	sentences := textrank.FindSentencesByRelationWeight(tr, 10)
	// Found sentences
	fmt.Println(sentences)

	// Get the most important 10 sentences. Importance by word occurrence.
	sentences = textrank.FindSentencesByWordQtyWeight(tr, 10)
	// Found sentences
	fmt.Println(sentences)

	// Get the first 10 sentences, start from 5th sentence.
	sentences = textrank.FindSentencesFrom(tr, 5, 10)
	// Found sentences
	fmt.Println(sentences)

	// Get sentences by phrase/word chains order by position in text.
	sentencesPh := textrank.FindSentencesByPhraseChain(tr, []string{"gnome", "shell", "extension"})
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
	
	"github.com/DavidBelicza/TextRank/v2"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.NewDefaultAlgorithm()

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get the rank graph.
	rankData := tr.GetRankData()

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
	
	"github.com/DavidBelicza/TextRank/v2"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.NewDefaultAlgorithm()

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	rawText2 := "Another book or article..."
	rawText3 := "Third another book or article..."

	// Add text to the previously added text.
	tr.Populate(rawText2, language, rule)
	// Add text to the previously added text.
	tr.Populate(rawText3, language, rule)
	// Run the ranking to the whole composed text.
	tr.Ranking(algorithmDef)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(tr)

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
	
	"github.com/DavidBelicza/TextRank/v2"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	// Using a little bit more complex algorithm to ranking text.
	algorithmChain := textrank.NewChainAlgorithm()

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmChain)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(tr)

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
	
	"github.com/DavidBelicza/TextRank/v2"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// 1th TextRank object
	tr1 := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.NewDefaultAlgorithm()

	// Add text.
	tr1.Populate(rawText, language, rule)
	// Run the ranking.
	tr1.Ranking(algorithmDef)

	// 2nd TextRank object
	tr2 := textrank.NewTextRank()

	// Using a little bit more complex algorithm to ranking text.
	algorithmChain := textrank.NewChainAlgorithm()

	// Add text to the second graph.
	tr2.Populate(rawText, language, rule)
	// Run the ranking on the second graph.
	tr2.Ranking(algorithmChain)

	// Get all phrases by weight from first graph.
	rankedPhrases := textrank.FindPhrases(tr1)

	// Most important phrase from first graph.
	fmt.Println(rankedPhrases[0])
	// Second important phrase from first graph.
	fmt.Println(rankedPhrases[1])

	// Get all phrases by weight from second graph.
	rankedPhrases2 := textrank.FindPhrases(tr2)

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
	
	"github.com/DavidBelicza/TextRank/v2"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()

	// Add Spanish stop words (just some example).
	language.SetWords("es", []string{"uno", "dos", "tres", "yo", "es", "eres"})
	// Active the Spanish.
	language.SetActiveLanguage("es")

	// Default algorithm for ranking text.
	algorithmDef := textrank.NewDefaultAlgorithm()

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(tr)

	// Most important phrase.
	fmt.Println(rankedPhrases[0])
	// Second important phrase.
	fmt.Println(rankedPhrases[1])
}
```

### Asynchronous usage by goroutines

It is thread safe. Independent graphs can receive texts in the same time and can be extended by more text also in the same time.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/DavidBelicza/TextRank/v2"
)

func main() {
	// A flag when program has to stop.
	stopProgram := false
	// Channel.
	stream := make(chan string)
	// TextRank object.
	tr := textrank.NewTextRank()

	// Open new thread/routine
	go func(tr *textrank.TextRank) {
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
	}(tr)

	// Open new thread/routine
	go func() {
		// Counter how many times texts added to the ranking.
		i := 1

		for {
			// Get text from stream channel when it got a new one.
			rawText := <-stream

			// Default Rule for parsing.
			rule := textrank.NewDefaultRule()
			// Default Language for filtering stop words.
			language := textrank.NewDefaultLanguage()
			// Default algorithm for ranking text.
			algorithm := textrank.NewDefaultAlgorithm()

			// Add text.
			tr.Populate(rawText, language, rule)
			// Run the ranking.
			tr.Ranking(algorithm)

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
	phrases := textrank.FindPhrases(tr)
	// Second important phrase.
	fmt.Println(phrases[0])
}
```

## A SIMPLE VISUAL REPRESENTATION

The below image is a representation how works the simplest text ranking algorithm. This algorithm can be replaced by an another one by inject different Algorithm interface implementation.

<img src="https://i.imgur.com/RUdDfBz.jpg" />
