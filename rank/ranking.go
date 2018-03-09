package rank

// Calculate function ranking words by the given algorithm implementation.
func Calculate(ranks *Rank, algorithm Algorithm) {
	updateRanks(ranks, algorithm)
}

func updateRanks(ranks *Rank, algorithm Algorithm) {
	for _, word := range ranks.Words {
		weight := algorithm.WeightingHits(word.ID, ranks)
		word.Weight = weight

		if ranks.Max < word.Weight {
			ranks.Max = word.Weight
		}

		if ranks.Min > word.Weight || ranks.Min == 0 {
			ranks.Min = word.Weight
		}
	}

	for _, word := range ranks.Words {
		word.Weight = normalize(word.Weight, ranks.Min, ranks.Max)
	}

	for x, xMap := range ranks.Relation.Node {
		for y := range xMap {
			sentenceIDs := ranks.Relation.Node[x][y].SentenceIDs
			weight := algorithm.WeightingRelation(x, y, ranks)

			ranks.Relation.Node[x][y] = Score{
				ranks.Relation.Node[x][y].Qty,
				weight,
				sentenceIDs,
			}

			if ranks.Relation.Max < weight {
				ranks.Relation.Max = weight
			}

			if ranks.Relation.Min > weight || ranks.Relation.Min == 0 {
				ranks.Relation.Min = weight
			}
		}
	}

	for x, xMap := range ranks.Relation.Node {
		for y := range xMap {
			weight := normalize(
				ranks.Relation.Node[x][y].Weight,
				ranks.Relation.Min,
				ranks.Relation.Max,
			)

			ranks.Relation.Node[x][y] = Score{
				ranks.Relation.Node[x][y].Qty,
				weight,
				ranks.Relation.Node[x][y].SentenceIDs,
			}
		}
	}
}

func normalize(weight float32, min float32, max float32) float32 {
	return (weight - min) / (max - min)
}
