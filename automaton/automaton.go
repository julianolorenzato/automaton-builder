package automaton

import (
	"automaton-builder/util"
	"errors"
)

type Automaton struct {
	name string
	// alphabet     []string
	states       []State
	currentState State
	initialState State
	// finalStates  []State
	transitions []Transition
}

type Transition struct {
	symbol string
	from   State
	to     State
}

type State struct {
	id      int
	isFinal bool
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
		// Verify if state "i" is a final state
		isFinal := util.Contains(fs, func(finalState int) bool {
			return i == finalState
		})
		
		// Each line index represent one state id
		if isFinal {
			a.states = append(a.states, State{id: i, isFinal: true})
		} else {
			a.states = append(a.states, State{id: i, isFinal: false})
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

func (a *Automaton) GetInfo() (name string, states, finalStates []State) {
	name = a.name
	states = a.states
	finalStates = util.FindAll(a.states, func (s State) bool {
		return s.isFinal
	})

	return
}

func (a *Automaton) GetName() string {
	return a.name
}

func (a *Automaton) GetAllStates() []State {
	return a.states
}

func (a *Automaton) GetFinalStates() []State {
	finalStates := util.FindAll(a.states, func (s State) bool {
		return s.isFinal
	})

	return finalStates
}

// Delta
func (a *Automaton) Transition(sym string) error {
	trs := util.FindAll(a.transitions, func(tr Transition) bool {
		return (tr.from == a.currentState) && (tr.symbol == sym)
	})

	if len(trs) == 0 {
		return errors.New("no transition for symbol " + sym)
	}

	// The automaton is using the first transition of that have the symbol sym, need refactor in the future,
	// if the automaton be deterministic must have only 1 transition per state per symbol
	// if the automaton be non-deterministic he should test all transitions per state per symbol
	a.currentState = trs[0].from

	return nil
}

// Delta Chapéu
func (a *Automaton) Perform(word []string) (*State, bool, error) {
	for _, v := range word {
		err := a.Transition(v)

		if err != nil {
			return &a.currentState, false, err
		}
	}

	// isFinal := util.Contains(a.finalStates, func(s State) bool {
	// 	return s == a.currentState
	// })

	if a.currentState.isFinal {
		return &a.currentState, false, errors.New("does not end in a final state")
	}

	return &a.currentState, true, nil
}
