package automaton

import (
	"automaton-builder/util"
	"errors"
)

type DFA struct {
	Name         string
	Alphabet     []string
	CurrentState *State
	InitialState *State
	States       []State
	Transitions  map[Pair]*State
}

// Params: m - matrix; fss - final states
func NewDFA(m [][]string, fss []int) (*DFA, error) {
	// re := regexp.MustCompile(`^\(.*\)|\s*-\s*$`)

	dfa := DFA{
		Transitions: make(map[Pair]*State),
	}

	// Add Alphabet Symbols to dfa (ignore first column)
	dfa.Alphabet = m[0][1:]

	// Remove first line after
	m = m[1:]

	// Add States to dfa (ignore first line)
	for i := range m {
		stateName := m[i][0]
		isFinal := stateName[0] == '*'

		dfa.States = append(dfa.States, State{
			name:    stateName,
			isFinal: isFinal,
		})
		// Remove first column after
		m[i] = m[i][1:]
	}

	dfa.InitialState = &dfa.States[0]
	dfa.CurrentState = dfa.InitialState

	// Add Transitions to dfa
	for i := range m {
		for j := range m[i] {
			to := util.Find(dfa.States, func(s State) bool {
				return m[i][j] == s.name
			})

			dfa.Transitions[Pair{
				from:   &dfa.States[i],
				symbol: &dfa.Alphabet[j],
			}] = &to
		}
	}

	return &dfa, nil
}

func (dfa *DFA) Transition(symbol string) error {
	var newState *State

	for pair, to := range dfa.Transitions {
		if *pair.symbol == symbol && pair.from == dfa.CurrentState {
			newState = to
		}
	}

	if newState == nil {
		return errors.New("no transitions with the symbol " + symbol + " from the current state")
	} else {
		dfa.CurrentState = newState
		return nil
	}
}

func (dfa *DFA) Perform(w []string) (bool, error) {
	for _, v := range w {

		err := dfa.Transition(v)

		if err != nil {
			return false, err
		}
	}

	if dfa.CurrentState.isFinal {
		return true, nil
	} else {
		return false, nil
	}
}

func (dfa *DFA) Reset() {
	dfa.CurrentState = dfa.InitialState
}
