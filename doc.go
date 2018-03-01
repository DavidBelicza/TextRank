/*
TextRank implementation in Golang with extendable features (automatic
summarization, phrase extraction). It supports multithreading by goroutines.

"Automatic summarization is the process of reducing a text document with a
computer program in order to create a summary that retains the most important
points of the original document. Technologies that can make a coherent summary
take into account variables such as length, writing style and syntax. Automatic
data summarization is part of machine learning and data mining. The main idea
of summarization is to find a representative subset of the data, which contains
the information of the entire set. Summarization technologies are used in a
large number of sectors in industry today." - Wikipedia

	rawText := "Lorem ipsum sit dolor amet..."
	tr := NewTextRank()
	rule := CreateDefaultRule()
	language := CreateDefaultLanguage()
	algorithm := CreateDefaultAlgorithm()

	tr.Populate(rawText, language, rule)
	tr.Ranking(algorithmDef)

	FindSentencesByPhraseChain(tr, []string{
		"captain",
		"james",
		"kirk",
	})

	FindPhrases(textRank)
*/
package textrank
