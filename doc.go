/*
Package textrank is an implementation of Text Rank algorithm in Go with
extendable features (automatic summarization, phrase extraction). It supports
multithreading by goroutines. The package is under The MIT Licence.

MOTIVATION

If there was a program what could rank book size text's words, phrases and
sentences continuously on multiple threads and it would be opened to modifing by
objects, written in a simple, secure, static language and if it would be very
well documented... Now, here it is.

FEATURES

- Find the most important phrases.
- Find the most important words.
- Find the most important N sentences.
- Importance by phrase weights.
- Importance by word occurrence.
- Find the first N sentences, start from Xth sentence.
- Find sentences by phrase chains ordered by position in text.
- Access to the whole ranked data.
- Support more languages.
- Algorithm for weighting can be modified by interface implementation.
- Parser can be modified by interface implementation.
- Multi thread support.

EXAMPLES

Find the most important phrases:

This is the most basic and simplest usage of textrank.

	package main

	import (
		"fmt"

		"github.com/DavidBelicza/TextRank"
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

All possible pre-defined finder queries:

After ranking, the graph contains a lot of valuable data. There are functions in
textrank package what contains logic to retrieve those data from the graph.

	package main

	import (
		"fmt"

		"github.com/DavidBelicza/TextRank"
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

Access to everything

After ranking, the graph contains a lot of valuable data. The GetRank function
allows access to the graph and every data can be retrieved from this structure.

	package main

	import (
		"fmt"

		"github.com/DavidBelicza/TextRank"
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
		// The node of this word, it contains the related words and the
		// relation weight.
		fmt.Println(rankData.Relation.Node[wordId])
	}

Adding text continuously:

It is possibe to add more text after another texts already have been added. The
Ranking function can merge these multiple texts and it can recalculate the
weights and all related data.

	package main

	import (
		"fmt"

		"github.com/DavidBelicza/TextRank"
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

Using different algorithm to ranking text:

There are two algorithm has implemented, it is possible to write custom
algorithm by Algorithm interface and use it instead of defaults.

	package main

	import (
		"fmt"

		"github.com/DavidBelicza/TextRank"
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

Using multiple graphs:

Graph ID exists because it is possible run multiple independent text ranking
processes.

	package main

	import (
		"fmt"

		"github.com/DavidBelicza/TextRank"
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

Using different non-English languages:

Engish is used by default but it is possible to add any language. To use other
languages a stop word list is required what you can find here:
https://github.com/stopwords-iso

	package main

	import (
		"fmt"

		"github.com/DavidBelicza/TextRank"
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

Asynchronous usage by goroutines:

It is thread safe. Independent graphs can receive texts in the same time and can
be extended by more text also in the same time.

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
*/
package textrank
