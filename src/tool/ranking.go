package tool

var rankGraph *Rank

func Run(currentGraph *Rank) {
	rankGraph = currentGraph

	updateRanks()
}

//@todo implement
func updateRanks() {
	/*for _, word := range rankGraph.GetWordData() {
		for wordID, count := range word.ConnectionLeft {

		}
	}*/
}
