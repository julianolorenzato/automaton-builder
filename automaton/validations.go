package automaton

import "errors"

func validateChain(am *AdjacencyMatrix, is *int, fs *[]int) error {
	return errors.Join(
		validateLinesAndColumns(*am),
	)
}

func validateLinesAndColumns(am AdjacencyMatrix) error {
	var lines int
	var columns int

	for i := range am {
		lines++

		for range am[i] {
			columns++
		}
	}

	if lines != columns {
		return errors.New("the adjacency matrix must be quadratic")
	}

	return nil
}

func validateInitialState(am AdjacencyMatrix) {

}
