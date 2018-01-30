package tool

var rankGraph *Graph

func Run(currentGraph *Graph) {
	rankGraph = currentGraph

	updateRanks()
}

func updateRanks() {
	for _, word := range rankGraph.GetWordData() {
		//@todo implement
		word.Score = float64(word.Count)
	}
}
