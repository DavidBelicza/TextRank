package textrank

import (
	"github.com/DavidBelicza/TextRank/convert"
	"github.com/DavidBelicza/TextRank/parse"
	"github.com/DavidBelicza/TextRank/rank"
)

var provider = make(map[int]*rank.Rank)

// CreateDefaultRule function retrieves a default Rule object what works in the
// most cases in English or similar Latin languages like French or Spanish. The
// Rule defines raw text how should be split to sentences and words. Because
// Rule is an interface it's possible modify the ranking by inject different
// Rule implementation. This is the 1th. step to use TextRank.
func CreateDefaultRule() *parse.RuleDefault {
	return parse.NewRule()
}

// CreateDefaultLanguage function retrieves a default Language object. It
// defines what words are real one and what words are just Stop Words or
// useless Junk Words. It uses the default English Stop Words, but it's
// possible to set different Stop Words in English or any other languages.
// Because Language is an interface it's possible to modify the ranking by
// inject different Language implementation. This is the 2nd. step to use
// TextRank.
func CreateDefaultLanguage() *convert.LanguageDefault {
	return convert.NewLanguage()
}

// CreateDefaultAlgorithm function retrieves an Algorithm object. It defines
// how should work the text ranking algorithm, the weighting. This is the
// general text rank by weighting the connection between the words to find
// the strongest phrases. Because Algorithm is an interface it's possible to
// modify the ranking  algorithm by inject different implementation. This is
// the 3rd. step to use  TextRank.
func CreateDefaultAlgorithm() *rank.AlgorithmDefault {
	return rank.NewAlgorithmDefault()
}

// CreateMixedAlgorithm function retrieves an Algorithm object. It defines
// how should work the text ranking algorithm, the weighting. This is an
// alternative way to ranking words by weighting the number of the words.
// Because Algorithm is an interface it's possible to modify the ranking
// algorithm by inject different implementation. This is  the 3rd. step to use
// TextRank.
func CreateMixedAlgorithm() *rank.AlgorithmMixed {
	return rank.NewAlgorithmMixed()
}

// Append function adds a raw text to the text-ranking graph. It parses,
// tokenize the raw text and prepares it to weighting and scoring. It's
// possible to append a new raw text to an existing one even if the previously
// text is already ranked. This is 4th. step to use TextRank.
// - text string must be a plain text from TXT or PDF or any document, it can
//   contain new lines, break lines or any unnecessary text parts, but it
//   should not contain html tags or codes.
// - lang Language object can be loaded from CreateDefaultLanguage function.
// - rule Rule object can be loaded from CreateDefaultRule function.
// - id int is the identifier of the TextRank. Because it's possible ranking
//   multiple texts in multiple languages with multiple rules in the same time
//   an ID is must have. If ranking is not exist by ID this function will
//   create a new ranking with the given ID.
func Append(text string, lang convert.Language, rule parse.Rule, id int) {
	var ranks *rank.Rank

	if savedTextRank, ok := provider[id]; ok {
		ranks = savedTextRank
	} else {
		ranks = rank.NewRank()
		provider[id] = ranks
	}

	parsedText := parse.TokenizeText(text, rule)

	for _, sentence := range parsedText.GetSentences() {
		convert.TextToRank(sentence, lang, provider[id])
	}
}

// Ranking function counts the words and connections between the words and
// weights the numbers then normalize them in type float32 between 0.00 and
// 1.00. This is the 5th step to use TextRank.
// - id int is the identifier of the TextRank. Because it's possible ranking
//   multiple texts in multiple languages with multiple rules in the same time
//   an ID is must have.
// - algorithm Algorithm is the object of the weighting and scoring methods.
func Ranking(id int, algorithm rank.Algorithm) {
	rank.Calculate(provider[id], algorithm)
}

// FindPhrases function retrieves a slice of Phrase structures by id what
// contain the sorted phrases with IDs, words, weights and quantities by weight
// from 1 to 0. Weight is calculated from quantities of relation between two
// words. One phrase is from two words, not less and more. (But it's possible
// to find chain of phrases by FindSentencesByPhraseChain function.)
func FindPhrases(id int) []rank.Phrase {
	return rank.FindPhrases(provider[id])
}

// FindSingleWords function retrieves a slice of SingleWord structures by id
// what contain the sorted words with IDs, words, weights and quantities by
// weight from 1 to 0. Weight is calculated from quantities of word.
func FindSingleWords(id int) []rank.SingleWord {
	return rank.FindSingleWords(provider[id])
}

// FindSentencesByRelationWeight function retrieves a slice of Sentence
// structures by id what contains the ID of the sentence and the sentence
// itself. The slice is sorted by weight of phrases from 1 to 0.
func FindSentencesByRelationWeight(id int, limit int) []rank.Sentence {
	return rank.FindSentences(provider[id], rank.ByRelation, limit)
}

// FindSentencesByWordQtyWeight function retrieves a slice of Sentence
// structures by id what contains the ID of the sentence and the sentence
// itself. The slice is sorted by weight of word quantities from 1 to 0.
func FindSentencesByWordQtyWeight(id int, limit int) []rank.Sentence {
	return rank.FindSentences(provider[id], rank.ByQty, limit)
}

// FindSentencesByPhraseChain function retrieves a slice of Sentence
// structures by id and slice of phrases what contains the ID of the sentence
// and the sentence itself. The slice is sorted by weight of word quantities
// from 1 to 0.
// - id int is the ID of the TextRank.
// - phrases []string is a slice of phrases. One phrase is from two words, so
//   when the slice contains 3 words the inner method will search for two
//   phrases.
//
//    rawText := "Long raw text, lorem ipsum..."
//    rule := CreateDefaultRule()
//    language := CreateDefaultLanguage()
//    algorithm := CreateDefaultAlgorithm()

//    Append(rawText, language, rule, 1)
//    Ranking(1, algorithm)
//
//    FindSentencesByPhraseChain(1, []string{
//        "captain",
//        "james",
//        "kirk",
//    })
//
// The above code search for captain james kirk, captain kirk james, james kirk
// captain, james captain kirk, kirk james captain and james kirk captain
// combinations in the graph. The 3 of words have to be related to each other
// in the same sentence but the search algorithm ignores the stop words. So if
// there is a sentence "James Kirk is the Captain." the sentence will be
// returned because the words "is" and "the" are stop words. In this case the
// possible combination is 3 factorial (3!) = 3 * 2 * 1.
func FindSentencesByPhraseChain(id int, phrases []string) []rank.Sentence {
	return rank.FindSentencesByPhrases(provider[id], phrases)
}

// FindSentencesFrom function retrieves a slice of Sentence structures by id of
// the TextRank and by id of the sentence what contains the ID of the sentence
// and the sentence itself. The returned slice contains sentences sorted by
// their IDs started from the given sentence ID.
func FindSentencesFrom(id int, sentenceID int, limit int) []rank.Sentence {
	return rank.FindSentencesFrom(provider[id], sentenceID, limit)
}

// GetRank function retrieves a TextRank by ID to that case if the developer
// want access to the whole graph and sentences, words, weights and all of the
// data to analyze it or just implement a new search or finder method.
func GetRank(id int) *rank.Rank {
	return provider[id]
}
