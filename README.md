# TextRank on Go

[![GoDoc](https://godoc.org/github.com/DavidBelicza/TextRank?status.svg)](https://godoc.org/github.com/DavidBelicza/TextRank)
[![License: MIT](https://img.shields.io/badge/License-MIT-ee00ee.svg)](https://github.com/DavidBelicza/TextRank/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/DavidBelicza/TextRank.svg?branch=master)](https://travis-ci.org/DavidBelicza/TextRank)
[![Go Report Card](https://goreportcard.com/badge/github.com/DavidBelicza/TextRank)](https://goreportcard.com/report/github.com/DavidBelicza/TextRank)
[![Coverage Status](https://coveralls.io/repos/github/DavidBelicza/TextRank/badge.svg?branch=master)](https://coveralls.io/github/DavidBelicza/TextRank?branch=master)

## TextRank or Automatic summarization
> Automatic summarization is the process of reducing a text document with a computer program in order to create a summary that retains the most important points of the original document. Technologies that can make a coherent summary take into account variables such as length, writing style and syntax. Automatic data summarization is part of machine learning and data mining. The main idea of summarization is to find a representative subset of the data, which contains the information of the entire set. Summarization technologies are used in a large number of sectors in industry today. - Wikipedia

## Simplest regular usage
```go
package main

import (
	"fmt"
	
	"github.com/DavidBelicza/TextRank"
)

func main() {
	newTRid := 1
	rawText := "Your raw text, lorem ipsum..."
	rule := textrank.CreateDefaultRule()
	language := textrank.CreateDefaultLanguage()
	algorithm := textrank.CreateDefaultAlgorithm()

	textrank.Append(rawText, language, rule, newTRid)
	textrank.Ranking(newTRid, algorithm)

	rankedPhrases := textrank.FindPhrases(newTRid)

	fmt.Println(rankedPhrases)
}
```

## A simple visual representation
<img src="http://i.picresize.com/images/2018/01/30/PTn3Y.png" />
