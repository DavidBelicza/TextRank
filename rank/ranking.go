package rank

import "sort"

func Calculate(ranks *Rank) {
	updateRanks(ranks)
}

type Phrase struct {
	Right  string
	Left   string
	Weight float32
	Qty    int
}

func GetPhrases(ranks *Rank) []Phrase {
	var phrases []Phrase

	for x, xMap := range ranks.Relation.Scores {
		for y, _ := range xMap {
			phrases = append(phrases, Phrase{
				ranks.Words[x].Value,
				ranks.Words[y].Value,
				ranks.Relation.Scores[x][y].Weight,
				ranks.Relation.Scores[x][y].Qty,
			})
		}
	}

	sort.Slice(phrases, func(i, j int) bool {
		return phrases[i].Weight > phrases[j].Weight
	})

	return phrases
}

type SingleWord struct {
	Word   string
	Weight float32
	Qty    int
}

func GetSingleWords(ranks *Rank) []SingleWord {
	var singleWords []SingleWord

	for _, word := range ranks.Words {
		singleWords = append(singleWords, SingleWord{
			word.Value,
			word.Weight,
			word.Qty,
		})
	}

	sort.Slice(singleWords, func(i, j int) bool {
		return singleWords[i].Weight > singleWords[j].Weight
	})

	return singleWords
}

//@todo
func GetSentences(ranks *Rank, kind int) {
	// by score - relations weights or word qtys
}

//@todo
func GetSentencesByPhrases() {
	// [w1, w2], [w1, w2], [w1], [w1, w2]
}

func updateRanks(ranks *Rank) {
	for x, xMap := range ranks.Relation.Scores {
		for y, _ := range xMap {
			qty := ranks.Relation.Scores[x][y].Qty

			if ranks.Relation.Max < qty {
				ranks.Relation.Max = qty
			}

			if ranks.Relation.Min > qty || ranks.Relation.Min == 0 {
				ranks.Relation.Min = qty
			}
		}
	}

	for x, xMap := range ranks.Relation.Scores {
		for y, _ := range xMap {
			qty := ranks.Relation.Scores[x][y].Qty
			weight := weighting(qty, ranks.Relation.Min, ranks.Relation.Max)
			ranks.Relation.Scores[x][y] = Score{ranks.Relation.Scores[x][y].Qty, weight}
		}
	}

	for _, word := range ranks.Words {
		if ranks.Max < word.Qty {
			ranks.Max = word.Qty
		}

		if ranks.Min > word.Qty || ranks.Min == 0 {
			ranks.Min = word.Qty
		}
	}

	for _, word := range ranks.Words {
		weight := weighting(word.Qty, ranks.Min, ranks.Max)
		word.Weight = weight
	}
}

func weighting(qty int, min int, max int) float32 {
	return (float32(qty) - float32(min)) / (float32(max) - float32(min))
}
