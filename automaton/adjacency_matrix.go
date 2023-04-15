package automaton

import (
	"errors"
	"math"
	"strings"
)

var text string = "- 0 1\n- 1 -\n- 0 -\n"

func NewAdjacencyMatrix() ([][]*string, error) {
	// Format the raw text
	symbols := strings.Fields(text)

	symbolsLen := len(symbols)
	degree := math.Sqrt(float64(symbolsLen))
	mod := math.Mod(degree, 1)

	// Throw an error if the degree is not an integer (not quadratic matrix)
	// There are other ways to a matrix not be quadratic, so, I will need to validate these ways too :(
	if mod != 0 {
		return nil, errors.New("the matrix must be quadratic")
	}

	// Initialize a 2D slice that will be the representation of the adjacency matrix
	m := make([][]*string, int(degree))
	for i := range m {
		m[i] = make([]*string, int(degree))
	}

	// Fill the 2D slice with the symbols or nil
	for i := 0; i < int(degree); i++ {
		for j := 0; j < int(degree); j++ {
			index := i*3 + j

			if symbols[index] == "-" {
				m[i][j] = nil
			} else {
				m[i][j] = &symbols[index]
			}

		}
	}

	return m, nil
}
