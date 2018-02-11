package rank

func Calculate(
	ranks *Rank,
	algorithm Algorithm,
) {
	updateRanks(ranks, algorithm)
}

func updateRanks(
	ranks *Rank,
	algorithm Algorithm,
) {
	for x, xMap := range ranks.Relation.Node {
		for y, _ := range xMap {
			qty := ranks.Relation.Node[x][y].Qty

			if ranks.Relation.Max < qty {
				ranks.Relation.Max = qty
			}

			if ranks.Relation.Min > qty || ranks.Relation.Min == 0 {
				ranks.Relation.Min = qty
			}
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
		weight := algorithm.WeightingHits(word.ID, word.Qty, ranks.Min, ranks.Max)
		word.Weight = weight
	}

	for x, xMap := range ranks.Relation.Node {
		for y, _ := range xMap {
			qty := ranks.Relation.Node[x][y].Qty
			sentenceIDs := ranks.Relation.Node[x][y].SentenceIDs
			weight := algorithm.WeightingRelation(
				x,
				y,
				qty,
				ranks.Relation.Min,
				ranks.Relation.Max,
				ranks.Words[x].Qty,
				ranks.Words[y].Qty,
				ranks.Min,
				ranks.Max,
			)
			ranks.Relation.Node[x][y] = Score{ranks.Relation.Node[x][y].Qty, weight, sentenceIDs}
		}
	}
}
