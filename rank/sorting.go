package rank

import (
	"sort"
)

// Phrase struct contains a single phrase and its data.
//
// LeftID is the ID of the word 1.
//
// RightID is the ID of the word 2.
//
// Left is the token of the word 1.
//
// Right is the token of the word 2.
//
// Weight is between 0.00 and 1.00.
//
// Qty is the occurrence of the phrase.
type Phrase struct {
	LeftID  int
	RightID int
	Left    string
	Right   string
	Weight  float32
	Qty     int
}

// FindPhrases function has wrapper textrank.FindPhrases. Use the wrapper
// instead.
func FindPhrases(ranks *Rank) []Phrase {
	var phrases []Phrase

	for x, xMap := range ranks.Relation.Node {
		for y := range xMap {
			phrases = append(phrases, Phrase{
				ranks.Words[x].ID,
				ranks.Words[y].ID,
				ranks.Words[x].Token,
				ranks.Words[y].Token,
				ranks.Relation.Node[x][y].Weight,
				ranks.Relation.Node[x][y].Qty,
			})
		}
	}

	sort.Slice(phrases, func(i, j int) bool {
		return phrases[i].Weight > phrases[j].Weight
	})

	return phrases
}

// SingleWord struct contains a single word and its data.
//
// ID of the word.
//
// Word itself, the token.
//
// Weight of the word between 0.00 and 1.00.
//
// Quantity of the word.
type SingleWord struct {
	ID     int
	Word   string
	Weight float32
	Qty    int
}

// FindSingleWords function has wrapper textrank.FindSingleWords. Use the
// wrapper instead.
func FindSingleWords(ranks *Rank) []SingleWord {
	var singleWords []SingleWord

	for _, word := range ranks.Words {
		singleWords = append(singleWords, SingleWord{
			word.ID,
			word.Token,
			word.Weight,
			word.Qty,
		})
	}

	sort.Slice(singleWords, func(i, j int) bool {
		return singleWords[i].Weight > singleWords[j].Weight
	})

	return singleWords
}

// Sentence struct contains a single sentence and its data.
type Sentence struct {
	ID    int
	Value string
}

// ByQty filter by occurrence of word.
const ByQty = 0

// ByRelation filter by phrase weight.
const ByRelation = 1

// FindSentences function has wrappers textrank.FindSentencesByRelationWeight
// and textrank.FindSentencesByWordQtyWeight. Use the wrappers instead.
func FindSentences(ranks *Rank, kind int, limit int) []Sentence {
	var sentences []Sentence

	cache := make(map[int]bool)

	collect := func(sentenceIDs []int) bool {
		for _, id := range sentenceIDs {
			if len(sentences) >= limit {
				return true
			}

			if !cache[id] {
				sentences = append(sentences, Sentence{id, ranks.SentenceMap[id]})
				cache[id] = true
			}
		}

		return false
	}

	if kind == ByQty {
		singleWords := FindSingleWords(ranks)

		for _, singleWord := range singleWords {
			sentenceIDs := ranks.Words[singleWord.ID].SentenceIDs

			if collect(sentenceIDs) {
				return sentences
			}
		}
	} else if kind == ByRelation {
		phrases := FindPhrases(ranks)

		for _, phrase := range phrases {
			sentenceIDs := ranks.Relation.Node[phrase.LeftID][phrase.RightID].SentenceIDs

			if collect(sentenceIDs) {
				return sentences
			}
		}
	}

	return sentences
}

// FindSentencesByPhrases function has wrapper
// textrank.FindSentencesByPhraseChain. Use the wrapper instead.
func FindSentencesByPhrases(ranks *Rank, words []string) []Sentence {
	var sentences []Sentence

	reqMatch := len(words) - 1
	sentenceIDs := make(map[int]int)

	for _, i := range words {
		for _, j := range words {
			x := ranks.WordValID[i]
			y := ranks.WordValID[j]

			if _, ok := ranks.Relation.Node[x][y]; ok {
				curSentenceIDs := ranks.Relation.Node[x][y].SentenceIDs

				for _, id := range curSentenceIDs {
					if _, ok := sentenceIDs[id]; ok {
						sentenceIDs[id]++
					} else {
						sentenceIDs[id] = 1
					}
				}
			}
		}
	}

	for sentenceID, v := range sentenceIDs {
		if v >= reqMatch {
			sentences = append(sentences, Sentence{sentenceID, ranks.SentenceMap[sentenceID]})
		}
	}

	sort.Slice(sentences, func(i, j int) bool {
		return sentences[i].ID < sentences[j].ID
	})

	return sentences
}

// FindSentencesFrom function has wrapper textrank.FindSentencesFrom. Use the
// wrapper instead.
func FindSentencesFrom(ranks *Rank, id int, limit int) []Sentence {
	var sentences []Sentence

	limit = id + limit - 1

	for i := id; i <= limit; i++ {
		sentences = append(sentences, Sentence{i, ranks.SentenceMap[i]})
	}

	return sentences
}
