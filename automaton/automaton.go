package automaton

import (
	"automaton-builder/util"
)

type Automaton struct {
	name string
	// alphabet     []string
	states       []State
	initialState State
	finalStates  []State
	transitions  []Transition
}

type Transition struct {
	symbol string
	from   State
	to     State
}

type State struct {
	id int
}

// Params (AdjacencyMatrix, InitialState, FinalStates)
func NewAutomaton(am AdjacencyMatrix, is int, fs []int) (*Automaton, error) {
	// Perfom validation in the input data
	err := validateChain(&am, is, fs)
	if err != nil {
		return nil, err
	}

	a := Automaton{}

	a.initialState = State{id: is}

	// Iterate on lines of adjacency matrix
	for i := range am {
		// Each line index represent one state id
		a.states = append(a.states, State{id: i})

		// Verify if state "i" is a final state
		isFinal := util.Contains(fs, func(finalState int) bool {
			return i == finalState
		})

		if isFinal {
			a.finalStates = append(a.finalStates, State{id: i})
		}

		//Iterate on columns of adjacency matrix
		for j, v := range am[i] {

			// If the value Aij is not nil there is a transition from i to j with symbol Aij
			if v != nil {
				a.transitions = append(a.transitions, Transition{
					symbol: *v,
					from:   State{id: i},
					to:     State{id: j},
				})
			}

		}
	}

	return &a, nil
}
