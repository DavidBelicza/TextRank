package rank

import "sort"

func Calculate(ranks *Rank) {
	updateRanks(ranks)
}

type Phrase struct {
	Right  string
	Left  string
	Weight float32
	Qty    int
}

func GetPhrases(rank *Rank) []Phrase {
	var phrases []Phrase

	for x, xMap := range rank.Relation.Scores {
		for y, _ := range xMap {
			phrases = append(phrases, Phrase{
				rank.Words[x].Value,
				rank.Words[y].Value,
				rank.Relation.Scores[x][y].Weight,
				rank.Relation.Scores[x][y].Qty,
			})
		}
	}

	sort.Slice(phrases, func(i, j int) bool {
		return phrases[i].Weight > phrases[j].Weight
	})

	return phrases
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
}

func weighting(qty int, min int, max int) float32 {
	return (float32(qty) - float32(min)) / (float32(max) - float32(min))
}
