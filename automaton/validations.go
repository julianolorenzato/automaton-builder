package automaton

import (
	"errors"
	"strconv"
)

// Perform all validations and return any errors found
func validateChain(m [][]string, is int, fs []int) error {
	return errors.Join(
		validateQuadratic(m),
		validateInitialState(m, is),
		validateFinalStates(m, fs),
	)
}

// Check if the adjacency matrix is quadratic
func validateQuadratic(am [][]string) error {
	lines := len(am)

	for i := range am {
		columns := len((am)[i])

		if lines != columns {
			return errors.New("the adjacency matrix must be quadratic")
		}
	}

	return nil
}

// Check if the initial state is present in adjacency matrix
func validateInitialState(am [][]string, is int) error {
	if len(am) < is {
		return errors.New("initial state must be present in adjacency matrix")
	}

	return nil
}

// Check if all final states are present in adjacency matrix
func validateFinalStates(am [][]string, fs []int) error {
	for _, v := range fs {
		if len(am) < v {
			str := strconv.Itoa(v)

			return errors.New("final state " + str + " is not present in adjacency matrix")
		}
	}

	return nil
}
