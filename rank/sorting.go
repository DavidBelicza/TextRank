package rank

import (
	"sort"
)

type Phrase struct {
	LeftID  int
	RightID int
	Right   string
	Left    string
	Weight  float32
	Qty     int
}

func GetPhrases(ranks *Rank) []Phrase {
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

type SingleWord struct {
	ID     int
	Word   string
	Weight float32
	Qty    int
}

func GetSingleWords(ranks *Rank) []SingleWord {
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

type Sentence struct {
	ID    int
	Value string
}

const ByQty = 0
const ByRelation = 1

func GetSentences(ranks *Rank, kind int, limit int) []Sentence {
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
		singleWords := GetSingleWords(ranks)

		for _, singleWord := range singleWords {
			sentenceIDs := ranks.Words[singleWord.ID].SentenceIDs

			if collect(sentenceIDs) {
				return sentences
			}
		}
	} else if kind == ByRelation {
		phrases := GetPhrases(ranks)

		for _, phrase := range phrases {
			sentenceIDs := ranks.Relation.Node[phrase.LeftID][phrase.RightID].SentenceIDs

			if collect(sentenceIDs) {
				return sentences
			}
		}
	}

	return sentences
}

func GetSentencesByPhrases(ranks *Rank, words []string) []Sentence {
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

func GetSentencesFrom(ranks *Rank, id int, limit int) []Sentence {
	var sentences []Sentence

	limit = id + limit - 1

	for i := id; i <= limit; i++ {
		sentences = append(sentences, Sentence{i, ranks.SentenceMap[i]})
	}

	return sentences
}
