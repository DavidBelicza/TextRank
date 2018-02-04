package rank

import "sort"

var rankGraph *Rank

func Calculate(currentGraph *Rank) {
	rankGraph = currentGraph

	updateRanks()
}

type Phrase struct {
	Word1  string
	Word2  string
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

func updateRanks() {
	for x, xMap := range rankGraph.Relation.Scores {
		for y, _ := range xMap {
			qty := rankGraph.Relation.Scores[x][y].Qty

			if rankGraph.Relation.Max < qty {
				rankGraph.Relation.Max = qty
			}

			if rankGraph.Relation.Min > qty || rankGraph.Relation.Min == 0 {
				rankGraph.Relation.Min = qty
			}
		}
	}

	for x, xMap := range rankGraph.Relation.Scores {
		for y, _ := range xMap {
			qty := rankGraph.Relation.Scores[x][y].Qty
			weight := weighting(qty, rankGraph.Relation.Min, rankGraph.Relation.Max)
			rankGraph.Relation.Scores[x][y] = Score{rankGraph.Relation.Scores[x][y].Qty, weight}
		}
	}
}

func weighting(qty int, min int, max int) float32 {
	return (float32(qty) - float32(min)) / (float32(max) - float32(min))
}
