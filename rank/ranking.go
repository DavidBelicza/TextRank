package rank

import "fmt"

var rankGraph *Rank

func Calculate(currentGraph *Rank) {
	rankGraph = currentGraph

	updateRanks()
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

	fmt.Println(rankGraph.Relation.Scores)
}

func weighting(qty int, min int, max int) float32  {
	return (float32(qty) - float32(min)) / (float32(max) - float32(min))
}
