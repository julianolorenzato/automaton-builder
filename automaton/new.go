package automaton

import (
	"automaton-builder/util"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type State struct {
	// data    any
	isFinal bool
}

type StatesPair struct {
	from *State
	to   *State
}

type Transition map[StatesPair]*string

type DFA struct {
	// Name         string
	// Alphabet     []string
	CurrentState *State
	InitialState *State
	States       []State
	Transitions  map[string][]StatesPair
}

// type NFA struct {
// 	// Name         string
// 	Alphabet     []string
// 	States       []State
// 	InitialState *State
// 	Transitions  map[StatesPair][]string
// }

// Params: m - matrix; fss - final states
func NewDFA(m [][]string, fss []int) (*DFA, error) {
	re := regexp.MustCompile(`^\(.*\)|\s*-\s*$`)

	dfa := DFA{
		// Transitions: map[StatesPair][]string{},
		Transitions: map[string][]StatesPair{},
	}

	// Add a state for each row of matrix.
	for i := 0; i < len(m); i++ {
		isFinal := util.Contains(fss, func(index int) bool {
			return i == index
		})

		dfa.States = append(dfa.States, State{
			isFinal: isFinal,
		})
	}

	dfa.InitialState = &dfa.States[0]
	dfa.CurrentState = dfa.InitialState

	rows := len(m)

	for i := range m {
		columns := len(m[i])

		if rows != columns {
			return nil, errors.New("the adjacency matrix must be quadratic")
		}

		for j := range m[i] {
			element := m[i][j]

			r := strconv.Itoa(i)
			c := strconv.Itoa(j)

			if !re.MatchString(element) {
				return nil, errors.New("the element M" + r + c + " must be a tuple or '-'.")
			}

			// If element is not a tuple skip the rest of this iteration
			if element[0] != '(' {
				continue
			}

			// Remove the brackets and whitespaces and split into possibly multiple symbols
			element = element[1 : len(element)-1]
			element = strings.ReplaceAll(element, " ", "")
			symbols := strings.Split(element, ";")

			// Add a new states pair for each symbol
			for _, v := range symbols {
				dfa.Transitions[v] = append(dfa.Transitions[v], StatesPair{
					from: &dfa.States[i],
					to:   &dfa.States[j],
				})
			}

			// dfa.Transitions[StatesPair{
			// 	from: &dfa.States[i],
			// 	to:   &dfa.States[j],
			// }] = symbols
		}
	}

	return &dfa, nil
}

func (dfa *DFA) Transition(s string) error {
	sp := util.Find(dfa.Transitions[s], func(sp StatesPair) bool {
		return sp.from == dfa.CurrentState
	})

	// if len(sps) > 1 {
	// 	panic("a deterministic finite automaton must not have more than one possibility of transition for the same symbol")
	// }

	if sp.to == nil {
		return errors.New("no transitions with the symbol " + s + " from the current state")
	}

	dfa.CurrentState = sp.to

	return nil
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
