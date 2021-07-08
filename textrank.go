package textrank

import (
	"github.com/DavidBelicza/TextRank/v2/convert"
	"github.com/DavidBelicza/TextRank/v2/parse"
	"github.com/DavidBelicza/TextRank/v2/rank"
)

// TextRank structure contains the Rank data object. This structure is a wrapper
// around the whole text ranking functionality.
type TextRank struct {
	rank *rank.Rank
}

// NewTextRank constructor retrieves a TextRank pointer. This is the 1th step to
// use TextRank.
func NewTextRank() *TextRank {
	return &TextRank{
		rank.NewRank(),
	}
}

// NewDefaultRule function retrieves a default Rule object what works in the
// most cases in English or similar Latin languages like French or Spanish. The
// Rule defines raw text how should be split to sentences and words. Because
// Rule is an interface it's possible modify the ranking by inject different
// Rule implementation. This is the 2nd step to use TextRank.
func NewDefaultRule() *parse.RuleDefault {
	return parse.NewRule()
}

// NewDefaultLanguage function retrieves a default Language object. It defines
// what words are real and what words are just Stop Words or useless Junk Words.
// It uses the default English Stop Words, but it's possible to set different
// Stop Words in English or any other languages. Because Language is an
// interface it's possible to modify the ranking by inject different Language
// implementation. This is the 3rd step to use TextRank.
func NewDefaultLanguage() *convert.LanguageDefault {
	return convert.NewLanguage()
}

// NewDefaultAlgorithm function retrieves an Algorithm object. It defines how
// should work the text ranking algorithm, the weighting. This is the general
// text rank by weighting the connection between the words to find the strongest
// phrases. Because Algorithm is an interface it's possible to modify the
// ranking algorithm by inject different implementation. This is the 4th step to
// use TextRank.
func NewDefaultAlgorithm() *rank.AlgorithmDefault {
	return rank.NewAlgorithmDefault()
}

// NewChainAlgorithm function retrieves an Algorithm object. It defines how
// should work the text ranking algorithm, the weighting. This is an alternative
// way to ranking words by weighting the number of the words. Because Algorithm
// is an interface it's possible to modify the ranking algorithm by inject
// different implementation. This is  the 4th step to use TextRank.
func NewChainAlgorithm() *rank.AlgorithmChain {
	return rank.NewAlgorithmChain()
}

// Populate method adds a raw text to the text-ranking graph. It parses,
// tokenize the raw text and prepares it to weighting and scoring. It's possible
// to append a new raw text to an existing one even if the previously text is
// already ranked. This is 5th step to use TextRank.
//
// text string must be a plain text from TXT or PDF or any document, it can
// contain new lines, break lines or any unnecessary text parts, but it should
// not contain HTML tags or codes.
//
// lang Language object can be loaded from NewDefaultLanguage function.
//
// rule Rule object can be loaded from NewDefaultRule function.
func (textRank *TextRank) Populate(
	text string,
	lang convert.Language,
	rule parse.Rule,
) {
	parsedText := parse.TokenizeText(text, rule)

	for _, sentence := range parsedText.GetSentences() {
		convert.TextToRank(sentence, lang, textRank.rank)
	}
}

// Ranking method counts the words and connections between the words, then it
// weights the numbers then normalize them in type float32 between 0.00 and
// 1.00. This is the 6th step to use TextRank.
//
// algorithm Algorithm is the object of the weighting and scoring methods.
func (textRank *TextRank) Ranking(algorithm rank.Algorithm) {
	rank.Calculate(textRank.rank, algorithm)
}

// GetRankData method retrieves the Rank data to that case if the developer want
// access to the whole graph and sentences, words, weights and all of the data
// to analyze it or just implement a new search logic or finder method.
func (textRank *TextRank) GetRankData() *rank.Rank {
	return textRank.rank
}

// FindPhrases function retrieves a slice of Phrase structures by TextRank
// object. The return value contains the sorted phrases with IDs, words, weights
// and quantities by weight from 1 to 0. Weight is calculated from quantities of
// relation between two words. A single phrase is from two words - not less and
// more. (But it's possible to find chain of phrases by
// FindSentencesByPhraseChain function.)
func FindPhrases(textRank *TextRank) []rank.Phrase {
	return rank.FindPhrases(textRank.rank)
}

// FindSingleWords function retrieves a slice of SingleWord structures by
// TextRank object. The return value contains the sorted words with IDs, words,
// weights and quantities by weight from 1 to 0. Weight is calculated from
// quantities of word.
func FindSingleWords(textRank *TextRank) []rank.SingleWord {
	return rank.FindSingleWords(textRank.rank)
}

// FindSentencesByRelationWeight function retrieves a slice of Sentence
// structures by TextRank object. The return value contains the ID of the
// sentence and the sentence text itself. The slice is sorted by weight of
// phrases from 1 to 0.
func FindSentencesByRelationWeight(
	textRank *TextRank,
	limit int,
) []rank.Sentence {

	return rank.FindSentences(textRank.rank, rank.ByRelation, limit)
}

// FindSentencesByWordQtyWeight function retrieves a slice of Sentence
// structures by TextRank object. The return value contains the ID of the
// sentence and the sentence text itself. The slice is sorted by weight of word
// quantities from 1 to 0.
func FindSentencesByWordQtyWeight(
	textRank *TextRank,
	limit int,
) []rank.Sentence {

	return rank.FindSentences(textRank.rank, rank.ByQty, limit)
}

// FindSentencesByPhraseChain function retrieves a slice of Sentence structures
// by TextRank object and slice of phrases. The return value contains the ID of
// the sentence and the sentence text itself. The slice is sorted by weight of
// word quantities from 1 to 0.
//
// textRank TextRank is the object of the TextRank.
//
// phrases []string is a slice of phrases. A single phrase is from two words, so
// when the slice contains 3 words the inner method will search for two phrases.
// The search algorithm seeks for "len(phrases)!". In case of three item the
// possible combination is 3 factorial (3!) = 3 * 2 * 1.
//
//    rawText := "Long raw text, lorem ipsum..."
//    rule := NewDefaultRule()
//    language := NewDefaultLanguage()
//    algorithm := NewDefaultAlgorithm()
//
//    Append(rawText, language, rule, 1)
//    Ranking(1, algorithm)
//
//    FindSentencesByPhraseChain(1, []string{
//        "captain",
//        "james",
//        "kirk",
//    })
//
// The above code searches for captain james kirk, captain kirk james, james
// kirk captain, james captain kirk, kirk james captain and james kirk captain
// combinations in the graph. The 3 of words have to be related to each other
// in the same sentence but the search algorithm ignores the stop words. So if
// there is a sentence "James Kirk is the Captain of the Enterprise." the
// sentence will be returned because the words "is" and "the" are stop words.
func FindSentencesByPhraseChain(
	textRank *TextRank,
	phrases []string,
) []rank.Sentence {

	return rank.FindSentencesByPhrases(textRank.rank, phrases)
}

// FindSentencesFrom function retrieves a slice of Sentence structures by
// TextRank object and by ID of the sentence. The return value contains the
// sentence text itself. The returned slice contains sentences sorted by their
// IDs started from the given sentence ID in ascending sort.
func FindSentencesFrom(
	textRank *TextRank,
	sentenceID int,
	limit int,
) []rank.Sentence {

	return rank.FindSentencesFrom(textRank.rank, sentenceID, limit)
}
